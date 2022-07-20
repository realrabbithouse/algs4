package streye

import (
	"fmt"
	"testing"
)

func TestAlphabet(t *testing.T) {
	encode1 := BASE64.ToIndices("NowIsTheTimeForAllGoodMen")
	decode1 := BASE64.ToChars(encode1)
	fmt.Println(decode1)

	encode2 := DNA.ToIndices("AACGAACGGTTTACCCCG")
	decode2 := DNA.ToChars(encode2)
	fmt.Println(decode2)

	encode3 := DECIMAL.ToIndices("01234567890123456789")
	decode3 := DECIMAL.ToChars(encode3)
	fmt.Println(decode3)
}
