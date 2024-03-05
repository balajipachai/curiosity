//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "forge-std/console.sol";

import {ERC191} from "../../src/basic-ercs/ERC191.sol";

contract ERC191Test is Test {
    ERC191 erc191;

    event Success(address target, bytes payload);
    event Failure(address target, bytes payload);

    uint256 nonce = 0;
    bytes payload = "0x";
    function setUp() public {
        erc191 = new ERC191();
    }

    /**
    To test ERC191 we have used 2 scenarios
    First, where the messageSender matches the addressRecovered
    Second, where the messageSender does not match the addressRecovered
    Upon match and no match we are emitting Success() & Failure() events
    */
    function test_SuccessfulSignatureBasedExecution() public {
        //Creates and returns an address & privateKey
        (address alice, uint256 alicePk) = makeAddrAndKey("alice");

        uint256 value = 0.1 ether;
        // The hash that will be signed
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0x19), bytes1(0), address(erc191), value, nonce, payload));
        // Use vm.sign to get the values of v, r & s
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(alicePk, hash);
        address target = address(erc191);

        // Since msg.sender == addressRecovered in ERC191 thus the prank
        vm.prank(alice);
        // Send 1 ether to alice, s.t. erc191.signatureBasedExecution(target, nonce, payload, v, r, s); call does not fail
        vm.deal(alice, 1 ether);
        vm.expectEmit(); // checks all event data
        emit Success(target, payload);
        erc191.signatureBasedExecution{value: 0.1 ether}(target, nonce, payload, v, r, s);
        vm.stopPrank();

        address signer = ecrecover(hash, v, r, s);

        assertEq(alice, signer);
    }

    function test_FailureSignatureBasedExecution() public {
        // Creates and returns an address & a privateKey
        (address alice, uint256 alicePk) = makeAddrAndKey("alice");
        address bob = address(0x1234);

        uint256 value = 0.1 ether;
        // The message that will be signed
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0x19), bytes1(0), address(erc191), value, nonce, payload));
        // Use vm.sign to get the values of v, r, & s
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(alicePk, hash);

        address target = address(this);

        // For the Failure event to be emitted
        // addressRecovered != msg.sender, thus, vm.prank via bob's address and not alice's address
        // as Alice is signing the message thus ecrecover returns Alice's public address
        // Since we are pranking with Bob's address hence
        // addressRecovered != msg.sender and thus Failure event is emitted
        vm.prank(bob);
        vm.deal(bob, 1 ether);
        vm.expectEmit(); // checks all event data
        emit Failure(target, payload);
        erc191.signatureBasedExecution{value: 0.1 ether}(target, nonce, payload, v, r, s);
        vm.stopPrank();

        address signer = ecrecover(hash, v, r, s);

        assertEq(alice, signer);
    }
}