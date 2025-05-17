package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	ebulicaoC := ebulicaoK - 273

	fmt.Printf("A temperatura de ebulição da água em Kelvin é %g, a temperatura de eblulição da água em Celsius é %g.", ebulicaoK, ebulicaoC)
}
