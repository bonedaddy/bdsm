pragma solidity >=0.6.0 <0.8.0;

import "./ProxyFactory.sol";
import "./StructAssignTest.sol";

contract TestProxyFactory is Factory {
    uint256 public numProxies;
    address[] public proxies;
    event ProxyCreated(address _proxy);

    constructor() Factory(type(StructAssignTest).creationCode) {}

    // initializerData is the encoded call to the initializer function.
    function deployProxyContract(address admin, bytes memory initializerData) public returns (address) {
        address prox = deploy(factory, admin, initializerData);
        emit ProxyCreated(prox);
        proxies.push(prox);
        numProxies += 1;
    }
}