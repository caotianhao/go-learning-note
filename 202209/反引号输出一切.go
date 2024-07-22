package main

import "fmt"

func main() {
	a := "golang go"
	fmt.Println("a =", a)
	b := "goo"
	fmt.Println("b =", b)
	c := `"package main
		import "fmt"
		func main() {
	//a := "golang go"
	/*fmt.Println(a)
	b := "goo"*/
	fmt.Println(b)"`
	fmt.Println("c =", c)
}
