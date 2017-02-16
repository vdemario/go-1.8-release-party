package main

type A struct { a int }

type B struct { a int }

type C B

type D = B

func main() {
	var ( a A; b B; c C; d D )
	a = b	// não funciona
	b = c	// tb não
	d = b	// novo no Go 1.9!
}
