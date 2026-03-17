// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// solhint-disable no-inline-assembly

/// @dev Metadata for a registered global plugin.
///      Packed into a single struct to minimize storage reads.
struct GlobalPluginInfo {
    address plugin;     // plugin contract address
    uint64 registeredAt; // block.timestamp when registered (fits until year 2554)
    bool active;        // soft-disable without removing
}

struct GlobalPluginStorage {
    /// @dev Ordered list of global plugins (max bounded by MAX_GLOBAL_PLUGINS)
    GlobalPluginInfo[] plugins;
    /// @dev plugin address → index+1 in plugins array (0 = not registered)
    ///      Avoids O(n) lookup on add/remove. Gas: 1 SLOAD instead of loop.
    mapping(address => uint256) pluginIndex;
}

/// @title LibGlobalPluginStorage
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Namespaced storage for protocol-wide global plugins.
///         Global plugins are cross-asset services (marketplace, AMM, governance)
///         that extend the Diamond protocol without being tied to a specific tokenId.
///         slot = keccak256("diamond.erc3643.globalplugin.storage") - 1
library LibGlobalPluginStorage {
    bytes32 internal constant POSITION =
        0x7ff8986cfa674a3f1057a68c0b1723f898c792c44fee3667f9ee7ffb0430217b;

    function layout() internal pure returns (GlobalPluginStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}
