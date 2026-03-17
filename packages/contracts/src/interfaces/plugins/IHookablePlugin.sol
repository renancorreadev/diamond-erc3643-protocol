// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {IPluginModule} from "./IPluginModule.sol";

/// @title IHookablePlugin
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Extended plugin interface for modules that receive automatic callbacks
///         on every balance-changing event (transfer, mint, burn).
///         Use this for features that need to track holder balances:
///         yield distribution, vesting schedules, loyalty programs, tax logic, etc.
///
/// @dev Unlike compliance modules (which gate transfers via canTransfer),
///      hookable plugins are purely reactive — they observe state changes
///      and update their own internal state accordingly.
///      Hooks fire AFTER balance mutation in the Diamond.
///
///      Action semantics:
///      - Mint:     actionType = Mint,     from = address(0), to = holder
///      - Burn:     actionType = Burn,     from = holder,     to = address(0)
///      - Transfer: actionType = Transfer, from = holder A,   to = holder B
///
///      Each module decides internally which action types to handle.
///      Unhandled types can simply return without side effects.
/// @custom:security-contact renan.correa@hubweb3.com
interface IHookablePlugin is IPluginModule {
    /// @notice Type of balance-changing action
    enum ActionType { Transfer, Mint, Burn }

    /// @notice Parameters passed to plugin hooks
    /// @param actionType The type of action (Transfer, Mint, Burn)
    /// @param tokenId The asset class affected
    /// @param operator The address that initiated the action (msg.sender in Diamond)
    /// @param from Source address (address(0) for mints)
    /// @param to Destination address (address(0) for burns)
    /// @param amount Amount of tokens affected
    struct ActionParams {
        ActionType actionType;
        uint256 tokenId;
        address operator;
        address from;
        address to;
        uint256 amount;
    }

    /// @notice Called after a balance-changing action is executed in the Diamond
    /// @dev Balance mutation has already occurred when this hook fires.
    ///      The module decides internally how to handle each ActionType.
    /// @param params The action parameters (calldata struct for gas efficiency)
    function onAction(ActionParams calldata params) external;
}
