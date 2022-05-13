package basic

import (
	"fmt"
	"testing"
)

func TestHex2Byte(t *testing.T) {
	hex := "0x840" // #define MAXVA (1L << (9 + 9 + 9 + 12 - 1))
	fmt.Printf("%s = %d Byte\n", hex, Hex2Byte(hex))
	fmt.Printf("%s = %d KB\n", hex, Hex2KByte(hex))
	fmt.Printf("%s = %d MB\n", hex, Hex2MByte(hex))
}

func TestDec2Hex(t *testing.T) {
	_ = Dec2Hex(0x80000000 + 128*1024*1024)
	_ = Dec2Hex(100 * 1024 * 1024)
}
