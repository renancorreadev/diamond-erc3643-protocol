// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";
import {IDiamond, IDiamondCut} from "../src/interfaces/core/IDiamond.sol";
import {DiamondABI} from "../src/DiamondABI.sol";

/// @title EIP1967Init
/// @notice Sets the EIP-1967 implementation slot to point to the DiamondABI contract.
///         This makes Polygonscan detect the proxy and show "Read/Write as Proxy" tab
///         with the combined ABI of all facets.
contract EIP1967Init {
    /// @dev EIP-1967 implementation slot: keccak256("eip1967.proxy.implementation") - 1
    bytes32 internal constant _IMPLEMENTATION_SLOT =
        0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;

    function init(address implementation) external {
        assembly {
            sstore(_IMPLEMENTATION_SLOT, implementation)
        }
    }
}

/// @title SetDiamondImplementation
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Deploys a DiamondABI stub and sets the EIP-1967 implementation slot
///         so Polygonscan shows "Read/Write as Proxy" with all facet functions.
contract SetDiamondImplementation is Script {
    function run() external {
        address diamond = vm.envAddress("DIAMOND_ADDRESS");

        vm.startBroadcast();

        // Deploy the DiamondABI stub (empty contract with all function signatures)
        DiamondABI diamondABI = new DiamondABI();

        // Deploy the init contract
        EIP1967Init eip1967Init = new EIP1967Init();

        // Set EIP-1967 implementation slot to DiamondABI via diamondCut + init
        IDiamond.FacetCut[] memory cuts = new IDiamond.FacetCut[](0);

        IDiamondCut(diamond).diamondCut(
            cuts,
            address(eip1967Init),
            abi.encodeCall(EIP1967Init.init, (address(diamondABI)))
        );

        vm.stopBroadcast();

        console2.log("DiamondABI stub deployed at:", address(diamondABI));
        console2.log("EIP-1967 implementation slot set on Diamond:", diamond);
        console2.log("Polygonscan will now show all Read/Write functions under 'Read/Write as Proxy'");
    }
}
