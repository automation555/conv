package char

type Char struct {
	b byte
}

func NewByte(c byte) *Char {
	return &Char{c}
}

func NewString(c string) *Char {
	t := []byte(c)
	return &Char{t[0]}
}

func (c *Char) GetByte() byte {
	return c.b
}

func (c *Char) GetString() string {
	return string(c.b)
}

func (c *Char) GetInt() int {
	return int(c.b)
}
