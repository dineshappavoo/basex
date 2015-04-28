package basex

import (
	"basex"
	"fmt"
)

func main() {
	input := "123456789012345678901234567890"
	fmt.Println("Input : ", input)

	encoded := basex.Encode(input)
	fmt.Println("Encoded : ", encoded)

	decoded := basex.Decode(encoded)
	fmt.Println("Decoded : ", decoded)

	if input == decoded {
		fmt.Println("Passed! decoded value is the same as the original. All set to Gooooooooo!!!")
	} else {
		fmt.Println("FAILED! decoded value is NOT the same as the original!!")
	}
}
