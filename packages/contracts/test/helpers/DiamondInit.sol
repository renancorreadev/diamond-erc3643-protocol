// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibDiamond, DiamondStorage} from "../../src/libraries/LibDiamond.sol";
import {IDiamondCut, IDiamondLoupe} from "../../src/interfaces/IDiamond.sol";
// solhint-disable-next-line import-path-check
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/// @dev Called via delegatecall during diamondCut to register ERC165 interfaces.
contract DiamondInit {
    function init() external {
        DiamondStorage storage ds = LibDiamond.diamondStorage();
        ds.supportedInterfaces[type(IERC165).interfaceId] = true;
        ds.supportedInterfaces[type(IDiamondCut).interfaceId] = true;
        ds.supportedInterfaces[type(IDiamondLoupe).interfaceId] = true;
    }
}
