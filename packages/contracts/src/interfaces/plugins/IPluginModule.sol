// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title IPluginModule
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Base interface for all pluggable modules that extend Diamond functionality.
///         Plugin modules are external contracts that add new capabilities to the
///         protocol without modifying the Diamond core (no diamondCut required).
///
///         Plugins are organized in two scopes:
///
///         **Global Plugins** — protocol-wide services registered via GlobalPluginFacet.
///         Operate across all tokenIds: marketplaces, AMMs, governance, voting.
///         Registered/removed cheaply (single SSTORE). Versionable: remove v1, add v2.
///
///         **Asset Plugins** — per-tokenId modules registered via AssetManagerFacet.
///         - **Hookable** (implement IHookablePlugin): receive automatic callbacks
///           on every balance-changing event (transfer, mint, burn).
///           Use for yield distribution, vesting, loyalty programs.
///         - **Service** (implement only IPluginModule): registered but not hooked.
///
/// @dev Plugins run as external contracts with their own storage.
///      They are NOT facets — they don't execute via delegatecall.
///      They inherit the Diamond's role-based permission system.
/// @custom:security-contact renan.correa@hubweb3.com
interface IPluginModule {
    /// @notice Returns the human-readable name of this plugin
    function name() external view returns (string memory);
}
