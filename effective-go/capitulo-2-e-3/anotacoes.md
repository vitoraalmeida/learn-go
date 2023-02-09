# Testes

## Abordagem em Go

Em go quase tudo está na stdlib, assim como ferramental de testes.

Em go existem dois atores principais em testes:
1. A ferramenta de CLI - Vem junto com a linguagem em si. Acha os testes, compila e executa o binário
2. O pacote de testes - Ajuda a escrever os testes

A cli acha as funções e elas se comunicam com o pacote
A ferramenta chama o compilador para as funções de testes e os pacotes que eles dependem e empacote como binário
A ferramenta executa o binário e o pacote toma controle
O pacote executa as funções de teste no binário e mostra os resultados

## Unidade em Go

Uma unidade testável em Go, em geral, é o pacote. Ainda que no pacote tenham
diversos arquivos diferentes para cada tipo e métodos, a unidade testada é o pacote.
Existe um arquivo de teste para cada arquivo, onde se escreve os testes paras as funções,
mas a unidade testada é o pacote.

Um único teste de unidade testa se uma função está trabalhando corretamente.

Um teste de unidade apropriado tem as seguintes caracteristicas

1. Isolado - testa a lógica de uma pequena parte do código;
2. Rápido - Executa rapidamente e dá feedback imadiato sobre o código
3. Deterministico - consistente e dá sempre o mesmo resultado

O que é considerado pequeno depende do time

## Criando um pacote em go

```
mkdir project_name
cd project_name
go mod init github.com/your_username/project_name

importar:


import github.com/your_username/project_name/url

```

Um arquivo tem seu arquivo de teste relacionado com o sufixo *_test.go*.

## testando

código mínimo para um teste

```
package url

import "testing"

func TestFuncao(t *testing.T) {
}
```

A função de testes precisa sinalizar o pacote de testes se ele falha ou é bem sucedido
Um desses macanismos de comunicação é a instância `t` do tipo `*testing.T`. E com isso
podemos chamar as funções `t.Log()`, `t.Fail()` etc.

Para testar uma funcão, você realiza a chamada passando o input e compara
com o resultado esperado para aquele input. Se o resultado for diferente 
do esperado, podemos chamar `t.Fail()`

It's a good practice to write a test using the following phases:

    Arrange—You define some input values for the code that you want to test and your expectations of what you want the code to produce when you run it.
    Act—You run the code with the input values and save what you get from the code.
    Assert—You check whether the code works correctly by comparing the expected values with the actual values.


`t.Fail()` avisa a falha do teste, mas não interrompe a execução. `t.FailNow()`
além de chamar Fail, interrompe a execução da função de teste. Apesar de interromper
a execução daqueal função, não interrompe outras, pois cada função de teste é executada
em uma goroutine diferente.


