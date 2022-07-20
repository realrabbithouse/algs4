package basic

import (
	"fmt"
	"strings"
)

func Hex2Byte(hex string) int {
	var x int
	reader := strings.NewReader(hex)
	_, err := fmt.Fscanf(reader, "%v", &x)
	if err != nil {
		fmt.Println("Fscanf err:", err)
	}
	return x
}

func Hex2KByte(hex string) int {
	var x int
	reader := strings.NewReader(hex)
	_, err := fmt.Fscanf(reader, "%v", &x)
	if err != nil {
		fmt.Println("Fscanf err:", err)
	}
	return x / 1024
}

func Hex2MByte(hex string) int {
	var x int
	reader := strings.NewReader(hex)
	_, err := fmt.Fscanf(reader, "%v", &x)
	if err != nil {
		fmt.Println("Fscanf err:", err)
	}
	return x / (1024 * 1024)
}

func Dec2Hex(i int64) string {
	//h := strconv.FormatInt(i, 16)
	h := fmt.Sprintf("%x", i)
	fmt.Printf("%d = 0x%s\n", i, h)
	return h
}
