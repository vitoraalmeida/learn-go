package caching

/*

Quanto mais longe do core do processador, mais demorado é o acesso
ao dado. Após o core, há 3 niveis de cache:

L1 - 64kb cache / core. 4 ciclos para acessar. 1ns. 16 instruçoes
L2 - 256kb cache / core. 15 ciclos. 4ns. 48 instruções
L3 - 8MB cache / core. 40 ciclos. 13ns. 160 instruções

Memória principal : 100 ciclos. 33.3 ns. 400 instruções

Se performace importa, é bom que se trabalhe com 8MB de dados/core

O processador busca dados na memória em "cache lines", pedaços de
memória que serão transportados até os caches. Geralmente 64 bytes

É mais performatico se a estrutura de dados for organizada de forma
que caiba em espaços proximos de memória, linearmente proximos, para
que caibam nas mesmas cache lines.

Construir os dados de forma que tenham padróes previsíveis de acesso
de memória.

Padrões previsiveis de acesso também são importantes para o cache do SO.
Dados contiguos também são preferíveis para serem armazenados nos pages do SO
Page tables = tabelas com endereços virtuais de mem criadas pelo SO
(varia de tamanho) 4k ~ 8 ~ 16k ... Exitem linux com 2mb de page.
Dados na mesma page table são mais faceis de trabalhar
*/

import "fmt"

const ( //														   (contiguos)
	rows = 5 * 1024 //atravessar fileiras é mais rapido, elementos proximos
	cols = 5 * 1024 //atravessar colunas custoso, elementos + afastados
	//				  ainda mais se as fileiras são longas como essas
)

//		   2048 x 2048 bytes
var matrix [rows][cols]byte

//nó para a linked list
type data struct {
	v byte
	p *data
}

//cabeça da lista
var list *data

func init() {
	//calda da lista
	var last *data

	//cria um nó para cada elemento do array
	//insere 0xFF em elementos pares tanto nos nós
	//quanto nos elementos do array. Serão contados
	//para ter mais trabalho no benchmark
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var d data
			if list == nil {
				list = &d
			}
			if last != nil {
				last.p = &d
			}
			last = &d

			if row%2 == 0 {
				matrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}

	var ctr int

	d := list
	for d != nil {
		ctr++
		d = d.p
	}

	fmt.Println("Elements in the link list", ctr)
	fmt.Println("Elements in the matrix", rows*cols)
}

//performace intermediária entre Colunas (-) e fileiras (+). Apesar de os
//ponteiros serem mais difíceis de estarem numa mesma cache line, são mais
//faceis de estarem numa mesma page table. Os primeiros elementos de uma linha
//muito grande são mais dificies de estarem numa mesma page table.
func LinkedListTraverse() int {
	var ctr int

	d := list
	for d != nil {
		if d.v == 0xFF {
			ctr++
		}

		d = d.p
	}
	return ctr
}

//como são fileiras longas, as colunas 0 de cada fileira estão mais afastadas
//na memória, mais custoso para acessar. Mais dificil de estarem num cache line,
//e mais dificil de estar na mesma page do SO
func ColumnTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

//cada coluna de uma fileira está ao lado da outra, elementos contiguos são
//preferidos pelo processador para serem levados ao cache. Mais rapido e mais
//consistente, criar padrões previsiveis e são "mais cacheados". Mais provaveis
//que estejam na mesma page do SO.
func RowTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}
