Arquivo hit.go

```
package main
 
import (
    "fmt"
    "log"
    "runtime"
)
 
const (
 
    bannerText = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/
`
 
    usageText = `
Usage:
  -url
       HTTP server URL to make requests (required)
  -n
       Number of requests to make
  -c
       Concurrency level`
)

func banner() string { return bannerText[1:] }
func usage() string  { return usageText[1:]  }
 
func main() {
    f := &flags{
        n: 100,
        c: runtime.NumCPU(),
    }
    var f flags
    if err := f.parse(); err != nil {
        fmt.Println(usage())
        log.Fatal(err)
    }
    fmt.Println(banner())
    fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n",
       f.n, f.url, f.c)
}
```

Arquivo flags.go

```Go
package main

import(
	"fmt"
	"os"
	"strconv"
	"strings"
)

 
type flags struct {
    url  string
    n, c int
}
type parseFunc func(string) error
 
func (f *flags) parse() (err error) {
    // a map of flag names and parsers.
    parsers := map[string]parseFunc{
        "url": f.urlVar(&f.url), // parses an url flag and updates f.url
        "n":   f.intVar(&f.n),   // parses an int flag and updates f.n
        "c":   f.intVar(&f.c),   // parses an int flag and updates f.c
    }
    for _, arg := range os.Args[1:] {
        n, v, ok := strings.Cut(arg, "=")
        if !ok {
            continue // can't parse the flag
        }
        parse, ok := parsers[strings.TrimPrefix(n, "-")]
        if !ok {
            continue // can't find a parser
        }
        if err = parse(v); err != nil {
            err = fmt.Errorf("invalid value %q for flag %s: %w", v, n, err)
            break    // parsing error
        }
    }
    return err
}
 
func (f *flags) urlVar(p *string) parseFunc {
    return func(s string) error {
        _, err := url.Parse(s)
        *p = s
        return err
    }
}
 
func (f *flags) intVar(p *int) parseFunc {
    return func(s string) (err error) {
        *p, err = strconv.Atoi(s)
        return err
    }
} 
```





O pacote *flag* ajuda a fazer o parsing de argumentos de linha de comando,
fazer a validação e mostrar mensagens de uso para cada flag.

Vem com a possibilidade de fazer parsing de strings, int, float, time duration etc.
Além disso é possível definir tipos próprios para que sejam parseados.

Primeiro executamos as funções de definição de flags, que irão armazenar num
*\*Flagset*, um conjunto de *\*Flag*s, retornando um ponteiro para a variável
interna do tipo da flag que foi definida.

Assim são os tipos como são definidos
```
+---------------+          +-----------------+  +-----------------+
| *FlagSet      |          |   *Flag         |  |  <<interface>>  |                                                    
| +-----------+ |          | Name     string |  |      value      |
| |    Flag   | |          | Usage    string |  |   Set(string)   |                                                
| |+---------+| |          | DefValue string |  |  String() string|
| ||  value *--------------->Value    Value *----->               |
| |+---------+| |          +-----------------+  +-----------------+
| +-----------+ |          
| +-----------+ |                  
| |    Flag   | |                
| |+---------+| |                  
| ||  value  || |              
| |+---------+| |          
| +-----------+ |          
+_--------------+         
```

ex.:

1 - `url := flag.String("url", ...)` -> regista a Flag no Flagset
2 - flagset [url *Flag | *string (retornado)] -> url
3 - `flag.Parse()` faz o parsing da flag
4 -> A flag parseada é inserida na variável interna cujo ponteiro foi retornado em 1/2

Para definir uma flag, deve-se chamar o pacote flag e a função de definição de
flags para cada tipo, passando o identificador da flag, um valor default e uma
mensagem de uso para a flag.

exemplo arquivo flags.go.go

```Go
package main
 
type flags struct {
    url  string
    n, c int
}

func (f *flags) parse() error {
   var (
       u = flag.String("url", "", "HTTP server URL to make requests (required)")
       n = flag.Int("n", f.n, "Number of requests to make")
       c = flag.Int("c", f.c, "Concurrency level")
   )
   flag.Parse()
   f.url = *u
   f.n = *n
   f.c = *c
   return nil
} 
```

Arquivo hit.go

```
package main
 
import (
    "fmt"
    "log"
    "runtime"
)
 
