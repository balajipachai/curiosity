// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";

contract CounterScript is Script {
     function test() public {} // To exclude from the coverage report
    function setUp() public {}

    function run() public {
        vm.broadcast();
    }
}
