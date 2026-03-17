// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";
import {IDiamond, IDiamondCut, IDiamondLoupe} from "../src/interfaces/core/IDiamond.sol";
import {AssetManagerFacet} from "../src/facets/token/AssetManagerFacet.sol";
import {ERC1155Facet} from "../src/facets/token/ERC1155Facet.sol";
import {SupplyFacet} from "../src/facets/token/SupplyFacet.sol";
import {GlobalPluginFacet} from "../src/facets/plugins/GlobalPluginFacet.sol";
import {PluginRouterFacet} from "../src/facets/routers/PluginRouterFacet.sol";

/// @title UpgradePluginSystem
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Upgrades the Diamond to add the Plugin System:
///         - Adds GlobalPluginFacet (8 selectors) — protocol-wide plugin registry
///         - Adds PluginRouterFacet (1 selector) — routes onAction to hookable plugins
///         - Adds 4 new selectors to AssetManagerFacet (addPluginModule, removePluginModule, etc.)
///         - Replaces ERC1155Facet and SupplyFacet implementations (plugin hook integration)
///
/// Usage:
///   DIAMOND_ADDRESS=0x... forge script script/UpgradePluginSystem.s.sol \
///     --rpc-url $RPC_URL --account deployer --broadcast -vvvv
contract UpgradePluginSystem is Script {
    error Upgrade__DiamondAddressZero();
    error Upgrade__UnexpectedFacetCount(uint256 expected, uint256 actual);

    function run() external {
        address diamond = vm.envAddress("DIAMOND_ADDRESS");
        if (diamond == address(0)) revert Upgrade__DiamondAddressZero();

        vm.startBroadcast();

        // ── 1. Deploy new facet implementations ──────────────────────

        GlobalPluginFacet globalPluginFacet = new GlobalPluginFacet();
        PluginRouterFacet pluginRouterFacet = new PluginRouterFacet();
        AssetManagerFacet newAssetManagerFacet = new AssetManagerFacet();
        ERC1155Facet newERC1155Facet = new ERC1155Facet();
        SupplyFacet newSupplyFacet = new SupplyFacet();

        // ── 2. Build facet cuts ──────────────────────────────────────
        //
        // 5 cuts total:
        //   [0] Add    — GlobalPluginFacet (8 new selectors)
        //   [1] Add    — PluginRouterFacet (1 new selector)
        //   [2] Add    — AssetManagerFacet new selectors only (4 plugin functions)
        //   [3] Replace — ERC1155Facet (same 7 selectors, new bytecode)
        //   [4] Replace — SupplyFacet (same 7 selectors, new bytecode)
        //
        // Note: AssetManagerFacet existing 14 selectors already point to the old
        // implementation. We only Add the 4 new selectors here. The existing selectors
        // still work — they route to the old bytecode which is functionally identical
        // for those functions (no logic change in registerAsset, addComplianceModule, etc.).
        // If you want all 18 selectors on the same new deployment, uncomment the
        // Replace cut below instead.

        IDiamond.FacetCut[] memory cuts = new IDiamond.FacetCut[](5);

        // [0] GlobalPluginFacet — Add all 8 selectors
        cuts[0] = IDiamond.FacetCut({
            facetAddress: address(globalPluginFacet),
            action: IDiamond.FacetCutAction.Add,
            functionSelectors: _globalPluginSelectors()
        });

        // [1] PluginRouterFacet — Add 1 selector
        cuts[1] = IDiamond.FacetCut({
            facetAddress: address(pluginRouterFacet),
            action: IDiamond.FacetCutAction.Add,
            functionSelectors: _pluginRouterSelectors()
        });

        // [2] AssetManagerFacet — Add 4 new plugin selectors
        cuts[2] = IDiamond.FacetCut({
            facetAddress: address(newAssetManagerFacet),
            action: IDiamond.FacetCutAction.Add,
            functionSelectors: _assetManagerNewSelectors()
        });

        // [3] ERC1155Facet — Replace (same ABI, internal plugin hooks added)
        cuts[3] = IDiamond.FacetCut({
            facetAddress: address(newERC1155Facet),
            action: IDiamond.FacetCutAction.Replace,
            functionSelectors: _erc1155Selectors()
        });

        // [4] SupplyFacet — Replace (same ABI, internal plugin hooks added)
        cuts[4] = IDiamond.FacetCut({
            facetAddress: address(newSupplyFacet),
            action: IDiamond.FacetCutAction.Replace,
            functionSelectors: _supplySelectors()
        });

        // ── 3. Execute diamond cut ───────────────────────────────────

        IDiamondCut(diamond).diamondCut(cuts, address(0), "");

        vm.stopBroadcast();

        // ── 4. Log addresses ─────────────────────────────────────────

        console2.log("");
        console2.log("=== Plugin System Upgrade ===");
        console2.log("");
        console2.log("Diamond              :", diamond);
        console2.log("");
        console2.log("--- New Facets (Add) ---");
        console2.log("GlobalPluginFacet    :", address(globalPluginFacet));
        console2.log("PluginRouterFacet    :", address(pluginRouterFacet));
        console2.log("");
        console2.log("--- Updated Facets ---");
        console2.log("AssetManagerFacet    :", address(newAssetManagerFacet), "(4 new selectors)");
        console2.log("ERC1155Facet         :", address(newERC1155Facet), "(replaced)");
        console2.log("SupplyFacet          :", address(newSupplyFacet), "(replaced)");
        console2.log("");

        // ── 5. Verify ────────────────────────────────────────────────

        uint256 facetCount = IDiamondLoupe(diamond).facetAddresses().length;
        // Was 19 facets (DiamondCutFacet + 18 from deploy script cuts)
        // Now 19 + 3 new facet addresses = 22
        // But AssetManagerFacet new selectors point to a NEW address, so
        // the old AssetManagerFacet address still has 14 selectors,
        // and the new address has 4 selectors. That's 2 facet addresses for AssetManager.
        // Total: 19 (original) + 3 (GlobalPlugin, PluginRouter, new AssetManager) = 22
        console2.log("Facet addresses      :", facetCount);
        console2.log("");
        console2.log("=== Upgrade Complete ===");
    }

    /*//////////////////////////////////////////////////////////////
                        SELECTOR BUILDERS
    //////////////////////////////////////////////////////////////*/

    function _globalPluginSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](8);
        sels[0] = GlobalPluginFacet.registerGlobalPlugin.selector;
        sels[1] = GlobalPluginFacet.removeGlobalPlugin.selector;
        sels[2] = GlobalPluginFacet.setGlobalPluginStatus.selector;
        sels[3] = GlobalPluginFacet.getGlobalPlugins.selector;
        sels[4] = GlobalPluginFacet.getActiveGlobalPlugins.selector;
        sels[5] = GlobalPluginFacet.getGlobalPluginInfo.selector;
        sels[6] = GlobalPluginFacet.isGlobalPlugin.selector;
        sels[7] = GlobalPluginFacet.globalPluginCount.selector;
    }

    function _pluginRouterSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](1);
        sels[0] = PluginRouterFacet.pluginAction.selector;
    }

    /// @dev Only the 4 NEW plugin management functions.
    ///      Existing AssetManagerFacet selectors (registerAsset, etc.) stay on the old implementation.
    function _assetManagerNewSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](4);
        sels[0] = AssetManagerFacet.addPluginModule.selector;
        sels[1] = AssetManagerFacet.removePluginModule.selector;
        sels[2] = AssetManagerFacet.setPluginModules.selector;
        sels[3] = AssetManagerFacet.getPluginModules.selector;
    }

    function _erc1155Selectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](7);
        sels[0] = ERC1155Facet.safeTransferFrom.selector;
        sels[1] = ERC1155Facet.safeBatchTransferFrom.selector;
        sels[2] = ERC1155Facet.setApprovalForAll.selector;
        sels[3] = ERC1155Facet.balanceOf.selector;
        sels[4] = ERC1155Facet.balanceOfBatch.selector;
        sels[5] = ERC1155Facet.isApprovedForAll.selector;
        sels[6] = ERC1155Facet.partitionBalanceOf.selector;
    }

    function _supplySelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](7);
        sels[0] = SupplyFacet.mint.selector;
        sels[1] = SupplyFacet.batchMint.selector;
        sels[2] = SupplyFacet.burn.selector;
        sels[3] = SupplyFacet.forcedTransfer.selector;
        sels[4] = SupplyFacet.totalSupply.selector;
        sels[5] = SupplyFacet.holderCount.selector;
        sels[6] = SupplyFacet.isHolder.selector;
    }
}
