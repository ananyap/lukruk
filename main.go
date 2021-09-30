package main

import "github.com/ananyap/lukruk/managers"

func main() {
	numberOne := managers.Number{TitleNumber: "1", DescriptionNumber: "NumberONE"}
	println("Go is " + numberOne.DescriptionNumber)
}
