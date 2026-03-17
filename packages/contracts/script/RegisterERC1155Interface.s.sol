// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";
import {IDiamond, IDiamondCut} from "../src/interfaces/core/IDiamond.sol";
import {LibDiamond, DiamondStorage} from "../src/libraries/LibDiamond.sol";

/// @title ERC1155InterfaceInit
/// @notice Registers ERC-1155 and ERC-1155 MetadataURI interface IDs in Diamond storage.
contract ERC1155InterfaceInit {
    /// @dev ERC-1155: 0xd9b67a26, ERC-1155 MetadataURI: 0x0e89341c
    function init() external {
        DiamondStorage storage ds = LibDiamond.diamondStorage();
        ds.supportedInterfaces[bytes4(0xd9b67a26)] = true; // IERC1155
        ds.supportedInterfaces[bytes4(0x0e89341c)] = true; // IERC1155MetadataURI
    }
}

/// @title RegisterERC1155Interface
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Registers ERC-1155 interface support on an existing Diamond deployment.
///
/// Usage:
///   DIAMOND_ADDRESS=0x... forge script script/RegisterERC1155Interface.s.sol \
///     --rpc-url $RPC_URL --account deployer --broadcast -vvvv
contract RegisterERC1155Interface is Script {
    function run() external {
        address diamond = vm.envAddress("DIAMOND_ADDRESS");

        vm.startBroadcast();

        ERC1155InterfaceInit initContract = new ERC1155InterfaceInit();

        IDiamond.FacetCut[] memory cuts = new IDiamond.FacetCut[](0);

        IDiamondCut(diamond).diamondCut(
            cuts,
            address(initContract),
            abi.encodeCall(ERC1155InterfaceInit.init, ())
        );

        vm.stopBroadcast();

        console2.log("ERC-1155 interface registered on Diamond:", diamond);
        console2.log("Init contract:", address(initContract));
    }
}
