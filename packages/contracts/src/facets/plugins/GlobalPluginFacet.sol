// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibDiamond} from "../../libraries/LibDiamond.sol";
import {LibAccessStorage} from "../../storage/LibAccessStorage.sol";
import {LibGlobalPluginStorage, GlobalPluginStorage, GlobalPluginInfo} from "../../storage/LibGlobalPluginStorage.sol";
import {IPluginModule} from "../../interfaces/plugins/IPluginModule.sol";

/*//////////////////////////////////////////////////////////////
                            ERRORS
//////////////////////////////////////////////////////////////*/

error GlobalPluginFacet__Unauthorized();
error GlobalPluginFacet__ZeroAddress();
error GlobalPluginFacet__AlreadyRegistered(address plugin);
error GlobalPluginFacet__NotRegistered(address plugin);
error GlobalPluginFacet__TooManyPlugins(uint256 count, uint256 max);
error GlobalPluginFacet__AlreadyActive(address plugin);
error GlobalPluginFacet__AlreadyInactive(address plugin);

/**
 * @title GlobalPluginFacet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Manages protocol-wide global plugins that extend Diamond functionality
 *         across all tokenIds. Global plugins are cross-asset services like
 *         marketplaces, AMMs, voting systems, or governance modules.
 *
 *         Unlike asset plugins (per-tokenId, hookable), global plugins:
 *         - Operate across all assets in the protocol
 *         - Do NOT receive balance-change hooks
 *         - Inherit the Diamond's role-based permission system
 *         - Can be versioned: remove v1, add v2
 *
 * @dev Uses indexed mapping (pluginIndex) for O(1) add/remove/lookup.
 *      Plugin addresses are validated via IPluginModule.name() call on registration
 *      to ensure the contract implements the base interface.
 *      Only Diamond owner or COMPLIANCE_ADMIN can manage global plugins.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract GlobalPluginFacet {
    /*//////////////////////////////////////////////////////////////
                                CONSTANTS
    //////////////////////////////////////////////////////////////*/

    /// @dev Maximum global plugins to prevent unbounded storage growth.
    uint256 internal constant MAX_GLOBAL_PLUGINS = 20;

    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event GlobalPluginRegistered(address indexed plugin, string name);
    event GlobalPluginRemoved(address indexed plugin);
    event GlobalPluginStatusChanged(address indexed plugin, bool active);

    /*//////////////////////////////////////////////////////////////
                        EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Registers a new global plugin.
    ///         Validates the plugin implements IPluginModule by calling name().
    /// @param plugin The plugin contract address
    function registerGlobalPlugin(address plugin) external {
        _enforceComplianceAdminOrOwner();
        if (plugin == address(0)) revert GlobalPluginFacet__ZeroAddress();

        GlobalPluginStorage storage gs = LibGlobalPluginStorage.layout();

        if (gs.pluginIndex[plugin] != 0) {
            revert GlobalPluginFacet__AlreadyRegistered(plugin);
        }
        if (gs.plugins.length >= MAX_GLOBAL_PLUGINS) {
            revert GlobalPluginFacet__TooManyPlugins(gs.plugins.length + 1, MAX_GLOBAL_PLUGINS);
        }

        // Validate plugin implements IPluginModule — reverts if not
        string memory pluginName = IPluginModule(plugin).name();

        gs.plugins.push(GlobalPluginInfo({
            plugin: plugin,
            registeredAt: uint64(block.timestamp),
            active: true
        }));
        // Store index+1 (0 means not registered)
        gs.pluginIndex[plugin] = gs.plugins.length;

        emit GlobalPluginRegistered(plugin, pluginName);
    }

    /// @notice Removes a global plugin from the registry.
    ///         Uses swap-and-pop with index update for O(1) removal.
    /// @param plugin The plugin contract address to remove
    function removeGlobalPlugin(address plugin) external {
        _enforceComplianceAdminOrOwner();

        GlobalPluginStorage storage gs = LibGlobalPluginStorage.layout();
        uint256 idx = gs.pluginIndex[plugin];
        if (idx == 0) revert GlobalPluginFacet__NotRegistered(plugin);

        uint256 arrayIdx = idx - 1;
        uint256 lastIdx = gs.plugins.length - 1;

        // Swap with last element if not already last
        if (arrayIdx != lastIdx) {
            GlobalPluginInfo storage lastPlugin = gs.plugins[lastIdx];
            gs.plugins[arrayIdx] = lastPlugin;
            gs.pluginIndex[lastPlugin.plugin] = idx; // update swapped element's index
        }

        gs.plugins.pop();
        delete gs.pluginIndex[plugin];

        emit GlobalPluginRemoved(plugin);
    }

    /// @notice Activates or deactivates a global plugin without removing it.
    ///         Useful for temporary disable during upgrades or incidents.
    /// @param plugin The plugin contract address
    /// @param active Whether the plugin should be active
    function setGlobalPluginStatus(address plugin, bool active) external {
        _enforceComplianceAdminOrOwner();

        GlobalPluginStorage storage gs = LibGlobalPluginStorage.layout();
        uint256 idx = gs.pluginIndex[plugin];
        if (idx == 0) revert GlobalPluginFacet__NotRegistered(plugin);

        GlobalPluginInfo storage info = gs.plugins[idx - 1];
        if (info.active == active) {
            if (active) revert GlobalPluginFacet__AlreadyActive(plugin);
            else revert GlobalPluginFacet__AlreadyInactive(plugin);
        }

        info.active = active;

        emit GlobalPluginStatusChanged(plugin, active);
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL VIEWS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns all registered global plugins (active and inactive).
    /// @return plugins Array of GlobalPluginInfo structs
    function getGlobalPlugins() external view returns (GlobalPluginInfo[] memory plugins) {
        plugins = LibGlobalPluginStorage.layout().plugins;
    }

    /// @notice Returns only active global plugins.
    /// @return activePlugins Array of addresses of active plugins
    function getActiveGlobalPlugins() external view returns (address[] memory activePlugins) {
        GlobalPluginInfo[] storage all = LibGlobalPluginStorage.layout().plugins;
        uint256 len = all.length;

        // Count active first to allocate exact size
        uint256 activeCount;
        for (uint256 i; i < len;) {
            if (all[i].active) {
                unchecked { ++activeCount; }
            }
            unchecked { ++i; }
        }

        activePlugins = new address[](activeCount);
        uint256 j;
        for (uint256 i; i < len;) {
            if (all[i].active) {
                activePlugins[j] = all[i].plugin;
                unchecked { ++j; }
            }
            unchecked { ++i; }
        }
    }

    /// @notice Returns the info for a specific global plugin.
    /// @param plugin The plugin contract address
    /// @return info The plugin info (address, registeredAt, active)
    function getGlobalPluginInfo(address plugin) external view returns (GlobalPluginInfo memory info) {
        GlobalPluginStorage storage gs = LibGlobalPluginStorage.layout();
        uint256 idx = gs.pluginIndex[plugin];
        if (idx == 0) revert GlobalPluginFacet__NotRegistered(plugin);
        info = gs.plugins[idx - 1];
    }

    /// @notice Returns true if a plugin is registered (active or inactive).
    /// @param plugin The plugin contract address
    /// @return registered Whether the plugin is registered
    function isGlobalPlugin(address plugin) external view returns (bool registered) {
        registered = LibGlobalPluginStorage.layout().pluginIndex[plugin] != 0;
    }

    /// @notice Returns the count of registered global plugins.
    /// @return count Number of registered global plugins
    function globalPluginCount() external view returns (uint256 count) {
        count = LibGlobalPluginStorage.layout().plugins.length;
    }

    /*//////////////////////////////////////////////////////////////
                        INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    bytes32 internal constant COMPLIANCE_ADMIN = keccak256("COMPLIANCE_ADMIN");

    function _enforceComplianceAdminOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        bool isAdmin = LibAccessStorage.layout().roles[COMPLIANCE_ADMIN][msg.sender];
        if (!isOwner && !isAdmin) revert GlobalPluginFacet__Unauthorized();
    }
}
