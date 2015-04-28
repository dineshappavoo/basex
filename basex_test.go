package basex

import (
	"testing"
)

func TestBasex(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"999999999999"}, 
		{"9007199254740992"}, 
		{"9007199254740989"} ,
		{"123456789012345678901234567890"}, 
		{"1234"},
	}
	for _, c := range cases {
		encode := Encode(c.in)
		decode := Decode(encode)
		if c.in != decode {
			t.Errorf("Encode(%q) == %q, Decode %q", c.in, encode, decode)
		}
	}
}