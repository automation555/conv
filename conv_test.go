package conv

import "testing"

func TestChar(t *testing.T) {

	var b byte
	var c string
	var in int

	text := "Printing byte, int and string value of this text"
	for i := 0; i < len(text); i++ {
		cb := NewByte(text[i])
		b = cb.GetByte()
		c = cb.GetString()
		in = cb.GetInt()
		print(b, "\t", c, "\t", in, "\n")
	}
}

func TestC2b(t *testing.T) {
	char := "T"
	println(char)
	b := C2b(char)
	println("S2b: ", b)
	s := B2c(b)
	println("B2s: ", s)
}

func TestS2bl(t *testing.T) {
	text := "Test String: test string"
	println(text)
	b := S2bl(text)
	println("S2bl: ", b)
	s := Bl2s(b)
	println("Bl2s: ", s)
}
