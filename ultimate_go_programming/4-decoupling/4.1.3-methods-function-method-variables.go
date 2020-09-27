package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) displayName() {
	fmt.Println("My name is", p.name)
}

func (p *person) setAge(age int) {
	p.age = age
	fmt.Println("New age is", p.age)
}

func main3() {
	p := person{name: "Bill"}

	//metodos são syntax sugar para mostrar que um dado tem certo comportamento
	//como fazemos:
	p.displayName()
	p.setAge(45)
	//o compilador faz chamada de função a partir do tipo passando a instância:
	//NÃO FAZER ISSO NA MÃO
	person.displayName(p)
	(*person).setAge(&p, 45)
	//tecnicamente, metodos são funções

	fmt.Println("==========================")

	//funções são dados, valores com tipo. Estrutura de 2 words
	f1 := p.displayName //f1 é um ponteiro para a função
	fmt.Print("f1(): ")
	f1()
	/*indireção dupla f1*->displayName*->        p
	  f1    displayName(value receiver)        +----+
	+---+    +---+                             |bill|
	| * ---> | * ----> Código de displayName   +----+
	+---+    +---+
	         | * ----> Cópia de p. A dupla indireção leva a alocar na heap
	         +---+

	alterações nos dados não serão refletidas por f1
	*/
	fmt.Println("Alterando name para 'Willian'")
	fmt.Print("bill.displayName(): ")
	p.name = "Willian"
	p.displayName()
	fmt.Print("f1(): ")
	f1()

	fmt.Println("==========================")

	f2 := p.setAge //f2 é um ponteiro para a função
	fmt.Print("f2(10): ")
	f2(10)
	/*indireção dupla f2*->setAge*->
	                                             p
	  f2     setAge(pointer receiver)          +----+  Ainda que p possa ficar
	+---+    +---+                           +>|bill|  na stack, a indireção
	| * ---> | * ----> Código de setAge      | +----+  dupla fará com que o
	+---+    +---+                           |         escape analsys aloque na
	         | * ----------------------------+         heap ainda assim. Falha.
	         +---+

		alterações nos dados serão refletidas
	*/
	fmt.Println("Age: ", p.age)
}
