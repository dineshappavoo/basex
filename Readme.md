basex
=====
A native golang impalementation for basex encoding which produces youtube like video id.
Unfortunately there are only 10 digits to work with, so if you have a lot of records, IDs tend to get very lengthy. We could borrow characters from the alphabet as have them pose as additional numbers.

Or how to create IDs similar to YouTube e.g. yzNjIBEdyww

The alphabet has 26 characters. That's a lot more than 10 digits. If we also distinguish upper- and lowercase, and add digits to the bunch or the heck of it, we already have (26 x 2 + 10) 62 options we can use per position in the ID.

###Usage

```go
package main

import (
	"github.com/dineshappavoo/basex"
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
```



##Install

$go get "github.com/dineshappavoo/basex"


  
##Project Contributor(s)

* Dinesh Appavoo ([@DineshAppavoo](https://twitter.com/DineshAppavoo))