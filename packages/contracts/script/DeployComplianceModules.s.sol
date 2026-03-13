// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";

import {CountryRestrictModule} from "../src/compliance/modules/CountryRestrictModule.sol";
import {MaxBalanceModule} from "../src/compliance/modules/MaxBalanceModule.sol";
import {MaxHoldersModule} from "../src/compliance/modules/MaxHoldersModule.sol";

/// @title DeployComplianceModules
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Deploys standalone compliance modules that plug into the Diamond.
///
/// Usage:
///   forge script script/DeployComplianceModules.s.sol \
///     --rpc-url $RPC_URL --account deployer --broadcast -vvvv
///
/// Environment:
///   DIAMOND_ADDRESS — Diamond proxy address (required)
///   OWNER_ADDRESS   — Module owner (defaults to msg.sender)
contract DeployComplianceModules is Script {
    function run() external {
        address diamond = vm.envAddress("DIAMOND_ADDRESS");
        address owner = vm.envOr("OWNER_ADDRESS", msg.sender);

        vm.startBroadcast();

        CountryRestrictModule countryRestrict = new CountryRestrictModule(diamond, owner);
        MaxBalanceModule maxBalance = new MaxBalanceModule(diamond, owner);
        MaxHoldersModule maxHolders = new MaxHoldersModule(diamond, owner);

        vm.stopBroadcast();

        console2.log("");
        console2.log("=== Compliance Modules ===");
        console2.log("");
        console2.log("Diamond              :", diamond);
        console2.log("Owner                :", owner);
        console2.log("");
        console2.log("CountryRestrictModule:", address(countryRestrict));
        console2.log("MaxBalanceModule     :", address(maxBalance));
        console2.log("MaxHoldersModule     :", address(maxHolders));
        console2.log("");
    }
}
