package main

import "fmt"

const ebulicaoK = 373

func main() {

	var tempK = ebulicaoK
	tempC := tempK - 273
	fmt.Printf("Ponto de ebulição da água em Kelvin é %d e o ponto de ebulição da água em Celsius é %d", tempK, tempC)

}
