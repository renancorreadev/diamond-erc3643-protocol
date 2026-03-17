// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IHookablePlugin} from "../../src/interfaces/plugins/IHookablePlugin.sol";

/// @dev Mock hookable plugin for testing PluginRouterFacet and hook integration.
///      Tracks all onAction calls with parameters for assertion.
contract MockPluginModule is IHookablePlugin {
    ActionParams[] public actionCalls;

    bool public shouldRevert;

    function setRevert(bool revert_) external {
        shouldRevert = revert_;
    }

    function onAction(ActionParams calldata params) external {
        if (shouldRevert) revert("MockPluginModule: revert");
        actionCalls.push(params);
    }

    function name() external pure returns (string memory) {
        return "MockPlugin";
    }

    function actionCallCount() external view returns (uint256) {
        return actionCalls.length;
    }

    function getAction(uint256 index) external view returns (ActionParams memory) {
        return actionCalls[index];
    }
}
