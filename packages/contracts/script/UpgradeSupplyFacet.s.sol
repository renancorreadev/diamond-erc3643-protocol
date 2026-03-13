// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";
import {IDiamond, IDiamondCut} from "../src/interfaces/core/IDiamond.sol";
import {SupplyFacet} from "../src/facets/token/SupplyFacet.sol";

/// @title UpgradeSupplyFacet
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Replaces the SupplyFacet on the Diamond to add ERC-1155 TransferSingle events
///         on mint, burn, and forcedTransfer — required for explorers and indexers.
///
/// Usage:
///   DIAMOND_ADDRESS=0x... forge script script/UpgradeSupplyFacet.s.sol \
///     --rpc-url $RPC_URL --account deployer --broadcast -vvvv
contract UpgradeSupplyFacet is Script {
    function run() external {
        address diamond = vm.envAddress("DIAMOND_ADDRESS");

        vm.startBroadcast();

        // Deploy new SupplyFacet with TransferSingle events
        SupplyFacet newSupplyFacet = new SupplyFacet();

        // Build selector array (same selectors, new implementation)
        bytes4[] memory sels = new bytes4[](7);
        sels[0] = SupplyFacet.mint.selector;
        sels[1] = SupplyFacet.batchMint.selector;
        sels[2] = SupplyFacet.burn.selector;
        sels[3] = SupplyFacet.forcedTransfer.selector;
        sels[4] = SupplyFacet.totalSupply.selector;
        sels[5] = SupplyFacet.holderCount.selector;
        sels[6] = SupplyFacet.isHolder.selector;

        // Replace action = 1
        IDiamond.FacetCut[] memory cuts = new IDiamond.FacetCut[](1);
        cuts[0] = IDiamond.FacetCut({
            facetAddress: address(newSupplyFacet),
            action: IDiamond.FacetCutAction.Replace,
            functionSelectors: sels
        });

        IDiamondCut(diamond).diamondCut(cuts, address(0), "");

        vm.stopBroadcast();

        console2.log("New SupplyFacet deployed at:", address(newSupplyFacet));
        console2.log("Diamond upgraded at:", diamond);
    }
}
