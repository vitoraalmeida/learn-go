// pacote terminado com _test indica que estamos escrevendo um teste externo.
// testes externos apenas testam partes do código que estão publicos para os
// usuários
package url_test

import (
	"fmt"
	"log"

	"github.com/vitoraalmeida/url"
)

func ExampleURL() {
	u, err := url.Parse("http://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	fmt.Println(u)
	// Output: https://foo.com/go
}
