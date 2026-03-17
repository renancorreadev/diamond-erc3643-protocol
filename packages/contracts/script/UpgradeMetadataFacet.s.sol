// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";
import {IDiamond, IDiamondCut} from "../src/interfaces/core/IDiamond.sol";
import {MetadataFacet} from "../src/facets/token/MetadataFacet.sol";
import {LibAppStorage, AppStorage} from "../src/libraries/LibAppStorage.sol";

/// @dev Initializer that sets contract-level name and symbol in AppStorage.
contract MetadataNameInit {
    function init(string calldata name_, string calldata symbol_) external {
        AppStorage storage app = LibAppStorage.layout();
        app.contractName = name_;
        app.contractSymbol = symbol_;
    }
}

/// @title UpgradeMetadataFacet
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Upgrades MetadataFacet to add name() and symbol() without args
///         so Polygonscan shows the token name in Token Tracker.
contract UpgradeMetadataFacet is Script {
    function run() external {
        address diamond = vm.envAddress("DIAMOND_ADDRESS");
        string memory tokenName = vm.envOr("TOKEN_NAME", string("Diamond RWA"));
        string memory tokenSymbol = vm.envOr("TOKEN_SYMBOL", string("dRWA"));

        vm.startBroadcast();

        // Deploy new MetadataFacet with name()/symbol() overloads
        MetadataFacet newMetadata = new MetadataFacet();

        // Deploy initializer for setting contract name/symbol
        MetadataNameInit nameInit = new MetadataNameInit();

        // Replace old MetadataFacet + add 2 new selectors
        IDiamond.FacetCut[] memory cuts = new IDiamond.FacetCut[](2);

        // 1. Replace existing selectors with new facet address
        bytes4[] memory replaceSelectors = new bytes4[](7);
        replaceSelectors[0] = bytes4(0x0e89341c); // uri(uint256)
        replaceSelectors[1] = bytes4(0x00ad800c); // name(uint256)
        replaceSelectors[2] = bytes4(0x4e41a1fb); // symbol(uint256)
        replaceSelectors[3] = MetadataFacet.supplyCap.selector;
        replaceSelectors[4] = MetadataFacet.issuer.selector;
        replaceSelectors[5] = MetadataFacet.allowedCountries.selector;
        replaceSelectors[6] = MetadataFacet.tokenInfo.selector;

        cuts[0] = IDiamond.FacetCut({
            facetAddress: address(newMetadata),
            action: IDiamond.FacetCutAction.Replace,
            functionSelectors: replaceSelectors
        });

        // 2. Add new selectors: name() and symbol()
        bytes4[] memory addSelectors = new bytes4[](2);
        addSelectors[0] = bytes4(0x06fdde03); // name()
        addSelectors[1] = bytes4(0x95d89b41); // symbol()

        cuts[1] = IDiamond.FacetCut({
            facetAddress: address(newMetadata),
            action: IDiamond.FacetCutAction.Add,
            functionSelectors: addSelectors
        });

        // Execute cut + init
        IDiamondCut(diamond).diamondCut(
            cuts, address(nameInit), abi.encodeCall(MetadataNameInit.init, (tokenName, tokenSymbol))
        );

        vm.stopBroadcast();

        console2.log("MetadataFacet upgraded:", address(newMetadata));
        console2.log("Contract name set to:", tokenName);
        console2.log("Contract symbol set to:", tokenSymbol);
    }
}
