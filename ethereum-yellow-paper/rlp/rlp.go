package main

import (
	"bytes"
	"fmt"
	"math/big"
)

func EncodeRLP(input interface{}) ([]byte, error) {
	switch data := input.(type) {
	case []byte:
		if len(data) == 1 && data[0] < 128 {
			return data, nil
		}
		return encodeBytes(data)

	case [][]byte:
		encodedItems := make([][]byte, len(data))
		for i, item := range data {
			encodedItem, err := EncodeRLP(item)
			if err != nil {
				return nil, err
			}
			encodedItems[i] = encodedItem
		}
		return encodeList(encodedItems)
	default:
		return nil, fmt.Errorf("Type not supported for conversion: %T, data")
	}

}

func encodeBytes(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return []byte{0x00}, nil
	}
	lengthPrefix := encodeLengthPrefix(len(data), 128)
	return bytes.Join([][]byte{lengthPrefix, data}, nil), nil
}

func encodeList(items [][]byte) ([]byte, error) {
	encodedItems := make([][]byte, len(items))
	for i, item := range items {
		encodedItem, err := EncodeRLP(item)
		if err != nil {
			return nil, err
		}
		encodedItems[i] = encodedItem
	}
	lengthPrefix := encodeLengthPrefix(len(bytes.Join(encodedItems, nil)), 192)
	return bytes.Join(append([][]byte{lengthPrefix}, encodedItems...), nil), nil
}

func encodeLengthPrefix(length, offset int) []byte {
	if length < 56 {
		return []byte{byte(length + offset)}
	}
	lengthBytes := []byte(fmt.Sprintf("%x", length))
	lengthBytesLen := len(lengthBytes)
	return bytes.Join([][]byte{{byte(lengthBytesLen + offset + 55)}}, append([]byte{0x00}, lengthBytes...))
}

func DecodeRLP(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("Empty input provided")
	}

	// Check the first byte to determine the encoding type
	switch {
	case data[0] <= 0x7f:
		// If the first byte is in the range [0x00, 0x7f], it represents a single byte
		return data[0], nil

	case data[0] <= 0xb7:
		// If the first byte is in the range [0x80, 0xb7], it represents a string with a length of 0-55 bytes
		length := int(data[0] - 0x80)
		if len(data) < 1+length {
			return nil, fmt.Errorf("insufficient data")
		}
		return data[1 : 1+length], nil

	case data[0] <= 0xbf:
		// If the first byte is in the range [0xb8, 0xbf], it represents a string with a length greater than 55 bytes
		lengthBytes := int(data[0] - 0xb7)
		if len(data) < 1+lengthBytes {
			return nil, fmt.Errorf("insufficient data")
		}
		length := decodeRLPSize(data[1 : 1+lengthBytes])
		if len(data) < 1+lengthBytes+length {
			return nil, fmt.Errorf("insufficient data")
		}
		return data[1+lengthBytes : 1+lengthBytes+length], nil

	case data[0] <= 0xf7:
		// If the first byte is in the range [0xc0, 0xf7], it represents a list with a length of 0-55 items
		length := int(data[0] - 0xc0)
		items := make([]interface{}, length)
		offset := 1
		for i := 0; i < length; i++ {
			item, err := decodeRLP(data[offset:])
			if err != nil {
				return nil, err
			}
			items[i] = item
			offset += rlpEncodedSize(data[offset:])
		}
		return items, nil

	case data[0] <= 0xff:
		// If the first byte is in the range [0xf8, 0xff], it represents a list with a length greater than 55 items
		lengthBytes := int(data[0] - 0xf7)
		if len(data) < 1+lengthBytes {
			return nil, fmt.Errorf("insufficient data")
		}
		length := decodeRLPSize(data[1 : 1+lengthBytes])
		if len(data) < 1+lengthBytes+length {
			return nil, fmt.Errorf("insufficient data")
		}
		items := make([]interface{}, length)
		offset := 1 + lengthBytes
		for i := 0; i < length; i++ {
			item, err := decodeRLP(data[offset:])
			if err != nil {
				return nil, err
			}
			items[i] = item
			offset += rlpEncodedSize(data[offset:])
		}
		return items, nil

	default:
		return nil, fmt.Errorf("unsupported encoding")
	}
}

func decodeRLPSize(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 && data[0] <= 0x7f {
		return int(data[0])
	}
	return decodeBigInt(data).Int64()
}

func decodeBigInt(data []byte) *big.Int {
	if len(data) == 0 {
		return big.NewInt(0)
	}
	return new(big.Int).SetBytes(data)
}

func main() {
	input := "balaji"
	inputInBytes := []byte(input)

	res, err := EncodeRLP(inputInBytes)

	if err != nil {
		fmt.Println("Error is: ", err)
	}
	fmt.Printf("Result is: %x", res)
}
