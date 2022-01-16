package prep

import (
	"fmt"
	"log"
	"testing"
)

func TestWriteRandomNumToFile(t *testing.T) {
	max, count := 8192, 4096
	fname := fmt.Sprintf("rand%d.out", count)
	err := WriteRandomNumsToFile(fname, max, count)
	if err != nil {
		log.Fatal("WriteRandomNumsToFile err:", err)
	}
}

func TestGenRandomNum(t *testing.T) {
	nums := GenRandomNums(256, 256)
	fmt.Println(nums)
}

func TestReadNumFromFile(t *testing.T) {
	nums, err := ReadNumsFromFile("rand256.out")
	if err != nil {
		log.Fatal("ReadNumsFromFile err:", err)
	}
	fmt.Println(nums)
}
