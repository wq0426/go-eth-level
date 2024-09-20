pragma solidity ^0.8.17

contract CrowdFunding{
   //受益人
   address public immutable beneficiary;
   //筹资目标
   uint256 public immutable fundingGoal;
   //当前的金额
   uint256 public fundingAmount;
   mapping(adress =>uint256) public funders;
   //可迭代的映射
   mapping(address =>bool) private fundersInserted; 
   //length
   address[] public fundersKey; 
   //不用自销毁方法，使用变量来控制
   bool public AVAILABLED = true;//状态
   //部署的时候，写入受益人+筹资目标数量
   constructor(address beneficiary_,uint256 goal_){
      beneficiary = beneficiary_;
      fundingGoal = goal_;
   }

   function contribute() external payable{
      require(AVAILABLED,"CrowdFunding is closed");

      uint256 potentialFundingAmount = fundingAmount + msg.value;
      uint256 refundAmount = 0;

      if(potentialFundingAmount > fundingGoal){
         refundAmount = potentialFundingAmount - fundingGoal;
         funders[msg.sender] +=(msg.value - refundAmount);
         fundingAmount+= (msg.value - refundAmount);
      }else{
         funders[msg.sender] += msg.value;
         fundingAmount += msg.value;
      }
      if(!fundersInserted[msg.sender]){
         fundersInserted[msg.sender] = true;
         fundersKey.push(msg.sender);
      }
      if(refundAmount > 0){
         payable(msg.sender).transfer(refundAmount);
      }
   }
   
   function close() external returns(bool){
      if(fundingAmount<fundingGoal){
         return false;
      }
      uint256 amount = fundingAmount;
      fundingAmount = 0;
      AVAILABLED = false;
      payable(beneficiary).transfer(amount);
      return true;
   }
   function fundersLength() public view returns(uint256){
      return fundersKey.length;
   }
}