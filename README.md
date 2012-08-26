char
====

####Converts characters from 

 * string to byte 
 * string to int 
 * byte   to string 
 * byte   to int 


Here is a little test program
```go

package main

import "gotamer/char"

func main() {

	var b byte
	var c string
	var i int

	a := char.NewString("h")
	b  = a.GetByte()
	c  = a.GetString()
	i  = a.GetInt()
	print(b, "\t", c, "\t", i, "\n")	
}

```