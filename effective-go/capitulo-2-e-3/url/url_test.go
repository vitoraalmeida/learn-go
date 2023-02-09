package url

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	const rawURL = "https://foo.com/go"

	u, err := Parse(rawURL)

	if err != nil {
		/*
			t.Logf("Parse(%q) err = %q, want nil", rawURL, err)
			//t.Fail() // não para a execução da função de teste, apenas avisa a falha
			t.FailNow()
		*/
		// Equivalente é chamar Logf e chamar FailNow
		t.Fatalf("Parse(%q) err = %q, want nil", rawURL, err)
	}

	want := "https"

	// got é inicializado apenas no escopo do if
	if got := u.Scheme; got != want {
		/*
			t.Logf("Parse(%q) -> Scheme = %q; want %q", rawURL, got, want)
			t.Fail()
		*/

		// Faz o mesmo que Logf seguido de Fail
		t.Errorf("Parse(%q).Scheme = %q; want %q", rawURL, got, want)
	}

	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawURL, got, want)
	}

	if got, want := u.Path, "go"; got != want {
		t.Errorf("Parse(%q).Path = %q; want %q", rawURL, got, want)
	}
}

// helper function não é chamada, pois o pacote de testing só chama funções
// top-level (que começam com letra maiúscula)
func testPort(t *testing.T, in, wantPort string) {
	// sem indicar que é uma helper funcion, o log do teste mostra que o erro
	// acontece na linha em que a função helper é definida, não onde é chamada
	// dificultando saber onde está falhando
	t.Helper()
	u := &URL{Host: in}
	if got := u.Port(); got != wantPort {
		t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
	}
}

func TestURLPortWithPort(t *testing.T) {
	const in = "foo.com:80"
	/*
		u := &URL{Host: in}
		if got, want := u.Port(), "80"; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	*/
	testPort(t, in, "80")
}

func TestURLPortWithEmptyPort(t *testing.T) {
	const in = "foo.com:"
	testPort(t, in, "")
}

func TestURLPortWithoutPort(t *testing.T) {
	const in = "foo.com"
	testPort(t, in, "")
}

func TestURLPortIPWithPort(t *testing.T) {
	const in = "1.2.3.4:90"
	testPort(t, in, "90")
}

func TestURLPortIPWithoutPort(t *testing.T) {
	const in = "1.2.3.4"
	testPort(t, in, "")
}

// table-driven testing
/*

   Reduces the amount of repetitive test code you need to write
   Reduces cognitive load by having related test cases in a single test function
   Allows you to add new test cases in the future quickly
   Makes it easier to see if you've covered the corner cases
   Avoids adding helper methods for shared logic

*/
func TestURLPort(t *testing.T) {
	tests := []struct {
		in   string // URL.Host field
		port string
	}{
		1: {in: "foo.com:80", port: "80"}, // with port
		2: {in: "foo.com:", port: ""},     // with empty port
		3: {in: "foo.com", port: ""},      // without port
		4: {in: "1.2.3.4:90", port: "90"}, // ip with port
		5: {in: "1.2.3.4", port: ""},      // ip without port
		// Add more tests in case of a need
	}
	for i, tt := range tests {
		u := &URL{Host: tt.in}
		if got, want := u.Port(), tt.port; got != want {
			t.Errorf("test %d: for host %q; got %q; want %q", i, tt.in, got, want)
		}
	}
}

//Table driven tests não permitem que possamos rodar apenas um teste sem que
//comentemos os outros, ou chamemos Fatalf (o que não faria que outras funções
// fossem paradas, apenas aquela top-level. Além disso, não podemos usar o
// a flag -shuffle para rodar os testes em ordem aleatória

// Para ter a aleatoriedade, poderiamos usar um map ou inves de slice, pois
// os elementos não tem ordem em maps

func TestURLPortMaps(t *testing.T) {
	tests := map[string]struct {
		in   string // URL.Host field
		port string
	}{
		"with port":       {in: "foo.com:80", port: "80"},
		"with empty port": {in: "foo.com", port: ""},
		"without port":    {in: "foo.com:", port: ""},
		"ip with port":    {in: "1.2.3.4:90", port: "90"},
		"ip without port": {in: "1.2.3.4", port: ""},
	}
	for name, tt := range tests {
		u := &URL{Host: tt.in}
		if got, want := u.Port(), tt.port; got != want {
			t.Errorf("test %q: for host %q; got %q; want %q", name, tt.in, got, want)
		}
	}
}

// Mas ainda não podemos rodar testes específicos. Para isso podemos usar subtests
func TestURLPortSubTests(t *testing.T) {
	// atribuição de uma função a uma variável por já estarmos numa função
	//high order function é uma função que retorna outra função
	testPort := func(in, wantPort string) func(*testing.T) {
		return func(t *testing.T) {
			t.Helper()
			u := &URL{Host: in}
			if got := u.Port(); got != wantPort {
				t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
			}
		}
	}
	t.Run("with port", testPort("foo.com:80", "80"))
	t.Run("with empty port", testPort("foo.com:", ""))
	t.Run("without port", testPort("foo.com", ""))
	t.Run("ip with port", testPort("1.2.3.4:90", "90"))
	t.Run("ip without port", testPort("1.2.3.4", ""))
}

// a forma idiomática é combinar as duas formas para ter os benefícios das duas
func TestURLPortSubTestsTable(t *testing.T) {
	tests := map[string]struct {
		in   string // url.host field
		port string
	}{
		"with port":       {in: "foo.com:80", port: "80"},
		"with empty port": {in: "foo.com", port: ""},
		"without port":    {in: "foo.com:", port: ""},
		"ip with port":    {in: "1.2.3.4:90", port: "90"},
		"ip without port": {in: "1.2.3.4", port: ""},
		// ...other tests
	}
	for name, tt := range tests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}

// para rodar um teste específico podemos passar a flag -run
/// go test -v -run=TestURLPortSubTests
// o -v mostra o resultado dos testes que passam também
// para rodar um subteste específico, podemos passar
/// go test -v -run=TestURLPortSubTests/with_port
// e todos os subtests que possuem with_port serão chamados
// para rodar apenas o with_port especificamente
/// go test -v -run=TestURLPortSubTests/^with_port
// para parar a execução dos testes no primeiro que falhar, podemos usar o
// -failfast
/// go test -v -run=TestURLPortSubTests -failfast

// podemos definir os casos de teste no level de pacote e passar para mais
// de uma função de teste
var hostTests = map[string]struct {
	in       string // URL.Host field
	hostname string
	port     string
}{
	"with port":       {in: "foo.com:80", hostname: "foo.com", port: "80"},
	"with empty port": {in: "foo.com", hostname: "foo.com", port: ""},
	"without port":    {in: "foo.com:", hostname: "foo.com", port: ""},
	"ip with port":    {in: "1.2.3.4:90", hostname: "1.2.3.4", port: "90"},
	"ip without port": {in: "1.2.3.4", hostname: "1.2.3.4", port: ""},
}

func TestURLHostnameFinal(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Hostname(), tt.hostname; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}

func TestURLPortFinal(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}
