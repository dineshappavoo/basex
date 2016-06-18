basex   
=======
[![GoDoc](https://godoc.org/github.com/dineshappavoo/basex?status.svg)](https://godoc.org/github.com/dineshappavoo/basex) [![Build Status](https://travis-ci.org/dineshappavoo/basex.svg?branch=master)](https://travis-ci.org/dineshappavoo/basex)

A native golang implementation for basex encoding which produces youtube like video id.
There are only 10 digits to work with, so if you have a lot of records to maintain in the application, IDs tend to get very lengthy. `uuidgen` gives a very lengthy value. We can use characters from the alphabet as have them pose as additional numbers.

Or how to create IDs similar to YouTube e.g. yzNjIBEdyww

The alphabet has 26 characters. That's a lot more than 10 digits. If we also distinguish upper- and lowercase, and add digits to the bunch for the heck of it, we already have (26 x 2 + 10) 62 options we can use per position in the ID. Please note that this package only takes numeric inputs.

####Note: 11/14/2015 version 0.1.0 has a breaking change which has new 'error' return type.
          06/17/2016 version 0.1.1 has new functions EncodeInt and DecodeInt for processing big integers.

###Usage

####Encode and Decode 
```go
package main

import (
        "fmt"
        "math/big"
        "github.com/dineshappavoo/basex"
)

func main() {
        input := "123456789012345678901234567890"
        inputBigInt := big.NewInt(0)
        inputBigInt.SetString(input, 10)

        fmt.Println("Input : ", input)

        // encode and decode functions
        encoded, err := basex.Encode(input)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("Encoded : ", encoded)

        decoded, err := basex.Decode(encoded)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("Decoded : ", decoded)

        if input == decoded {
                fmt.Println("Passed! decoded value is the same as the original.")
        } else {
                fmt.Println("FAILED! decoded value is NOT the same as the original!!")
        }

        // encode int and decode int functions
        encodedInt, err := basex.EncodeInt(inputBigInt)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("Encoded using big int: ", encodedInt)

        decodedInt, err := basex.DecodeInt(encodedInt)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("Decoded using big int: ", decodedInt)

        if inputBigInt.Cmp(decodedInt) == 0 {
                fmt.Println("Passed! decoded int value is the same as the intput big int.")
        } else {
                fmt.Println("FAILED! decoded int value is NOT the same as the original!!")
        }
}
```

####output looks like,

```go
Input :  123456789012345678901234567890

Encoded :  2aYls9bkamJJSwhr0
Decoded :  123456789012345678901234567890
Passed! decoded value is the same as the original.

Encoded using big int:  2aYls9bkamJJSwhr0
Decoded using big int:  123456789012345678901234567890
Passed! decoded int value is the same as the intput big int.
```

##Install

```shell
go get github.com/dineshappavoo/basex
```
##Referrence
* [Kevin van Zonneveld's Blog](http://kvz.io/blog/2009/06/10/create-short-ids-with-php-like-youtube-or-tinyurl/)

  
##Project Contributor(s)

* Dinesh Appavoo ([@DineshAppavoo](https://twitter.com/DineshAppavoo))
* Daved ([@Daved](https://github.com/daved))
* Christian Mauduit ([@ufootorg](https://twitter.com/ufootorg)) 
* Jim Tittsler ([@jtittsler](https://twitter.com/jtittsler))
