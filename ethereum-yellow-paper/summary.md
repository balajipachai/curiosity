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