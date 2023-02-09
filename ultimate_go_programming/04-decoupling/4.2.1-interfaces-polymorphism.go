package main

//Polimorfismo significa que voce escreve um certo programa e ele se comporta
//de forma diferente dependendo do dado(concreto) em que ele opera.
//Tom Kertz (BASIC)

//Uma boa razão para um dado ter comportamento, é a necessidae de polimorfismo

import "fmt"

//interfaces definem um conjunto de métodos, um contrato de comportamento
type reader interface {
	read(b []byte) (int, error) //quem chama deve alocar e compartilhar
	/*
		read(n int) ([]byte, error)
		má escolha, pois tem que alocar []byte na heap, por ser retornado
		func read() ([]byte, error) {
			s := make([]byte, 1024)
			return s  //retorna o slice para main. Escape
		}
	*/

}

//é um reader pois implementa o método de mesma assinatura da interface.
//Go prefere convenção à configuração:
//implementou o método da interface, implementou a interface automaticamente
type file struct {
	name string
}

func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

//é um reader. implementou a interface
type pipe struct {
	name string
}

func (pipe) read(b []byte) (int, error) {
	s := `{name: "Bill", title: "Developer"}`
	copy(b, s)
	return len(s), nil
}

func main4() {
	var r reader   //NÃO cria um dado real, concreto, não é manipulável
	fmt.Println(r) //nil
	// r.??? não tem dado nenhum para preencher
	//retrieve(r)  -> error

	f := file{".bashrc"}
	p := pipe{"cfg_service"}

	retrieve(f) //retrieve recebe sua própria cópia de f e p
	retrieve(p)
}

//função polimórfica
//pede por qualquer tipo concreto que satisfaça o contrato de reader
func retrieve(r reader) error {
	/*
		quando um dado é passado para a função, r (valor da interface) passa a
		apontar para uma cópia do valor passado.
		r é um tipo de 2 words, 1 ponteiro para iTable, 1 ponteiro para a cópia
		do valor passado.
		iTable é uma matrix que relaciona os tipos existentes com suas funções

		      +---> iTable [ pipe||file ]            main:       f
		      |            [ * p/ função]                     [".bashrc"]
		r:	[ * ]            |
			[ * ]            read()                              p
		      |     copy         /\  chama a função           ["cfg_service"]
			  +---> p||f          |  passando a cópia
		            [  ] ---------+
	*/
	data := make([]byte, 1024)

	len, err := r.read(data) //consequências diferentes a depender do concreto
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}
