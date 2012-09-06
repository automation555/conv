// Converts characters from 
//
// string to byte 
// string to int 
// byte   to string 
// byte   to int

package char

import "strconv"

type Char struct {
	b byte
}

// Set a Byte character
func NewByte(c byte) *Char {
	return &Char{c}
}

// Set a String character
func NewString(c string) *Char {
	t := []byte(c)
	return &Char{t[0]}
}

// Returns a Byte representation
func (c *Char) GetByte() byte {
	return c.b
}

// Returns a String representation
func (c *Char) GetString() string {
	return string(c.b)
}

// Returns an Int representation
func (c *Char) GetInt() int {
	return int(c.b)
}

// Doesn't work, returns the UTF-8 representation of the string in int
func S2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Converts a uint to a string
func U2s(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Converts an int to a string
func I2s(i int) string {
	return strconv.Itoa(i)
}
