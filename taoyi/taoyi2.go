package main

func f() int {
	x := 1
	a := make([]int, 0, 20)
	b := make([]int, 0, 20000)
	l := 20
	c := make([]int, 0, l)

	a[0] = 1
	b[0] = 2
	c[0] = 3

	return x
}

func main() {
	f()
}