const (
 
    bannerText = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/
`
)

func banner() string { return bannerText[1:] }
 
func main() {
    f := &flags{
        n: 100,
        c: runtime.NumCPU(),
    }
    if err := f.parse(); err != nil {
        os.Exit(1)
    }
    fmt.Println(banner())
    fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n",
        f.n, f.url, f.c)
}
```

Dessa forma, o programa já fica com uma flag -h mostrando como usar

```
$ go run . –h
Usage of .../hit:
  -c int
        Concurrency level (default 10)
  -n int
        Number of requests to make (default 100)
  -url string
       HTTP server URL to make requests (required)
```


```
$ go run . -url=http://localhost:9090
Making 100 requests to http://localhost:9090 with a concurrency level of 10.
```

É possível evitar declaração de variáveis para receber o ponteiro do valor
que será lido das flags, passando o próprio ponteiro para os campos do
struct de definição do tipo `flags`. Mas o método do pacote flag é diferente

```Go
package main
 
type flags struct {
    url  string
    n, c int
}

func (f *flags) parse() error {
   flag.StringVar(&f.url, "url", "", "HTTP server `URL` to make requests (required)")
   flag.IntVar(&f.n, "n", f.n, "Number of requests to make")
   flag.IntVar(&f.c, "c", f.c, "Concurrency level")
   flag.Parse()
   return nil
} 
```
O pacote flag valida se o tipo passado é igual ao tipo esperado, bem como 
o tamanho do dado para numerais, mas não valida a presença ou ausência de
argumentos

As funções TipoVar criam um valor de um tipo interno do pacote que implementa
a interface Value, de forma que uma funçao "genérica" Var registra aquele valor
dentro do tipo flag, que guarda o value como sendo o tipo Value

```                                                                               
flag.StringVar(...)                                                              
      \/                                            +---------------+                            
  *stringValue                                      | *FlagSet      |                            
      \/                                            | +-----------+ |                            
flag.Var(value Value, name, usage string)---------------> *Flag   | |                          
                                                    | |+---------+| |                            
                                                    | ||  value  || |
                                                    | |+---------+| |                             
                                                    | +-----------+ |                           
flag.Var(value Value, name, usage string)           | +-----------+ |                          
                  /\                                | |    Flag   | |                           
              *intValue                             | |+---------+| |                                     
                  /\                                | ||  value  || |
         flag.StringVar(...)                        | |+---------+| |                                      
                                                    | +-----------+ |
                                                    +_--------------+

```

$ go run .
Making 100 requests to  with a concurrency level of 10.
-----------------------^ !
```

Podemos criar uma função de validação para isso.

```Go
package main
 
type flags struct {
    url  string
    n, c int
}

func (f *flags) parse() error {
   flag.StringVar(&f.url, "url", "", "HTTP server `URL` to make requests (required)")
   flag.IntVar(&f.n, "n", f.n, "Number of requests to make")
   flag.IntVar(&f.c, "c", f.c, "Concurrency level")
   flag.Parse()
       if err := f.validate(); err != nil {
       fmt.Fprintln(os.Stderr, err)
       flag.Usage()
       return err
   }
   return nil
} 

// validate post-conditions after parsing the flags.
func (f *flags) validate() error {
    if strings.TrimSpace(f.url) == "" {
        return errors.New("-url: required")
    }
    return nil
}
```
A função TrimSpace remove espaços em branco para garantir que não seja passado
um valor vazio `-url="    "`

Podemos usar o pacote url para fazer validações mais corretas, além de validar
se o número passado como o nível de concorrência tem que ser menor ou igual ao 
número de requests, pois não faria sentido termos 2 goroutines para uma request

```Go
package main
 
type flags struct {
    url  string
    n, c int
}

func (f *flags) parse() error {
   flag.StringVar(&f.url, "url", "", "HTTP server `URL` to make requests (required)")
   flag.IntVar(&f.n, "n", f.n, "Number of requests to make")
   flag.IntVar(&f.c, "c", f.c, "Concurrency level")
   flag.Parse()
       if err := f.validate(); err != nil {
       fmt.Fprintln(os.Stderr, err)
       flag.Usage()
       return err
   }
   return nil
} 

// o pacote flag consegue validar o tipo passado no argumento, mas precisamos adicionar outras validações
func (f *flags) validate() error {
	// não faz sentido usar mais goroutines que o numero de requests a serem feitos
	if f.c > f.n {
		return fmt.Errorf("-c=%d: should be less than or equal to -n=%d", f.c, f.n)
	}
	if err := validateURL(f.url); err != nil {
		//return fmt.Errorf("invalid value %q for flag -url: %w", f.url, err)
		return fmt.Errorf("url: %w", err)
	}
	return nil
}

func validateURL(s string) error {
	u, err := url.Parse(s)
	switch {
	// TrimSpace remove todos espaços ao redor
	case strings.TrimSpace(s) == "":
		err = errors.New("required")
	case err != nil:
		err = errors.New("parse error")
	case u.Scheme != "http":
		err = errors.New("only supported scheme is http")
	case u.Host == "":
		err = errors.New("missing host")
	}
	return err
}
```

Testando

```
$ go run . -url=http://somewhere -n=10 -c=20
-c=20: should be less than or equal to -n=10
<<usage>>
```

Podemos extender o pacote flag definindo tipos customizados que implementam
a interface Value.

Por exemplo, podemos criar um tipo numérico que apenas aceita valores
positivos para nossas flags C e N, pois nunca podem ser negativas. E com o
tipo customizado, implmentar a interface Value, de forma que a validação
possa ser feita diretamente no tipo.

Isso funciona pois o Método `Set(s string)` error da interface Value é responsável
por converter o valor que vem da linha de comando no tipo interno, e dentro
desse método podemos fazer as validações para tentar converter de fato o valor
textual. As funções `flag.TipoVar`, na verdade, convertem o valor passado num
Value e passam para a função `flag.Var` que é quem de fato cria a Flag e insere
no FlagSet, como mostrado anteriormente.

```                                                                               
flag.StringVar(...)                                                              
      \/                                            +---------------+                            
  *stringValue                                      | *FlagSet      |                            
      \/                                            | +-----------+ |                            
flag.Var(value Value, name, usage string)---------------> *Flag   | |                          
                                                    | |+---------+| |                            
                                                    | ||  value  || |
                                                    | |+---------+| |                             
                                                    | +-----------+ |                           
flag.Var(value Value, name, usage string)           | +-----------+ |                          
                  /\                                | |    Flag   | |                           
              *intValue                             | |+---------+| |                                     
                  /\                                | ||  value  || |
         flag.StringVar(...)                        | |+---------+| |                                      
                                                    | +-----------+ |
                                                    +_--------------+

```

O arquivo flags.go fica dessa forma

```
package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type flags struct {
	url       string
	n, c, rps int //n = número de requests | c = nível de concorrência | rps = throtling
}

func (f *flags) parse() error {
   flag.StringVar(&f.url, "url", "", "HTTP server `URL` to make requests (required)")
   flag.Var(toNumber(&f.n), "n", "Number of requests to make")
   flag.Var(toNumber(&f.c), "c", "Concurrency level")
   flag.Parse()
       if err := f.validate(); err != nil {
       fmt.Fprintln(os.Stderr, err)
       flag.Usage()
       return err
   }
   return nil
} 

// o pacote flag consegue validar o tipo passado no argumento, mas precisamos adicionar outras validações
func (f *flags) validate() error {
	// não faz sentido usar mais goroutines que o numero de requests a serem feitos
	if f.c > f.n {
		return fmt.Errorf("-c=%d: should be less than or equal to -n=%d", f.c, f.n)
	}
	if err := validateURL(f.url); err != nil {
		//return fmt.Errorf("invalid value %q for flag -url: %w", f.url, err)
		return fmt.Errorf("url: %w", err)
	}
	return nil
}

func validateURL(s string) error {
	u, err := url.Parse(s)
	switch {
	case strings.TrimSpace(s) == "":
		err = errors.New("required")
	case err != nil:
		err = errors.New("parse error")
	case u.Scheme != "http":
		err = errors.New("only supported scheme is http")
	case u.Host == "":
		err = errors.New("missing host")
	}
	// TrimSpace remove todos espaços ao redor
	return err
}

// number é um número natural
type number int

func toNumber(p *int) *number {
	return (*number)(p) // converte um *int para um *number, incluindo os métodos
	// de number a *int
}

func (n *number) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	switch {
	case err != nil:
		err = errors.New("parse error")
	case v <= 0:
		err = errors.New("should be positive")
	}
	*n = number(v)
	return err
}

func (n *number) String() string {
	return strconv.Itoa(int(*n))
}
```

Assim, a função toNumber é reponsável por converter o `*int`em `*number`, e a 
flag.Var vai receber esse `*number` e chamar a função Set do tipo number, que vai
fazer a converter a string passada na linha de comando em um `*number` e no processo
fazermos a verificação do valor
