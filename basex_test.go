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
		{"9007199254740989"},
		{"123456789012345678901234567890"},
		{"1234"},
		{"test/test/123"},
		{"https://tour.golang.org"},
		{"https://blog.golang.org"},
		{"http://golang.org/doc/#learning"},
	}
	for _, c := range cases {
		encode := Encode(c.in)
		decode := Decode(encode)
		if c.in != decode {
			t.Errorf("Encode(%q) == %q, Decode %q", c.in, encode, decode)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	s := "9007199254740989"
	for n := 0; n < b.N; n++ {
		_ = Encode(s)
	}
}

func BenchmarkDecode(b *testing.B) {
	s := "2aYls9bkamJJSwhr0"
	for n := 0; n < b.N; n++ {
		_ = Decode(s)
	}
}
