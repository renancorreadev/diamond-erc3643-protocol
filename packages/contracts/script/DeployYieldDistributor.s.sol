// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";

import {YieldDistributorModule} from
    "../src/plugins/modules/YieldDistributorModule/YieldDistributorModule.sol";

/// @title DeployYieldDistributor
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Deploys the YieldDistributorModule as a standalone contract.
///         After deploy, register it per-asset via:
///           Diamond.addPluginModule(tokenId, yieldModuleAddress)
///
/// Usage:
///   DIAMOND_ADDRESS=0xAb07CEf1BEeDBb30F5795418c79879794b31C521 \
///   forge script script/DeployYieldDistributor.s.sol \
///     --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv
///
/// Environment:
///   DIAMOND_ADDRESS — Diamond proxy address (required)
///   OWNER_ADDRESS   — Module owner (defaults to msg.sender)
contract DeployYieldDistributor is Script {
    function run() external {
        address diamond = vm.envAddress("DIAMOND_ADDRESS");
        address owner = vm.envOr("OWNER_ADDRESS", msg.sender);

        vm.startBroadcast();

        YieldDistributorModule yieldModule = new YieldDistributorModule(diamond, owner);

        vm.stopBroadcast();

        console2.log("");
        console2.log("=== YieldDistributorModule ===");
        console2.log("");
        console2.log("Diamond              :", diamond);
        console2.log("Owner                :", owner);
        console2.log("YieldDistributor     :", address(yieldModule));
        console2.log("");
        console2.log("Next step: register per-asset via Diamond.addPluginModule(tokenId, address)");
        console2.log("");
    }
}
