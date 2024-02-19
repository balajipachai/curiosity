//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/Pandora.sol";

contract PandoraTest is Test {
    Pandora pandora;
    address owner = address(1);
    function setUp() public {
        pandora = new Pandora(owner);
    }

    function test_ConstructorInvocation() public {
        uint256 total = 10000 * 10 ** 18;
        assertEq(pandora.name(), "Pandora");
        assertEq(pandora.symbol(), "PANDORA");
        assertEq(pandora.decimals(), 18);
        assertEq(pandora.totalSupply(), total);
        assertEq(pandora.balanceOf(owner), total);
    }
}