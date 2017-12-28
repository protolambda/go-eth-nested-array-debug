pragma solidity ^0.4.18;

contract TestContract {

    uint64[3][4][5] public deepUint64Array;

    function TestContract() public {

    }

    function storeDeepUintArray(uint64[3][4][5] arr) public {
        deepUint64Array = arr;
    }

    function retrieveDeepArray() public view returns (uint64[3][4][5]) {
        return deepUint64Array;
    }

}
