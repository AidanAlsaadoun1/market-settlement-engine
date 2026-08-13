// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.20;


contract Market {
    address public owner;

    constructor() {
        owner = msg.sender;
    }
    
    struct MarketData {
        string question;
        uint256 closeTime;
        bool resolved;
        bool outcome;
        uint256 yesPool;
        uint256 noPool;
    }

    mapping(uint256 => MarketData) public markets;
    mapping(uint256 => mapping(address => uint256)) public yesStakes;
    mapping(uint256 => mapping(address => uint256)) public noStakes;
    mapping(uint256 => mapping(address => bool)) public claimed;

    event MarketCreated(uint256 indexed id, string question, uint256 closeTime);
    event SharesPurchased(uint256 indexed marketId, address indexed buyer, bool yes, uint256 amount);
    event MarketResolved(uint256 indexed marketId, bool outcome);
    event Claimed(uint256 indexed marketId, address indexed claimer, uint256 payout);

    uint256 public marketCount;

    // Create a market (validation checks depending) and emits the event
    function createMarket(string calldata question, uint256 closeTime) external returns (uint256 id)
    {
        require(closeTime > block.timestamp, "close time in past");
        id = marketCount++;
        markets[id] = MarketData(question, closeTime, false, false, 0, 0);
        emit MarketCreated(id, question, closeTime);
    }

    // byuy function to buy into a market with require to reverse transations if conditions aren't met
    function buy(uint256 marketId, bool yes) external payable {
        require(marketId < marketCount, "Market does not exist");
        require(!markets[marketId].resolved, "Market already resolved");
        require(block.timestamp <= markets[marketId].closeTime, "Market closed");
        require(msg.value > 0, "No value received");

        if (yes) 
        {
            markets[marketId].yesPool += msg.value;
            yesStakes[marketId][msg.sender] += msg.value;
        }
        else
        {
            markets[marketId].noPool += msg.value;
            noStakes[marketId][msg.sender] += msg.value;
        }

        emit SharesPurchased(marketId, msg.sender, yes, msg.value);
    }
    
    function resolve(uint256 marketId, bool outcome) external {
        require(msg.sender == owner);
        require(marketId < marketCount, "Market does not exist");
        require(!markets[marketId].resolved, "Market already resolved");
        require(block.timestamp >= markets[marketId].closeTime, "Market closed");
        
        markets[marketId].resolved = true;
        markets[marketId].outcome = outcome;

        emit MarketResolved(marketId, outcome);
    }

function claim(uint256 marketId) external {
    MarketData storage m = markets[marketId];

    require(m.resolved, "market not resolved");
    require(!claimed[marketId][msg.sender], "already claimed");

    uint256 stake;      
    uint256 winningPool;
    uint256 losingPool; 

    if (m.outcome) {
        // The answer was YES -> yes-bucket people win.
        stake = yesStakes[marketId][msg.sender];
        winningPool = m.yesPool;
        losingPool = m.noPool;
    } else {
        // The answer was NO -> no-bucket people win.
        stake = noStakes[marketId][msg.sender];
        winningPool = m.noPool;
        losingPool = m.yesPool;
    }

    require(stake > 0, "nothing to claim");

    uint256 payout = stake + (stake * losingPool) / winningPool;

    // Mark as paid BEFORE sending money = reentrancy protection
    // (checks-effects-interactions). If we transferred first, a hostile
    // contract could re-enter claim() and drain the pool.
    claimed[marketId][msg.sender] = true;

    // Actually send the ETH.
    (bool ok, ) = payable(msg.sender).call{value: payout}("");
    require(ok, "Transfer failed");

    emit Claimed(marketId, msg.sender, payout);
}
}
