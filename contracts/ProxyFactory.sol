pragma solidity >=0.6.0 <0.8.0;

import "./proxy/Proxies.sol";

contract Factory is Proxies {

    // This struct contains the address of the implementation that is
    // used for the proxies it creates.
    ProxyFactory factory;

    constructor(bytes memory _creationCode) {
        // Initializes the factory struct by deploying the implementation
        // and storing its address.
        factory = Proxies.newFactory(_creationCode);
    }

    // initializerData is the encoded call to the initializer function.
    function deployInstance(address admin, bytes memory initializerData) public returns (address) {
        return deploy(factory, admin, initializerData);
    }
}