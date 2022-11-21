package main

import "fmt"

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	var v1, v2 = 10, 20
	swap(&v1, &v2)

	fmt.Println("v1=", v1, "v2=", v2)
}
