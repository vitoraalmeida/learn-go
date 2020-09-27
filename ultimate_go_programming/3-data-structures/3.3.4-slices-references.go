package main

import "fmt"

type user struct {
	likes int
}

func main6() {
	users := make([]user, 3) //cria um backing array de users

	shareUser := &users[1] //cria um ponteiro para um elem no backing array

	shareUser.likes++ //modifica

	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	fmt.Println("*********************")

	users = append(users, user{})
	// não há espaço, cria outro array, copia os valores, aponta o slice para
	// o novo array e retorna

	shareUser.likes++
	//não vai modificar o novo slive, pois está apontando para o array antigo.
	//mantem uma referencia para um array inutlizado. memory leak. Não vai ser
	//coletado a menos que o ponteiro seja deletado.
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	fmt.Println(shareUser)
	shareUser = nil
	fmt.Println(shareUser)

	friends := []string{"sdfsd", "mnmnm", "popop", "tytyty"}

	//mostra que essa forma de range itera realmente sobre uma cópia
	for _, f := range friends { //semantica de valor. itera sobre uma cópia
		friends = friends[:2] //altera o slice original
		fmt.Println(f)
		//continua mostrando os 4 originais, pois
		// está iterando sobre a copia
	}
	fmt.Println(friends) //mostra que foi realmente alterado

}
