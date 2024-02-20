//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../src/Pandora.sol";

contract PandoraTest is Test {
    // Events
    event NameSymbolUpdate(
        string oldName,
        string oldSymbol,
        string newName,
        string newSymbol
    );

    event WhiteListSet(
        address indexed target,
        bool state
    );

    Pandora pandora;
    address owner = address(1);
    address alice = address(0x1234);
    address bob = address(0x5678);


    function setUp() public {
        pandora = new Pandora(owner);
        vm.prank(owner);
        // approve current contract to spend tokens on behalf of owner
        pandora.approve(address(this), type(uint256).max);
        vm.stopPrank();
    }

    function test_ConstructorInvocation() public {
        uint256 total = 10000 * 10 ** 18;
        assertEq(pandora.name(), "Pandora");
        assertEq(pandora.symbol(), "PANDORA");
        assertEq(pandora.decimals(), 18);
        assertEq(pandora.totalSupply(), total);
        assertEq(pandora.balanceOf(owner), total);
    }

    function test_RevertDeployWhenOwnerIsZeroAddress() public {
        bytes4 invalidOwnerSelector = bytes4(keccak256(abi.encodePacked("InvalidOwner()")));
        vm.expectRevert(invalidOwnerSelector);
        pandora = new Pandora(address(0));
    }

    function test_SetNameAndSymbolWorksAsOwner() public {
        string memory oldName = pandora.name();
        string memory oldSymbol = pandora.symbol();
        assertEq(oldName, "Pandora");
        assertEq(oldSymbol, "PANDORA");

        // Update name & symbol
        vm.prank(owner);
        string memory newName = "Balaji";
        string memory newSymbol = "BALAJI";
        vm.expectEmit();
        emit NameSymbolUpdate(oldName, oldSymbol, newName, newSymbol);
        pandora.setNameSymbol(newName, newSymbol);
        vm.stopPrank();

        // Check if new name & symbol are set
        assertEq(pandora.name(), newName);
        assertEq(pandora.symbol(), newSymbol);

        vm.prank(owner);
        vm.expectEmit();
        emit NameSymbolUpdate(newName, newSymbol, oldName, oldSymbol);
        // Set the name & symbol to the previous old ones
        pandora.setNameSymbol(oldName, oldSymbol);
        assertEq(pandora.name(), oldName);
        assertEq(pandora.symbol(), oldSymbol);
        vm.stopPrank();
    }

    function test_SetNameAndSymbolFailAsNonOwner() public {
        bytes4 unauthorizedSelector = bytes4(keccak256(abi.encodePacked("Unauthorized()")));
        vm.expectRevert(unauthorizedSelector);
        pandora.setNameSymbol("Dai StableCoin", "DAI");
    }

    function test_SetWhitelistWorksAsOwner() public {
        address toWhitelist = address(2);
        // Before calling whitelist whitelist[address[2]] should return false
        assertEq(pandora.whitelist(toWhitelist), false);

        // Now whitelist the address(2)
        vm.prank(owner);
        pandora.setWhitelist(toWhitelist, true);
        vm.stopPrank();

        // After calling whitelist whitelist[address[2]] should return true
        assertEq(pandora.whitelist(toWhitelist), true);
    }

    function test_RevertSetWhitelistWhenTargetIsAddressZero() public {
        bytes4 selector = bytes4(keccak256(abi.encodePacked("InvalidTarget()")));
        vm.expectRevert(selector);
        vm.prank(owner);
        pandora.setWhitelist(address(0), true);
        vm.stopPrank();
    }

    function test_SetWhitelistEmitsWhiteListSet() public {
        vm.prank(owner);
        vm.expectEmit(); // checks all event data
        emit WhiteListSet(address(2), false);
        pandora.setWhitelist(address(2), false);
        vm.stopPrank();
    }

    function test_RevertOwnerOfForNonExistingToken() public {
        bytes4 selector = bytes4(keccak256(abi.encodePacked("NotFound()")));
        vm.expectRevert(selector);
        vm.prank(owner);
        pandora.ownerOf(1);
        vm.stopPrank();
    }


}