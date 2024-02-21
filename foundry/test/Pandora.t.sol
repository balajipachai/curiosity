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

    event WhiteListSet(address indexed target, bool state);

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
        bytes4 invalidOwnerSelector = bytes4(
            keccak256(abi.encodePacked("InvalidOwner()"))
        );
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
        bytes4 unauthorizedSelector = bytes4(
            keccak256(abi.encodePacked("Unauthorized()"))
        );
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
        bytes4 selector = bytes4(
            keccak256(abi.encodePacked("InvalidTarget()"))
        );
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

    function test_Transfer() public {
        //Before transfer alice balance is 0
        assertEq(pandora.balanceOf(alice), 0);

        //Transfer 1 PANDORA to alice
        vm.prank(owner);
        pandora.setWhitelist(owner, true);
        vm.stopPrank();

        vm.prank(owner);
        pandora.transfer(alice, 1 ether);

        //Check alice's balance after transfer
        assertEq(pandora.balanceOf(alice), 1 ether);

        // Transfer from alice to bob
        vm.prank(alice);
        pandora.transfer(bob, 1 ether);
        vm.stopPrank();

        assertEq(pandora.balanceOf(alice), 0);
        assertEq(pandora.balanceOf(bob), 1 ether);
        assertEq(pandora.ownerOf(2), bob);
    }

    function test_TransferFromERC20() public {
        // Bob transfers 1 token to Alice on behalf of the owner account
        vm.prank(owner);
        pandora.setApprovalForAll(bob, true);
        vm.stopPrank();

        vm.prank(owner);
        pandora.approve(bob, 1 ether);
        vm.stopPrank();

        // Check the allowance for Bob is 1 PANDORA token i.e. 1 ether since 1 ether = 10 ** 18 wei
        // Hence instead of using 10 ** 18, we are using 1 ether representing token decimals
        assertEq(pandora.allowance(owner, bob), 1 ether);

        //Should whitelist owner inorder for transferFrom to succeed
        vm.prank(owner);
        pandora.setWhitelist(owner, true);
        vm.stopPrank();

        // This will transfer PANDORA ERC-20 Tokens to Alice
        vm.prank(bob);
        pandora.transferFrom(owner, alice, 0.5 ether);
        vm.stopPrank();

        assertEq(pandora.balanceOf(alice), 0.5 ether); // Alice's balance
        assertEq(pandora.allowance(owner, bob), 0.5 ether); // Bob's allowance
    }

    function test_RevertsTransferFromForInvalidSender() public {
        test_Transfer(); // This will set the value of minted to 2
        bytes4 selector = bytes4(
            keccak256(abi.encodePacked("InvalidSender()"))
        );
        vm.expectRevert(selector);
        pandora.transferFrom(alice, bob, 2);
    }

    function test_RevertsTransferFromWhenToIsAddressZero() public {
        test_Transfer(); // This will set the value of minted to 2
        bytes4 selector = bytes4(
            keccak256(abi.encodePacked("InvalidRecipient()"))
        );
        vm.expectRevert(selector);
        pandora.transferFrom(bob, address(0), 2);
    }

    function test_RevertsTransferFromWhenCallerIsUnauthorized() public {
        test_Transfer(); // This will set the value of minted to 2
        bytes4 selector = bytes4(keccak256(abi.encodePacked("Unauthorized()")));
        vm.expectRevert(selector);
        vm.prank(address(0xbad));
        pandora.transferFrom(bob, alice, 2);
        vm.stopPrank();
    }

    function test_TransferFromERC721() public {
        test_Transfer(); // This will set the value of minted to 2

        // Approval required for if() Unauthorized block to bypass
        vm.prank(bob);
        pandora.setApprovalForAll(address(this), true);
        vm.stopPrank();

        pandora.transferFrom(bob, alice, 2);

        // Check Balance of Bob & Alice
        assertEq(pandora.balanceOf(bob), 0);
        assertEq(pandora.balanceOf(alice), 1 ether);

        // Check owner of id 2 is now Alice
        assertEq(pandora.ownerOf(2), alice);
    }

    /*
    function test_SafeTransferFromERC20() public {
        // Bob transfers 1 token to Alice on behalf of the owner account
        vm.prank(owner);
        pandora.setApprovalForAll(bob, true);
        vm.stopPrank();

        vm.prank(owner);
        pandora.approve(bob, 1000);
        vm.stopPrank();

        //Should whitelist owner inorder for transferFrom to succeed
        vm.prank(owner);
        pandora.setWhitelist(owner, true);
        vm.stopPrank();

        // This will transfer PANDORA ERC-20 Tokens to Alice
        vm.prank(bob);
        pandora.safeTransferFrom(owner, alice, 500);
        vm.stopPrank();

        assertEq(pandora.balanceOf(alice), 500);
    }

    function test_SafeTransferFromERC721() public {
        // Bob transfers 1 token to Alice on behalf of the owner account
        vm.prank(owner);
        pandora.setApprovalForAll(bob, true);
        vm.stopPrank();

        vm.prank(owner);
        pandora.approve(bob, 1);
        vm.stopPrank();

        //Should whitelist owner inorder for transferFrom to succeed
        vm.prank(owner);
        pandora.setWhitelist(owner, true);
        vm.stopPrank();

        // This will transfer PANDORA ERC-20 Tokens to Alice
        vm.prank(bob);
        pandora.safeTransferFrom(owner, alice, 1);
        vm.stopPrank();

        assertEq(pandora.balanceOf(alice), 1);
    }

    function test_RevertsTransferOwnershipWhenNewOwnerIsAddressZero() public {
        bytes4 selector = bytes4(keccak256(abi.encodePacked("InvalidOwner()")));
        vm.expectRevert(selector);
        vm.prank(owner);
        pandora.transferOwnership(address(0));
    }

    function test_RevertTransferOwnershipAsNonOwner() public {
        bytes4 selector = bytes4(keccak256(abi.encodePacked("Unauthorized()")));
        vm.expectRevert(selector);
        pandora.transferOwnership(alice);
    }

    function test_TransferOwnershipWorksAsOwner() public {
        assertEq(pandora.owner(), owner);
        //Transfers ownership to Alice
        vm.prank(owner);
        pandora.transferOwnership(alice);
        vm.stopPrank();

        assertEq(pandora.owner(), alice);
    }

    function test_RevertsRevokeOwnershipAsNonOwner() public {
        bytes4 selector = bytes4(keccak256(abi.encodePacked("Unauthorized()")));
        vm.expectRevert(selector);
        pandora.revokeOwnership();
    }

    function test_RevokeOwnershipWorksAsOwner() public {
        assertEq(pandora.owner(), owner);

        //Revoke ownership
        vm.prank(owner);
        pandora.revokeOwnership();
        vm.stopPrank();

        //Check owner is updated to be address(0)
        assertEq(pandora.owner(), address(0));
    }
    */
}
