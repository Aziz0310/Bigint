package main

import (
	"fmt"

	"github.com/Aziz0310/bootcamp/bigint/bigint"
)

func main() {
	a, err := bigint.NewInt("2222222222222")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("11")
	if err != nil {
		panic(err)
	}

	// err = b.Set("1") // b = "1"
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(b)

	c := bigint.Add(a, b)
	fmt.Println(c)        // c = "988847123412385995937737458960"
	d := bigint.Sub(a, b) // d = "988847123412385995937737458958"
	fmt.Println(d)
	e := bigint.Multiply(a, b) // e = "988847123412385995937737458959"
	fmt.Println(e)
	f := bigint.Mod(a, b) // f = "0"
	fmt.Println(f)
	// l := f.Abs()               // always returns pos num
	// fmt.Println(l)

	// fmt.Println(c, d, e, f, l)
}
