contract Demo{
   struct  Todo {
      string name;
      bool isCompleted;
   }
   Todo[] public list; 

   function create(string memory name_) external {
      list.push(
         Todo({
            name:name_,
            isCompleted:false
         })
      );
   }

   function modiName1(uint256 index_,string memory name_) external{
      list[index_].name = name_;
   }
   function modiName2(uint256 index_,string memory name_) external {
       Todo storage temp = list[index_];
       temp.name = name_;
   }
   function modiStatus1(uint256 index_,bool status_) external {
      list[index_].isCompleted = status_;
   }
   function modiStatus2(uint256 index_) external {
      list[index_].isCompleted = !list[index_].isCompleted;
   }
   function get1(uint256 index_) external view returns(string memory name_,bool status_){
        Todo memory temp = list[index_];
        return(temp.name,temp.isCompleted);
   }
   function get2(uint256 index_) external view returns(string memory name_,bool status_){
       Todo storage temp = list[index_];
       return(temp.name,temp.isCompleted);
   }
}