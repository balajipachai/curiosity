pragma solidity ^0.8.0;

contract ERC191 {
    // The below defined events are not part of ERC191
    event Success(address target, bytes payload);
    event Failure(address target, bytes payload);

    function signatureBasedExecution(address target, uint256 nonce, bytes memory payload, uint8 v, bytes32 r, bytes32 s) public payable {
        
    // Arguments when calculating hash to validate
    // 1: byte(0x19) - the initial 0x19 byte
    // 2: byte(0) - the version byte
    // 3: address(this) - the validator address
    // 4-6 : Application specific data

    bytes32 hash = keccak256(abi.encodePacked(bytes1(0x19), bytes1(0), address(this), msg.value, nonce, payload));

    // recovering the signer from the hash and the signature
    address addressRecovered = ecrecover(hash, v, r, s);
   
    // logic of the wallet
     if (addressRecovered == msg.sender) {
        emit Success(target, payload);
     } else {
        emit Failure(target, payload);
     }
}
}
