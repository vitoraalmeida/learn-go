//o pacote pode ser esse arquivo sozinho ou mais qualquer arquivo que
//declare package main
package main

import "fmt"

//todo o pacote consegue acessar
//o tipo foi inferido para o padrão de floats (o float64)
const boilingF = 212.0

func amain() {
	//variaveis locais à função
	//tipo inferido para o tipo de boilingF
	var f = boilingF
	var c = fToC(f)
	fmt.Printf("boiling point = %gºF or %gºC\n", f, c)
	// Output:
	// boiling point = 212ºF or 100ºC
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
