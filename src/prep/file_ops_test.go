package prep

import (
	"fmt"
	"log"
	"testing"
)

func TestWriteRandomNumToFile(t *testing.T) {
	max, count := 8192, 4096
	fname := fmt.Sprintf("rand%d.out", count)
	err := WriteRandomNumToFile(fname, max, count)
	if err != nil {
		log.Fatal("WriteRandomNumToFile err:", err)
	}
}

func TestGenRandomNum(t *testing.T) {
	nums := GenRandomNum(256, 256)
	fmt.Println(nums)
}

func TestReadNumFromFile(t *testing.T) {
	nums, err := ReadNumFromFile("rand256.out")
	if err != nil {
		log.Fatal("ReadNumFromFile err:", err)
	}
	fmt.Println(nums)
}
