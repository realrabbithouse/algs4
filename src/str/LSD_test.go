package str

import (
	"fmt"
	"testing"
)

func TestLSDSort(t *testing.T) {
	a := []string{
		"AVFA",
		"AVDX",
		"AVDF",
		"IOCS",
		"DFSC",
		"QWER",
		"DSJK",
		"DFJS",
		"DFSC",
	}
	LSDSort(a, 4)
	fmt.Println(a)
}
