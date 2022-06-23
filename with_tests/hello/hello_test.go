// define que este arquivo pertence a um pacote
// um pacote agrega código que estão sob o mesmo contexto
package main

// importa um pacote
// fmt é um pacote que pertence a biblioteca padrão
import "fmt"

// constantes inserem o valor no "texto" do binário, então
// não memória auxiliar não é alocada para salvar o valor,
// por isso são mais econômicas em uso de memória
const spanish = "Spanish" // strings são definidas com " "
const french = "French"
const english = "English"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

// func declara uma nova função
// greetingPrefix é o nome da função
// (language string) é o parâmetro que a função espera
// para excecutar seu trabalho, sendo que "language" é o nome
// da variável que armazenará o valor passado quando ela for chamada
// e "string" declara o tipo do argumento
// (prefix string) é o valor de retorno da função
// quando o retorno é nomeado, a função terá acesso à variável
// que foi declarada como retorno, sendo que ela terá o valor padrão para o tipo
// definido
// func nome(argumento tipoDoArgumento) (retorno tipoDoRetorno)
func greetingPrefix(language string) (prefix string) {
    // switch é uma estrutura de controle que compara um valor
    // e se esse valor for igual ao das opções definidas, o código
    // da opção será executado
    switch language {
    case spanish:
        prefix = spanishHelloPrefix
    case french:
        prefix = frenchHelloPrefix
    case english:
        prefix = englishHelloPrefix
    // default é o caso que será executado caso nenhuma das opções bater com o 
    // valor comparado
    default:
        prefix = portugueseHelloPrefix
    }
    return // quando o valor de retorno é nomeado na assinatura da função,
    // é ele que será retornado caso retorn estiver sozinho
}

// funções que começam com letra maiúscula tem visibilidade entre pacotes
// em minúsculo possuem visibilidade apenas dentro do pacote
func Hello(name, language string) string {
    if name == "" {
        name = "Mundo"
    }
    // é possivel concatenar string usando a operação de soma
    return greetingPrefix(language) + name
}

func main() {
    // fmt.Println() imprime valores na saída padrão
    // nesse caso, foi passada uma função para ela, pois é importante isolar a regra
    // de negócio (o problema que está sendo resolvido) de fatores externos a ela
    // (imprimir o valor no terminal). Poderiamos retornar o valor da função Hello
    // para qualquer tipo de saída que aceite texto (arquivo, resposta http etc)
    // isso permite que o código seja mais testável, inclusive, já que ele em si
    // tem menos dependências
    fmt.Println(Hello("Batata", ""))
}
