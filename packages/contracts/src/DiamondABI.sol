// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IDiamond, IDiamondCut, IDiamondLoupe} from "./interfaces/core/IDiamond.sol";
import {AssetConfig} from "./storage/LibAssetStorage.sol";
import {AssetGroup} from "./storage/LibAssetGroupStorage.sol";

/// @title DiamondABI
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice ABI-only stub whose verified source on Polygonscan exposes all Diamond facet
///         functions under "Read/Write as Proxy". Has NO logic — every function reverts.
/// @dev Deploy once, verify on Polygonscan, then set EIP-1967 implementation slot on
///      the Diamond to point here. Polygonscan reads this ABI for the proxy UI.
contract DiamondABI {
    /*//////////////////////////////////////////////////////////////
                              STRUCTS
    //////////////////////////////////////////////////////////////*/

    struct RegisterAssetParams {
        string name;
        string symbol;
        string uri;
        uint256 supplyCap;
        uint32 identityProfileId;
        address[] complianceModules;
        address issuer;
        uint16[] allowedCountries;
    }

    struct CreateGroupParams {
        string name;
        uint256 parentTokenId;
        uint256 maxUnits;
    }

    struct MintUnitParams {
        uint256 groupId;
        string name;
        string symbol;
        string uri;
        uint256 supplyCap;
        address investor;
        uint256 amount;
    }

    struct Facet {
        address facetAddress;
        bytes4[] functionSelectors;
    }

    struct FacetCut {
        address facetAddress;
        IDiamond.FacetCutAction action;
        bytes4[] functionSelectors;
    }

    /*//////////////////////////////////////////////////////////////
                        DIAMOND CORE
    //////////////////////////////////////////////////////////////*/

    function diamondCut(FacetCut[] calldata _diamondCut, address _init, bytes calldata _calldata) external {}

    function facets() external view returns (Facet[] memory) {}
    function facetFunctionSelectors(address _facet) external view returns (bytes4[] memory) {}
    function facetAddresses() external view returns (address[] memory) {}
    function facetAddress(bytes4 _functionSelector) external view returns (address) {}
    function supportsInterface(bytes4 _interfaceId) external view returns (bool) {}

    /*//////////////////////////////////////////////////////////////
                          OWNERSHIP
    //////////////////////////////////////////////////////////////*/

    function transferOwnership(address _newOwner) external {}
    function acceptOwnership() external {}
    function owner() external view returns (address) {}
    function pendingOwner() external view returns (address) {}

    /*//////////////////////////////////////////////////////////////
                        ACCESS CONTROL
    //////////////////////////////////////////////////////////////*/

    bytes32 public constant GOVERNANCE_ROLE = keccak256("GOVERNANCE_ROLE");
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    bytes32 public constant COMPLIANCE_ADMIN = keccak256("COMPLIANCE_ADMIN");
    bytes32 public constant TRANSFER_AGENT = keccak256("TRANSFER_AGENT");
    bytes32 public constant RECOVERY_AGENT = keccak256("RECOVERY_AGENT");
    bytes32 public constant CLAIM_ISSUER_ROLE = keccak256("CLAIM_ISSUER_ROLE");

    function grantRole(bytes32 role, address account) external {}
    function revokeRole(bytes32 role, address account) external {}
    function renounceRole(bytes32 role) external {}
    function setRoleAdmin(bytes32 role, bytes32 adminRole) external {}
    function hasRole(bytes32 role, address account) external view returns (bool) {}
    function getRoleAdmin(bytes32 role) external view returns (bytes32) {}

    /*//////////////////////////////////////////////////////////////
                            PAUSE
    //////////////////////////////////////////////////////////////*/

    function pauseProtocol() external {}
    function unpauseProtocol() external {}
    function pauseAsset(uint256 tokenId) external {}
    function unpauseAsset(uint256 tokenId) external {}
    function isProtocolPaused() external view returns (bool) {}
    function isAssetPaused(uint256 tokenId) external view returns (bool) {}

    /*//////////////////////////////////////////////////////////////
                          EMERGENCY
    //////////////////////////////////////////////////////////////*/

    function emergencyPause() external {}
    function isEmergencyPaused() external view returns (bool) {}

    /*//////////////////////////////////////////////////////////////
                          ERC-1155
    //////////////////////////////////////////////////////////////*/

    function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes calldata data) external {}
    function safeBatchTransferFrom(
        address from,
        address to,
        uint256[] calldata ids,
        uint256[] calldata amounts,
        bytes calldata data
    ) external {}
    function setApprovalForAll(address operator, bool approved) external {}
    function balanceOf(address account, uint256 id) external view returns (uint256) {}
    function balanceOfBatch(address[] calldata accounts, uint256[] calldata ids)
        external
        view
        returns (uint256[] memory)
    {}
    function isApprovedForAll(address account, address operator) external view returns (bool) {}
    function partitionBalanceOf(address account, uint256 id)
        external
        view
        returns (uint256 free, uint256 locked, uint256 custody, uint256 pendingSettlement)
    {}

    /*//////////////////////////////////////////////////////////////
                            SUPPLY
    //////////////////////////////////////////////////////////////*/

    function mint(uint256 tokenId, address to, uint256 amount) external {}
    function batchMint(uint256[] calldata tokenIds, address[] calldata recipients, uint256[] calldata amounts)
        external
    {}
    function burn(uint256 tokenId, address from, uint256 amount) external {}
    function forcedTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes32 reasonCode) external {}
    function totalSupply(uint256 tokenId) external view returns (uint256) {}
    function holderCount(uint256 tokenId) external view returns (uint256) {}
    function isHolder(uint256 tokenId, address account) external view returns (bool) {}

    /*//////////////////////////////////////////////////////////////
                           METADATA
    //////////////////////////////////////////////////////////////*/

    function name() external view returns (string memory) {}
    function symbol() external view returns (string memory) {}
    function uri(uint256 id) external view returns (string memory) {}
    function name(uint256 tokenId) external view returns (string memory) {}
    function symbol(uint256 tokenId) external view returns (string memory) {}
    function supplyCap(uint256 tokenId) external view returns (uint256) {}
    function issuer(uint256 tokenId) external view returns (address) {}
    function allowedCountries(uint256 tokenId) external view returns (uint16[] memory) {}
    function tokenInfo(uint256 tokenId)
        external
        view
        returns (
            string memory name_,
            string memory symbol_,
            string memory uri_,
            uint256 totalSupply_,
            uint256 supplyCap_,
            uint256 holderCount_,
            address issuer_,
            bool paused_
        )
    {}

    /*//////////////////////////////////////////////////////////////
                        ASSET MANAGER
    //////////////////////////////////////////////////////////////*/

    function registerAsset(RegisterAssetParams calldata p) external returns (uint256 tokenId) {}
    function addComplianceModule(uint256 tokenId, address module) external {}
    function removeComplianceModule(uint256 tokenId, address module) external {}
    function setComplianceModules(uint256 tokenId, address[] calldata modules) external {}
    function setIdentityProfile(uint256 tokenId, uint32 profileId) external {}
    function setIssuer(uint256 tokenId, address newIssuer) external {}
    function setSupplyCap(uint256 tokenId, uint256 cap) external {}
    function setAllowedCountries(uint256 tokenId, uint16[] calldata countries) external {}
    function setAssetUri(uint256 tokenId, string calldata newUri) external {}
    function getAssetConfig(uint256 tokenId) external view returns (AssetConfig memory) {}
    function getComplianceModules(uint256 tokenId) external view returns (address[] memory) {}
    function getRegisteredTokenIds() external view returns (uint256[] memory) {}
    function assetExists(uint256 tokenId) external view returns (bool) {}
    function nextTokenId() external view returns (uint256) {}

    /*//////////////////////////////////////////////////////////////
                            FREEZE
    //////////////////////////////////////////////////////////////*/

    function setWalletFrozen(address wallet, bool frozen) external {}
    function setAssetWalletFrozen(uint256 tokenId, address wallet, bool frozen) external {}
    function setFrozenAmount(uint256 tokenId, address wallet, uint256 amount) external {}
    function setLockupExpiry(uint256 tokenId, address wallet, uint64 expiry) external {}
    function isWalletFrozen(address wallet) external view returns (bool) {}
    function isAssetWalletFrozen(uint256 tokenId, address wallet) external view returns (bool) {}
    function getFrozenAmount(uint256 tokenId, address wallet) external view returns (uint256) {}
    function getLockupExpiry(uint256 tokenId, address wallet) external view returns (uint64) {}

    /*//////////////////////////////////////////////////////////////
                          RECOVERY
    //////////////////////////////////////////////////////////////*/

    function recoverWallet(address lostWallet, address newWallet) external {}

    /*//////////////////////////////////////////////////////////////
                          SNAPSHOT
    //////////////////////////////////////////////////////////////*/

    function createSnapshot(uint256 tokenId) external returns (uint256 snapshotId) {}
    function recordHolder(uint256 snapshotId, address holder) external {}
    function recordHoldersBatch(uint256 snapshotId, address[] calldata holders) external {}
    function getSnapshot(uint256 snapshotId)
        external
        view
        returns (uint256 tokenId, uint256 totalSupplyVal, uint64 timestamp, uint256 holderCountVal)
    {}
    function getSnapshotBalance(uint256 snapshotId, address holder)
        external
        view
        returns (uint256 balance, bool recorded)
    {}
    function getTokenSnapshots(uint256 tokenId) external view returns (uint256[] memory) {}
    function getLatestSnapshotId(uint256 tokenId) external view returns (uint256) {}
    function nextSnapshotId() external view returns (uint256) {}

    /*//////////////////////////////////////////////////////////////
                          DIVIDEND
    //////////////////////////////////////////////////////////////*/

    function createDividend(uint256 snapshotId, uint256 totalAmount, address paymentToken)
        external
        payable
        returns (uint256 dividendId)
    {}
    function claimDividend(uint256 dividendId) external {}
    function getDividend(uint256 dividendId)
        external
        view
        returns (
            uint256 snapshotId,
            uint256 tokenId,
            uint256 totalAmount,
            address paymentToken,
            uint256 claimedAmount,
            uint64 createdAt
        )
    {}
    function hasClaimed(uint256 dividendId, address holder) external view returns (bool) {}
    function claimableAmount(uint256 dividendId, address holder) external view returns (uint256) {}
    function getTokenDividends(uint256 tokenId) external view returns (uint256[] memory) {}

    receive() external payable {}

    /*//////////////////////////////////////////////////////////////
                        ASSET GROUPS
    //////////////////////////////////////////////////////////////*/

    function createGroup(CreateGroupParams calldata params) external returns (uint256 groupId) {}
    function mintUnit(MintUnitParams calldata params) external returns (uint256 childTokenId) {}
    function mintUnitBatch(MintUnitParams[] calldata params) external returns (uint256[] memory childTokenIds) {}
    function getGroup(uint256 groupId) external view returns (AssetGroup memory) {}
    function getGroupChildren(uint256 groupId) external view returns (uint256[] memory) {}
    function getChildGroup(uint256 childTokenId) external view returns (uint256) {}
    function getRegisteredGroupIds() external view returns (uint256[] memory) {}
    function groupExists(uint256 groupId) external view returns (bool) {}

    /*//////////////////////////////////////////////////////////////
                      IDENTITY REGISTRY
    //////////////////////////////////////////////////////////////*/

    function registerIdentity(address wallet, address identity, uint16 country) external {}
    function deleteIdentity(address wallet) external {}
    function updateIdentity(address wallet, address identity) external {}
    function updateCountry(address wallet, uint16 country) external {}
    function batchRegisterIdentity(
        address[] calldata wallets,
        address[] calldata identities,
        uint16[] calldata countries
    ) external {}
    function isVerified(address wallet, uint32 profileId) external returns (bool verified) {}
    function getIdentity(address wallet) external view returns (address) {}
    function getCountry(address wallet) external view returns (uint16) {}
    function contains(address wallet) external view returns (bool) {}

    /*//////////////////////////////////////////////////////////////
                      TRUSTED ISSUERS
    //////////////////////////////////////////////////////////////*/

    function addTrustedIssuer(uint32 profileId, address issuerAddr) external {}
    function removeTrustedIssuer(uint32 profileId, address issuerAddr) external {}
    function isTrustedIssuer(uint32 profileId, address issuerAddr) external view returns (bool) {}

    /*//////////////////////////////////////////////////////////////
                       CLAIM TOPICS
    //////////////////////////////////////////////////////////////*/

    function createProfile(uint256[] calldata claimTopics) external returns (uint32 profileId) {}
    function setProfileClaimTopics(uint32 profileId, uint256[] calldata claimTopics) external {}
    function getProfileClaimTopics(uint32 profileId) external view returns (uint256[] memory) {}
    function getProfileVersion(uint32 profileId) external view returns (uint64) {}
    function profileExists(uint32 profileId) external view returns (bool) {}

    /*//////////////////////////////////////////////////////////////
                     COMPLIANCE ROUTER
    //////////////////////////////////////////////////////////////*/

    function transferred(uint256 tokenId, address from, address to, uint256 amount) external {}
    function minted(uint256 tokenId, address to, uint256 amount) external {}
    function burned(uint256 tokenId, address from, uint256 amount) external {}
    function canTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes calldata data)
        external
        view
        returns (bool ok, bytes32 reason)
    {}
}
