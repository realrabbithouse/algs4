package fileops

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func WriteRandomNumsToFile(path string, max, count int) error {
	rand.Seed(time.Now().UnixNano())
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	for i := 0; i < count; i++ {
		elem := rand.Intn(max)
		_, err := f.WriteString(strconv.Itoa(elem) + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func GenRandomNums(max, count int) (ret []int) {
	rand.Seed(time.Now().UnixNano())
	ret = make([]int, count)
	for i := 0; i < count; i++ {
		ret[i] = rand.Intn(max)
	}
	return
}

func ReadNumsFromFile(path string) (ret []int, err error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, fmt.Errorf("open file err: %q", err)
	}
	fstate, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat err: %q", err)
	}
	ret = make([]int, 0, fstate.Size()/8)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		elem, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("strconv err: %q", err)
		}
		ret = append(ret, elem)
	}
	return
}
