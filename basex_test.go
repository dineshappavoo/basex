package basex

import (
	"math/big"
	"strconv"
	"testing"
)

var successCases []struct{ in string }

func init() {
	successCases = []struct {
		in string
	}{
		{"999999999999"},
		{"9007199254740992"},
		{"9007199254740989"},
		{"123456789012345678901234567890"},
		{"1234"},
	}
}

func TestBasexSuccess(t *testing.T) {
	for _, c := range successCases {
		encode, err := Encode(c.in)
		if err != nil {
			t.Errorf("Encode error:%q", err)
		}
		decode, err := Decode(encode)
		if err != nil {
			t.Errorf("Decode error:%q", err)
		}
		if c.in != decode {
			t.Errorf("Encode(%q) == %q, Decode %q", c.in, encode, decode)
		}
	}
}

func TestBasexIntSuccess(t *testing.T) {
	for _, c := range successCases {
		b := big.NewInt(0)
		b.SetString(c.in, 10)
		encode, err := EncodeInt(b)
		if err != nil {
			t.Errorf("Encode error:%q", err)
		}
		decode, err := DecodeInt(encode)
		if err != nil {
			t.Errorf("Decode error:%q", err)
		}
		if b.Cmp(decode) != 0 {
			t.Errorf("Encode(%q) == %q, Decode %q (b == %q)", c.in, encode, decode.String(), b.String())
		}
	}
}

func TestBasexFailure(t *testing.T) {
	failureCases := []struct {
		in string
	}{
		{"test/test/123"},
		{"https://tour.golang.org"},
		{"https://blog.golang.org"},
		{"http://golang.org/doc/#learning"},
	}
	for _, c := range failureCases {
		encode, _ := Encode(c.in)
		decode, _ := Decode(encode)
		if c.in == decode {
			t.Errorf("Encode(%q) == %q, Decode %q", c.in, encode, decode)
		}
	}
}

func TestForLargeInputs(t *testing.T) {
	if testing.Short() {
		t.Logf("skipping large input test")
		return
	}
	for i := 1000; i < 3000000; i++ {
		encode, err := Encode(strconv.Itoa(i))
		if err != nil {
			t.Errorf("Encode error:%q", err)
		}
		decode, err := Decode(encode)
		if err != nil {
			t.Errorf("Decode error:%q", err)
		}
		if strconv.Itoa(i) != decode {
			t.Errorf("Encode(%q) == %q, Decode %q", i, encode, decode)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	s := "9007199254740989"
	for n := 0; n < b.N; n++ {
		_, _ = Encode(s)
	}
}

func BenchmarkDecode(b *testing.B) {
	s := "2aYls9bkamJJSwhr0"
	for n := 0; n < b.N; n++ {
		_, _ = Decode(s)
	}
}

func BenchmarkEncodeInt(b *testing.B) {
	var i big.Int
	i.SetString("9007199254740989", 10)
	for n := 0; n < b.N; n++ {
		_, _ = EncodeInt(&i)
	}
}

func BenchmarkDecodeInt(b *testing.B) {
	s := "2aYls9bkamJJSwhr0"
	for n := 0; n < b.N; n++ {
		_, _ = DecodeInt(s)
	}
}
