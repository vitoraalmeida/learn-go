# Exemplos testáveis

Podemos definir exemplos para o uso dos nossos pacotes, mas na medida em que
mudamos o pacote, os exemplos vão ficando desatualizados.

Para garantir que isso não aconteça, podemos adicionar um exemplo que caso não
esteja conforme as alterações no pacote, cause um erro de compilação.

Os exemplos não podem ser executados, mas serão compilados igualmente e não
serão compilados se não estiver usando o pacote como deveria ser usado.

é possível adicionar comentários no fim do exemplo que indicam qual seria o 
output correto para a função

fmt.Println()
// Output:   //indica que o pacote de teste deve capturar o output da função de exemplo
// Output esperado

Se o output da função for diferent do esperado, o teste falha

Quando o output é de uma linha unica podemos fazer o seguinte:

// Output: output-esperado


Para outputs que não são ordenados, podemos usar 

// Unordered output:


Os exemplos testáveis são mostrados na documentação usando godoc e podemos executá-los

# Cobertura de testes

go test -coverprofile cover.out

Gera um arquivo que as linhas terminadas com 0 indicam quais linhas do código não
foram cobertas por testes

go tool cover -html=cover.out

Abre uma janlea de browser mostrandos os resultados

$ go tool cover -html=cover.out -o coverage.html

gera um arquivo html

go tool cover -func=cover.out

Mostra no terminal

Cobertura de testes não indica um código livre de bugs, pois ele apenas
mostra se foi testado os retornos das funções, mas não indica que cada caso
possível está testado

# Benchmarks

Funcionam como funções de test, Importamos um ponteiro para o tipo B, ao invés
de T.

Exemplo

```
func BenchmarkURLString(b *testing.B) {
    u := &URL{Scheme: "https", Host: "foo.com", Path: "go"}
    u.String()
}
```

Executamos a função que queremos testar e assim ela será testada a performace

Para executar o benchmark

`go test -bench .`

Os testes continuam sendo executados em paralelo, se quiser apenas executar
o benchmark

`go test -run=^$ -bench .`

Assim selecionar 0 testes para serem executados

A execução da função de benchmark retorna muito rapido sem que haja tempo de
o runtime ajustar o numero de vezes que será executado, quantas execuções etc

Então precisamos alterar a função para que o runtime tenha tmepo para se 
ajustar

``
func BenchmarkURLString(b *testing.B) {
    b.Logf("Loop %d times\n", b.N)
    u := &URL{Scheme: "https", Host: "foo.com", Path: "go"}
    for i := 0; i < b.N; i++ {
        u.String()
    }
}
``

b.N é o numero de vezes que o methodo vai ser executado, rodando a função u.String
o número de vezes que o runtime definiu que será executado garante que a função não
terminará de executar antes de o benchmark ser conlcuido

adiconar a chamada `b.ReportAllocs()` também mostra quantas vezes ocorreram alocações
por cada chamada da função e quanto de memória foi alocado por operação

