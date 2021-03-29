pragma solidity ^0.8.1;

/*
Notes

- contract is trusted and usable by only 1 (owner)
- storing strings in solidity is not ideal due to expense; demo purposes only
- storing hashes of ipfs hashes is not a solution

*/

contract Store {
    // state variables
    address public owner;
    
    // array of CIDs
    string[] public cidArray;

    // modifiers
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    // constructor
    constructor() {
        owner = msg.sender;
    }

    // setters

    // function that adds a CID to the stored cidArray
    // only the owner of the contract is allowed to add a CID
    // @param _cid is string of file data (expensive)
    function addCID(string memory _cid) public onlyOwner {
        // push new CID to array
        cidArray.push(_cid);
    }

    // getters
    // programagically created with public state variables

    // function that returns the contents of the entire CID array
    // @return array of CIDs
    function getCIDArray() public view returns (string[] memory) {
        return cidArray;
    }
}
