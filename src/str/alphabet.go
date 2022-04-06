package str

import (
	"fmt"
	"strings"
)

const (
	CharLength = 256
)

var (
	BINARY         = NewAlphabet("01")
	OCTAL          = NewAlphabet("01234567")
	DECIMAL        = NewAlphabet("0123456789")
	HEXADECIMAL    = NewAlphabet("0123456789ABCDEF")
	DNA            = NewAlphabet("ACGT")
	LOWERCASE      = NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	UPPERCASE      = NewAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	PROTEIN        = NewAlphabet("ACDEFGHIKLMNPQRSTVWY")
	BASE64         = NewAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	ASCII          = newAlphabet(128)
	EXTENDED_ASCII = newAlphabet(256)
)

type Alphabet struct {
	r        int    // 基数
	alphabet []byte // Alphabet can contain 256 different alphabets at most due to using byte.
	inverse  []int
}

func NewAlphabet(str string) Alphabet {
	var (
		unicode  = make([]bool, CharLength)
		alphabet = []byte(str)
		inverse  = make([]int, CharLength)
	)
	for _, b := range alphabet {
		if unicode[b] {
			panic(fmt.Sprintf("illegal alphabet: repeated character = %c", b))
		}
		unicode[b] = true
	}
	for i := range inverse {
		inverse[i] = -1
	}
	for i, b := range alphabet {
		inverse[b] = i
	}
	return Alphabet{
		r:        len(alphabet),
		alphabet: alphabet,
		inverse:  inverse,
	}
}

// newAlphabet maybe problematic when radix > 256.
func newAlphabet(radix int) Alphabet {
	alphabet := make([]byte, radix)
	inverse := make([]int, radix)
	for i := 0; i < radix; i++ {
		alphabet[i] = byte(i)
	}
	for i := 0; i < radix; i++ {
		inverse[i] = i
	}
	return Alphabet{
		r:        radix,
		alphabet: alphabet,
		inverse:  inverse,
	}
}

func (alp *Alphabet) Radix() int {
	return alp.r
}

func (alp *Alphabet) LgR() int {
	var logR int
	for t := alp.r; t >= 1; t = t / 2 {
		logR++
	}
	return logR
}

func (alp *Alphabet) Contains(b byte) bool {
	if int(b) >= len(alp.inverse) || alp.inverse[b] == -1 {
		return false
	}
	return true
}

func (alp *Alphabet) ToChar(index int) byte {
	if index < 0 || index >= alp.r {
		panic(fmt.Sprintf("index must between 0 and %d: %d", alp.r, index))
	}
	return alp.alphabet[index]
}

func (alp *Alphabet) ToIndex(b byte) int {
	if int(b) >= len(alp.inverse) || alp.inverse[b] == -1 {
		panic(fmt.Sprintf("character %c not in alphabet", b))
	}
	return alp.inverse[b]
}

func (alp *Alphabet) ToChars(indices []int) string {
	var builder strings.Builder
	for i := range indices {
		builder.WriteByte(alp.ToChar(indices[i]))
	}
	return builder.String()
}

func (alp *Alphabet) ToIndices(str string) []int {
	indices := make([]int, len(str))
	chars := []byte(str)
	for i, b := range chars {
		indices[i] = alp.ToIndex(b)
	}
	return indices
}
