package _2_context_with_value

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextC, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	/*
		context.Background
		context.Background.WithValue(b, B)
		context.Background.WithValue(c, C)
		context.Background.WithValue(b, B).WithValue(d, D)
		context.Background.WithValue(c, C).WithValue(e, E)
		context.Background.WithValue(c, C).WithValue(f, F)
	*/

	// Ambil value dari context
	fmt.Println(contextF.Value("f")) // Dapat mengambil value dari diri sendiri
	fmt.Println(contextF.Value("c")) // Dapat mengambil value dari Parent C
	fmt.Println(contextF.Value("b")) // Tidak dapat mengambil value dari Context B (Beda Parent, tidak ada relasi)
	fmt.Println(contextA.Value("b")) // Tidak dapat mengambil value dari Child
	/*
		F
		C
		<nil>
		<nil>
	*/
}
