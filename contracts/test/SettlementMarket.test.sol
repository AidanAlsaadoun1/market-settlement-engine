// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "forge-std/Test.sol";
import "../src/SettlementMarket.sol";

contract MarketTest is Test {
    Market m;
    address alice = address(0xA11CE);
    address bob   = address(0xB0B);
    address carol = address(0xCA401);

    function setUp() public {
        m = new Market();
        vm.deal(alice, 10 ether);
        vm.deal(bob, 10 ether);
        vm.deal(carol, 10 ether);
    }

    function testAliceBobCarolStory() public {
        uint256 id = m.createMarket("rain saturday?", block.timestamp + 1 days);

        vm.prank(alice);
        m.buy{value: 2 ether}(id, true);

        vm.prank(bob);
        m.buy{value: 1 ether}(id, true);

        vm.prank(carol);
        m.buy{value: 3 ether}(id, false);

        vm.warp(block.timestamp + 1 days + 1);
        m.resolve(id, true);

        // Alice: 2 back + (2 * 3) / 3 = 2 winnings -> balance 10 - 2 + 4 = 12
        vm.prank(alice);
        m.claim(id);
        assertEq(alice.balance, 12 ether);

        // Bob: 1 back + (1 * 3) / 3 = 1 winnings -> 10 - 1 + 2 = 11
        vm.prank(bob);
        m.claim(id);
        assertEq(bob.balance, 11 ether);

        // Carol lost: her claim reverts, balance stays at 10 - 3 = 7
        vm.prank(carol);
        vm.expectRevert("nothing to claim");
        m.claim(id);
        assertEq(carol.balance, 7 ether);

        // Alice tries to double-claim
        vm.prank(alice);
        vm.expectRevert("already claimed");
        m.claim(id);
    }

    function testBuyAfterCloseReverts() public {
        uint256 id = m.createMarket("rain saturday?", block.timestamp + 1 days);
        vm.warp(block.timestamp + 1 days + 1);
        vm.prank(alice);
        vm.expectRevert("Market closed");
        m.buy{value: 1 ether}(id, true);
    }
}