package main

import (
	"fmt"
	"sort"
)

type user2 struct {
	name    string
	surname string
}

func main() {
	//maps não são uteis no seu zero value, tem que ser construido com make
	//ou na forma literal
	users := make(map[string]user2)
	//nem tudo pode ser uma key, mas tudo pode ser um value

	users["Roy"] = user2{"Rob", "Roy"}
	users["Michael"] = user2{"Michael", "Jackson"}
	users["Peter"] = user2{"Peter", "Parker"}
	users["Dare"] = user2{"Dare", "Devil"}

	for key, value := range users { // itera sobre copias da key e de value
		fmt.Println(key, value) //é aleatório, maps não são ordenados
	}

	fmt.Println()

	for key := range users {
		fmt.Println(key)
	}

	fmt.Println()

	delete(users, "Roy") //buint in function para deletar valores de maps

	roy := users["Roy"] //busca um valor, retorna zero value para user se não achar
	//não é tão útil pois o user pode ter sido colocado vazio e não terá como
	//se foi inicializado assim ou apenas não foi encontrado
	fmt.Println("Roy: ", roy)

	//buscar um valor num map dado uma key, tendo o valor e se foi encontrado.
	u, found := users["Roy"]
	fmt.Println("Roy > Found:", found, u)
	u, found = users["Peter"]
	fmt.Println("Peter > Found:", found, u)

	fmt.Println("===============")
	//map declaração literal
	users = map[string]user2{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}
	for key := range users {
		fmt.Println(key, users[key])
	}

	fmt.Println("===============")

	var keys []string
	//coleta as keys num slice, de forma aleatória
	for key := range users {
		keys = append(keys, key)
	}

	//ordena
	sort.Strings(keys)
	//olhando o map de forma ordenada atraves das keys ordenas num slice
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}
