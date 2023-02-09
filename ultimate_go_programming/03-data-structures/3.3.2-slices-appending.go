package main

import "fmt"

func main() {

	var data []string //zero value, não aloca o backing array
	//len 0 cap 0 -> [nil][0][0]
	//datainit := []string{} // [*][0][0] inicializado com um ponteiro para um struct vazio
	//                         |
	//                         +-------- struct{}
	inspectSlice(data)
	lastCap := cap(data)

	//fazer append num slice com zero value (nil, 0, 0) cria o backing
	//array na heap de tamanho 1, o slice passado é mod p/ len 1 cap 1
	//e seu valor é copiado para a variável na esquerda da atribuição.
	//o slice criado na stack frame é limpo após o retorno de append.
	//cada vez que um novo array é criado, o antigo é coletado pelo GC

	//se o tamanho final já é conhecido, criar um slice com o tamanho evita
	//varias alocaçoes diferentes e a chamada de append
	//data := make([]string, 0, 1e5)

	//insere 100_000 elementos no slice
	for record := 1; record <= 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value) //append usa semantica de valor
		//o primeiro append será num slice zero valued.
		//o segundo cria outro backing array com o dobro da capacidade
		//anterior, copia o valor do anterior e adiciona o valor passado.
		//a cópia do slice é modificada para len2 e cap2 e seu valor é
		//retornado. Um terceiro append novamente dobra e fica len3 cap4.
		//no quarto não será necessaŕio novo array pois resta 1 de cap.
		//no quinto o processo se repete.

		//calcula porcentagem de aumento
		//quando o capacidade chega  a 10_000, append para de dobrar e passa a
		//aumentar em 25% a capacidade
		if lastCap != cap(data) {
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)

			//o endereço dos elementos muda quando o novo array é criado.
			//mostra o resultado - cada linha, uma nova alocação
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)
		}
	}

}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice { //semantica de valor
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
}
