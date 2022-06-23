// Arquivos que contem testes devem seguir o padrão XXX_test.go
package main

// o pacote padrão para realização de testes
import "testing"

// funções que executam testes devem seguir o padrão TestNomeDaFuncaoTestada
// e receber como parâmetro um ponteiro para o tipo testing.T, que possui as 
// funcionalidades de testes
func TestHello(t *testing.T) {
    // função auxiliar de teste criada quando o teste específico se repetirá
    // assertCorrectMessage está armazenando uma função anônima (quando não declarada com func)
    // o tipo testing.TB é um tipo que permite tanto testes comuns (T) quando banckmark (B)
    assertCorrectMessage := func(t testing.TB, got, want string) {
        // marcar o test como helper permite que, se o teste falhar, a saída tenha
        // informação da linha específica do teste que falhou. Caso contrário,
        // será a linha da definição do helper
        t.Helper()
        if got != want {
            // define um caso de falha para o teste e a mensagem que deve ser exibida
            // %q formata o valor com aspas duplas
            t.Errorf("got %q want %q", got, want)
        }
    }

    // t.Run("Nome do caso de teste", funçãoQueDefineOTesteExecutado)
    t.Run("saying Hello to people", func(t *testing.T) {
        got := Hello("Xibungo", "")
        want := "Olá, Xibungo"
        // uso do helper em cada caso de teste, assim fica mais legível e limpo
        assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Olá, mundo' when an empty string is supplied", func(t *testing.T){
        got := Hello("", "")
        want := "Olá, Mundo"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T) {
        got := Hello("Joana", "French")
        want := "Bonjour, Joana"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in English", func(t *testing.T){
        got := Hello("Frank", "English")
        want := "Hello, Frank"
        assertCorrectMessage(t, got, want)
    })
}
