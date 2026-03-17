// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IHookablePlugin} from "../../interfaces/plugins/IHookablePlugin.sol";
import {LibAssetStorage, AssetStorage, AssetConfig} from "../../storage/LibAssetStorage.sol";

/**
 * @title PluginRouterFacet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Routes post-hook calls to per-tokenId hookable plugin modules.
 *         Plugin modules are pluggable extensions that react to token
 *         state changes (transfer, mint, burn) without gating them.
 * @dev Called by ERC1155Facet and SupplyFacet after balance mutations.
 *      If no plugin modules are set for a tokenId, hooks are no-ops
 *      (single SLOAD for the empty array length).
 *      Uses a single `pluginAction` entry point — each module receives
 *      an ActionParams struct and decides internally which types to handle.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract PluginRouterFacet {
    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error PluginRouterFacet__AssetNotRegistered(uint256 tokenId);

    /*//////////////////////////////////////////////////////////////
                        EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Forwards a balance-changing action to all plugin modules for a tokenId
    /// @param params The action parameters (actionType, tokenId, operator, from, to, amount)
    function pluginAction(IHookablePlugin.ActionParams calldata params) external {
        address[] storage modules = _getModules(params.tokenId);
        uint256 len = modules.length;
        for (uint256 i; i < len;) {
            IHookablePlugin(modules[i]).onAction(params);
            unchecked { ++i; }
        }
    }

    /*//////////////////////////////////////////////////////////////
                        INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    function _getModules(uint256 tokenId) internal view returns (address[] storage) {
        return LibAssetStorage.layout().configs[tokenId].pluginModules;
    }
}
