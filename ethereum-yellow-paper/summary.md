# Trie

- When it comes to `Trie` datastructure, the important thing is the definition of the `TrieNode` structure.
- After defining the structure, give some time to go over it and think about the `Insert`, `Delete` & `Search` functionalities.
- Take a pen and paper and work around the CRUD operations for a Trie.
- Once the above step is done, then sit for coding, and you  will be amazed about your understanding of the `Trie` datastructure and at the same time be wondered about the beauty of datastructures.
- If you understand the basics and fundamentals then programming becomes fun and an act of magic.

`Go through the comments added above the function definitions in trie.go to understand the implementations.`

# Recursive Length Prefix (RLP) encoding/decoding

## RLP encoding is defined as follows:

1. If a string is a single byte whose value is in the range `0x00 - 0x7f [0 - 127]` its RLP encoding is it's own value.
2. If a string is `0-55 bytes long`, it's RLP encoding is defined as follows:

    2.1 A single byte with value `0x80 [128]` plus the length of the string followed by the string itself.

3. If a string is `> 55 bytes long`, its RLP encoding is defined as follows:

    3.1 A single byte with value `0xb7 [183]` plus the length of the string in binary form, followed by the length of the string followed by the string.

4. If total payload of a list (i.e. combined length of all its items being RLP encoded) is `0-55 bytes long`, thne its RLP encoding is defined as follows:

    4.1 A single byte with value `0xc0` plus the length of the list followed by the cancatenation of the RLP encoding of the items.

5. If total payload of a list (i.e. combined length of all its items being RLP encoded) is `> 55 bytes long`, thne its RLP encoding is defined as follows:

    5.1 A single byte with value `0xf7` plus the length in bytes of the length of the payload in binary form, followed by the length of the payload followed by the concatenation of the RLP encodings of the items.

6. First byte in RLP encoding & what it means:

    - [0x00 .. 0x7f] : Data is of type String and should be decoded as it is
    - [0x80 .. 0xb7] : String and its a short string
    - [0xb8 .. 0xbf ] :String and its a long string
    - [0xc0 .. 0xf7] : List and short list
    - [0xf8 .. 0xff] : List and it’s a long list




References:

- [RECURSIVE-LENGTH PREFIX (RLP) SERIALIZATION](https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp/)
- [Data structure in Ethereum | Episode 1: Recursive Length Prefix (RLP) Encoding/Decoding.](https://medium.com/coinmonks/data-structure-in-ethereum-episode-1-recursive-length-prefix-rlp-encoding-decoding-d1016832f919)
- [Ethereum Under The Hood Part 3 (RLP Decoding)](https://medium.com/coinmonks/ethereum-under-the-hood-part-3-rlp-decoding-df236dc13e58)

# EVM Opcodes & Ethereum Virtual Machine

1. There are in all 256 Ethereum opcodes however not all of them exist at the moment.
2. Some of the opcodes are not defined and are left for future implmentation.
3. Ethereum Virtual Machine is a stack-based big-endian system.
4. Stack based implies that each instruction is put on top of the stack and its execution happends by popping those instructions off the stack.
5. Big Endian: The most significant byte MSB is stored at the lowest memory address and the Least Significant Byte LSB is stored at the largest memory address.
6. What is the diffrence between a normal and contract creation transaction?

    `Answer`: The receiver or to of contract creation transaction is a `null` address i.e. to = address(0), this informs the EVM that the current transaction is a contract creation transaction. Interesting thing is after contract deployment the bytecode doesn't contain the constructor code, it is run only during contract deployment, which is logical too that after contract deployment anyways we are not going to call the constructor hence it is a good design decision to not include constructor in the deployed contract bytecode.

7. How EVM knows that `transfer()` of `DAI` contract is called?

    `Answer`: Having `to=address(DAI)` informs the EVM that the transaction is headed towards the DAI contract, however, inside `DAI` there are multiple functions then how does the EVM decides to call which function, this is possible by the method selector which is `bytes4(keccak256(transfer(address to, uint256 amount)))`. If you inspect the transaction data on etherscan the first four bytes are always the method selector or function selector. Thus, in this way the EVM determines which function to call.

    - Another simpler analogy to understand above concept is, you live in `Burj Khalifa` (This is the contract address) and your flat number is `567`(This is the method selector)

8. Storage in EVM is a persistent associative map having `uint256s as both keys and values` i.e. mapping(uint256 => uint256). Or you can think of it as a JSON object having keys and values both to be `uint256s`. All contract state variables are stored in storage.

9. How does EVM decides how much gas a transaction will use?

    `Answer`: Each transaction is sent as bytecode to the EVM for execution. Bytecode consists of Opcodes and gas requirement of each opcode is known in advance which is then used to calculate the amount of gas a transaction will take.

10. [EVM OPCODES DEFINITION](https://github.com/ethereum/go-ethereum/blob/master/core/vm/instructions.go#L27)

11. [Solidity Block & Transaction Properties](https://docs.soliditylang.org/en/develop/units-and-global-variables.html#block-and-transaction-properties)

12. [Layout of state variables in storage](https://docs.soliditylang.org/en/latest/internals/layout_in_storage.html)

