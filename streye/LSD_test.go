package streye

import (
	"fmt"
	"testing"
)

func TestLSD(t *testing.T) {
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
	LSD(a, 4)
	fmt.Println(a)
}
