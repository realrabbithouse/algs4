// hello.go
package main

//void SayHello(const char* s); // declare SayHello
import "C"

import "fmt"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}
