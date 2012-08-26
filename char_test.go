package char

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
