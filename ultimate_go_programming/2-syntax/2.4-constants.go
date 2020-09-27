package main

func main() {
	//constantes existem em tempo de compilação
	//tem um sistema de tipos paralelo
	//o compilador pode performar conversões implicitas em constantes sem tipo

	//tipo não está completamente definido, ainda podem ter conversões.
	//uma constante com um kind pode estar numa operação com uma tipada int64
	//aí ela vai ser convertida para int64
	const ui = 12354   // kind integer
	const uf = 3.14159 // kind floating-point

	const ti int = 12345       //type int64
	const tf float64 = 2.14159 //type float64

	//kind floating-point
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)
	//kind integer
	const zero = 1 / 3 //KindInteger(1) / KindInteger(3)

	const maxInt = 9223372036854775807

	//constantes tem precisão de 256 bits
	const biggerInt = 3429873094824987321947211239038210938

	//ERRO
	//var anothoerbiggerInt int = 3810293810298301298310298231231

}
