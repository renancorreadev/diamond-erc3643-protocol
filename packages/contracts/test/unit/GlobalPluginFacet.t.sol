// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {
    GlobalPluginFacet,
    GlobalPluginFacet__Unauthorized,
    GlobalPluginFacet__ZeroAddress,
    GlobalPluginFacet__AlreadyRegistered,
    GlobalPluginFacet__NotRegistered,
    GlobalPluginFacet__TooManyPlugins,
    GlobalPluginFacet__AlreadyActive,
    GlobalPluginFacet__AlreadyInactive
} from "../../src/facets/plugins/GlobalPluginFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IPluginModule} from "../../src/interfaces/plugins/IPluginModule.sol";
import {GlobalPluginInfo} from "../../src/storage/LibGlobalPluginStorage.sol";

/// @dev Minimal mock that implements IPluginModule
contract MockGlobalPlugin is IPluginModule {
    string internal _name;

    constructor(string memory name_) {
        _name = name_;
    }

    function name() external view returns (string memory) {
        return _name;
    }
}

/// @dev Contract that does NOT implement IPluginModule (no name())
contract NotAPlugin {
    function doSomething() external pure returns (uint256) {
        return 42;
    }
}

contract GlobalPluginFacetTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal admin = makeAddr("admin");
    address internal stranger = makeAddr("stranger");

    GlobalPluginFacet internal gp;
    AccessControlFacet internal ac;

    bytes32 internal constant COMPLIANCE_ADMIN = keccak256("COMPLIANCE_ADMIN");

    MockGlobalPlugin internal marketplace;
    MockGlobalPlugin internal amm;
    MockGlobalPlugin internal governance;

    function setUp() public {
        d = deployDiamond(owner);
        gp = GlobalPluginFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));

        marketplace = new MockGlobalPlugin("Marketplace v1");
        amm = new MockGlobalPlugin("AMM");
        governance = new MockGlobalPlugin("Governance");

        vm.prank(owner);
        ac.grantRole(COMPLIANCE_ADMIN, admin);
    }

    /*//////////////////////////////////////////////////////////////
                        REGISTER — HAPPY PATH
    //////////////////////////////////////////////////////////////*/

    function test_RegisterGlobalPlugin_AsOwner() public {
        vm.prank(owner);
        gp.registerGlobalPlugin(address(marketplace));

        assertTrue(gp.isGlobalPlugin(address(marketplace)));
        assertEq(gp.globalPluginCount(), 1);
    }

    function test_RegisterGlobalPlugin_AsAdmin() public {
        vm.prank(admin);
        gp.registerGlobalPlugin(address(marketplace));

        assertTrue(gp.isGlobalPlugin(address(marketplace)));
    }

    function test_RegisterGlobalPlugin_EmitsEvent() public {
        vm.expectEmit(true, false, false, true);
        emit GlobalPluginFacet.GlobalPluginRegistered(address(marketplace), "Marketplace v1");

        vm.prank(owner);
        gp.registerGlobalPlugin(address(marketplace));
    }

    function test_RegisterGlobalPlugin_SetsActiveByDefault() public {
        vm.prank(owner);
        gp.registerGlobalPlugin(address(marketplace));

        GlobalPluginInfo memory info = gp.getGlobalPluginInfo(address(marketplace));
        assertTrue(info.active);
        assertEq(info.plugin, address(marketplace));
        assertEq(info.registeredAt, uint64(block.timestamp));
    }

    function test_RegisterMultiplePlugins() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.registerGlobalPlugin(address(amm));
        gp.registerGlobalPlugin(address(governance));
        vm.stopPrank();

        assertEq(gp.globalPluginCount(), 3);

        address[] memory active = gp.getActiveGlobalPlugins();
        assertEq(active.length, 3);
    }

    /*//////////////////////////////////////////////////////////////
                        REGISTER — REVERTS
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_RegisterGlobalPlugin_Unauthorized() public {
        vm.expectRevert(GlobalPluginFacet__Unauthorized.selector);
        vm.prank(stranger);
        gp.registerGlobalPlugin(address(marketplace));
    }

    function test_RevertWhen_RegisterGlobalPlugin_ZeroAddress() public {
        vm.expectRevert(GlobalPluginFacet__ZeroAddress.selector);
        vm.prank(owner);
        gp.registerGlobalPlugin(address(0));
    }

    function test_RevertWhen_RegisterGlobalPlugin_AlreadyRegistered() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));

        vm.expectRevert(
            abi.encodeWithSelector(GlobalPluginFacet__AlreadyRegistered.selector, address(marketplace))
        );
        gp.registerGlobalPlugin(address(marketplace));
        vm.stopPrank();
    }

    function test_RevertWhen_RegisterGlobalPlugin_NotIPluginModule() public {
        NotAPlugin bad = new NotAPlugin();
        vm.expectRevert(); // call to name() reverts
        vm.prank(owner);
        gp.registerGlobalPlugin(address(bad));
    }

    function test_RevertWhen_RegisterGlobalPlugin_TooMany() public {
        vm.startPrank(owner);
        for (uint256 i; i < 20;) {
            MockGlobalPlugin p = new MockGlobalPlugin("Plugin");
            gp.registerGlobalPlugin(address(p));
            unchecked { ++i; }
        }

        MockGlobalPlugin extra = new MockGlobalPlugin("Extra");
        vm.expectRevert(abi.encodeWithSelector(GlobalPluginFacet__TooManyPlugins.selector, 21, 20));
        gp.registerGlobalPlugin(address(extra));
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                        REMOVE — HAPPY PATH
    //////////////////////////////////////////////////////////////*/

    function test_RemoveGlobalPlugin() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.removeGlobalPlugin(address(marketplace));
        vm.stopPrank();

        assertFalse(gp.isGlobalPlugin(address(marketplace)));
        assertEq(gp.globalPluginCount(), 0);
    }

    function test_RemoveGlobalPlugin_EmitsEvent() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));

        vm.expectEmit(true, false, false, false);
        emit GlobalPluginFacet.GlobalPluginRemoved(address(marketplace));

        gp.removeGlobalPlugin(address(marketplace));
        vm.stopPrank();
    }

    function test_RemoveGlobalPlugin_SwapAndPop() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.registerGlobalPlugin(address(amm));
        gp.registerGlobalPlugin(address(governance));

        // Remove middle element — amm should be replaced by governance
        gp.removeGlobalPlugin(address(amm));
        vm.stopPrank();

        assertEq(gp.globalPluginCount(), 2);
        assertFalse(gp.isGlobalPlugin(address(amm)));
        assertTrue(gp.isGlobalPlugin(address(marketplace)));
        assertTrue(gp.isGlobalPlugin(address(governance)));
    }

    function test_RemoveGlobalPlugin_LastElement() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.registerGlobalPlugin(address(amm));

        gp.removeGlobalPlugin(address(amm));
        vm.stopPrank();

        assertEq(gp.globalPluginCount(), 1);
        assertTrue(gp.isGlobalPlugin(address(marketplace)));
    }

    function test_RemoveAndReRegister() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.removeGlobalPlugin(address(marketplace));
        gp.registerGlobalPlugin(address(marketplace));
        vm.stopPrank();

        assertTrue(gp.isGlobalPlugin(address(marketplace)));
        assertEq(gp.globalPluginCount(), 1);
    }

    /*//////////////////////////////////////////////////////////////
                        REMOVE — REVERTS
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_RemoveGlobalPlugin_Unauthorized() public {
        vm.prank(owner);
        gp.registerGlobalPlugin(address(marketplace));

        vm.expectRevert(GlobalPluginFacet__Unauthorized.selector);
        vm.prank(stranger);
        gp.removeGlobalPlugin(address(marketplace));
    }

    function test_RevertWhen_RemoveGlobalPlugin_NotRegistered() public {
        vm.expectRevert(
            abi.encodeWithSelector(GlobalPluginFacet__NotRegistered.selector, address(marketplace))
        );
        vm.prank(owner);
        gp.removeGlobalPlugin(address(marketplace));
    }

    /*//////////////////////////////////////////////////////////////
                        STATUS — HAPPY PATH
    //////////////////////////////////////////////////////////////*/

    function test_SetGlobalPluginStatus_Deactivate() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.setGlobalPluginStatus(address(marketplace), false);
        vm.stopPrank();

        GlobalPluginInfo memory info = gp.getGlobalPluginInfo(address(marketplace));
        assertFalse(info.active);
    }

    function test_SetGlobalPluginStatus_Reactivate() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.setGlobalPluginStatus(address(marketplace), false);
        gp.setGlobalPluginStatus(address(marketplace), true);
        vm.stopPrank();

        GlobalPluginInfo memory info = gp.getGlobalPluginInfo(address(marketplace));
        assertTrue(info.active);
    }

    function test_SetGlobalPluginStatus_EmitsEvent() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));

        vm.expectEmit(true, false, false, true);
        emit GlobalPluginFacet.GlobalPluginStatusChanged(address(marketplace), false);

        gp.setGlobalPluginStatus(address(marketplace), false);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                        STATUS — REVERTS
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_SetStatus_Unauthorized() public {
        vm.prank(owner);
        gp.registerGlobalPlugin(address(marketplace));

        vm.expectRevert(GlobalPluginFacet__Unauthorized.selector);
        vm.prank(stranger);
        gp.setGlobalPluginStatus(address(marketplace), false);
    }

    function test_RevertWhen_SetStatus_NotRegistered() public {
        vm.expectRevert(
            abi.encodeWithSelector(GlobalPluginFacet__NotRegistered.selector, address(marketplace))
        );
        vm.prank(owner);
        gp.setGlobalPluginStatus(address(marketplace), false);
    }

    function test_RevertWhen_SetStatus_AlreadyActive() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));

        vm.expectRevert(
            abi.encodeWithSelector(GlobalPluginFacet__AlreadyActive.selector, address(marketplace))
        );
        gp.setGlobalPluginStatus(address(marketplace), true);
        vm.stopPrank();
    }

    function test_RevertWhen_SetStatus_AlreadyInactive() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.setGlobalPluginStatus(address(marketplace), false);

        vm.expectRevert(
            abi.encodeWithSelector(GlobalPluginFacet__AlreadyInactive.selector, address(marketplace))
        );
        gp.setGlobalPluginStatus(address(marketplace), false);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                            VIEWS
    //////////////////////////////////////////////////////////////*/

    function test_GetGlobalPlugins_Empty() public view {
        GlobalPluginInfo[] memory plugins = gp.getGlobalPlugins();
        assertEq(plugins.length, 0);
    }

    function test_GetActiveGlobalPlugins_FiltersInactive() public {
        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(marketplace));
        gp.registerGlobalPlugin(address(amm));
        gp.registerGlobalPlugin(address(governance));

        gp.setGlobalPluginStatus(address(amm), false);
        vm.stopPrank();

        address[] memory active = gp.getActiveGlobalPlugins();
        assertEq(active.length, 2);
    }

    function test_IsGlobalPlugin_ReturnsFalseForUnregistered() public view {
        assertFalse(gp.isGlobalPlugin(address(marketplace)));
    }

    function test_GlobalPluginCount_Empty() public view {
        assertEq(gp.globalPluginCount(), 0);
    }

    function test_RevertWhen_GetGlobalPluginInfo_NotRegistered() public {
        vm.expectRevert(
            abi.encodeWithSelector(GlobalPluginFacet__NotRegistered.selector, address(marketplace))
        );
        gp.getGlobalPluginInfo(address(marketplace));
    }

    /*//////////////////////////////////////////////////////////////
                        GAS — VERSIONING SCENARIO
    //////////////////////////////////////////////////////////////*/

    function test_VersionUpgrade_RemoveV1_AddV2() public {
        MockGlobalPlugin v1 = new MockGlobalPlugin("Marketplace v1");
        MockGlobalPlugin v2 = new MockGlobalPlugin("Marketplace v2");

        vm.startPrank(owner);
        gp.registerGlobalPlugin(address(v1));
        assertTrue(gp.isGlobalPlugin(address(v1)));

        gp.removeGlobalPlugin(address(v1));
        gp.registerGlobalPlugin(address(v2));
        vm.stopPrank();

        assertFalse(gp.isGlobalPlugin(address(v1)));
        assertTrue(gp.isGlobalPlugin(address(v2)));
        assertEq(gp.globalPluginCount(), 1);

        GlobalPluginInfo memory info = gp.getGlobalPluginInfo(address(v2));
        assertEq(info.plugin, address(v2));
    }
}
