pragma solidity ^0.8.1;

import "./Store.sol";

contract TestStore {
    // string of CID to store
    string public cid;

    // declare boolean to store result comparison
    bool public resultBool;

    // alias for readability of contract instance
    Store public store;

    // create new instance of Store before each test run
    function beforeEach() public {
        store = new Store();
    }

    // test add CID to contract
    function testAddCID() public {
        // sample CID value
        // using a test string instead of of an IPFS hash
        cid = "QmdUWc45kg8QhuG5e86rVisaGHUMcVSG9DqkGhd8pq1HdK";

        // call contract method
        // convert bytes to string as argument
        store.addCID(cid);

        // test retrieving the CID (index 0)
        string memory result = store.cidArray(0);

        // test values stored/retrieved are equal
        // compare values easily by converting string to bytes then hashing
        if (keccak256(bytes(result)) == keccak256(bytes(cid))) {
            resultBool = true;
        }

        assert(resultBool);
    }

    // test return CIDs from contract
    function testReturnCIDs() public view returns (string[] memory) {
        // get array of CIDs from contract
        string[] memory result = store.getCIDArray();

        return result;
    }
}
