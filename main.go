package main

import "github.com/ananya/lukruk/managers"

func main() {
	numberOne := managers.Number{
		TitleNumber: "1", DescriptionNumber: "Number ONE",
	}
	println("Hello Golang!!! xxxx " + numberOne.DescriptionNumber)
}
