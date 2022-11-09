pragma solidity >=0.4.24 <0.6.11;
pragma experimental ABIEncoderV2;

import "./Table.sol";

contract TableSize {
    event CreateResult(int256 count);
    event InsertResult(int256 count);
    event UpdateResult(int256 count);
    event RemoveResult(int256 count);

    TableFactory tableFactory;
    constructor() public {
        tableFactory = TableFactory(0x1001); //The fixed address is 0x1001 for TableFactory
        // the parameters of createTable are tableName,keyField,"vlaueFiled1,vlaueFiled2,vlaueFiled3,..."
        // tableFactory.createTable(TABLE_NAME, "name", "item_id,item_name");
    }

    //select records
    function size(string memory tablename, string memory name)
    public
    view
    returns (int)
    {
        Table table = tableFactory.openTable(tablename);// open the table

        Condition condition = table.newCondition();

        Entries entries = table.select(name, condition);
        
        return (entries.size());
    }
}
