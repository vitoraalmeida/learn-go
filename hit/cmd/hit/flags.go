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

// Mensagem adicional de uso
const usageText = `
Usage:
  hit [options] url
Options:`

type flags struct {
	url  string
	n, c int //n = número de requests | c = nível de concorrência
}

func (f *flags) parse(s *flag.FlagSet, args []string) (err error) {
	/* comentado pois passando diretamente os campos do struct que já temos ao invés de salvar na variável interna do flag

		var (
			// o pacote flag tem um objeto em que são armazenadas as Flags com seu nome
			// e um ponteiro do tipo especificado para a flag, que é retornado pelo
			// método flag.Tipo(). Quando chamamos flag.Parse(), os argumetos são passadps
			// , parseados pelos nomes definidos, o valor é armazenado e então podemos
			// ter acesso através do ponteiro.

			//       tipo    nome  default  uso
			u = flag.String("url", "", "HTTP server `URL` to make requests (required)")
			// marcar um valor com ` ` faz com que a mensagem de uso mostre que o valor
			// daquele argumento é do tipo entre as crases

			// defaults já definidos na função main
			n = flag.Int("n", f.n, "Number of requests to make")
			c = flag.Int("c", f.c, "Concurrency level")
			// podemos passar as flags em qualquer ordem, inclusive passar mais de
			// uma vez a mesma, sendo que o último valor será o definido
			// cada flag pode ser usada com - ou --
			// por padrão vem um -h ou -help implementado, que inclusive mostra os valores default
		)
	flag.Parse()
	f.url = *u
	f.n = *n
	f.c = *c
	*/

	s.Usage = func() {
		fmt.Fprintln(os.Stderr, usageText[1:])
		s.PrintDefaults()
	}
	// usando como positional argument -> tem que ser por último
	//flag.StringVar(&f.url, "url", "", "HTTP server `URL` to make requests (required)")

	// utlizando um tipo customizado que só permite inteiros positivos
	s.Var(toNumber(&f.n), "n", "Number of requests to make")
	s.Var(toNumber(&f.c), "c", "Concurrency level")
	if err := s.Parse(args); err != nil {
		return err
	}

	// o flag parseia as flags definidas e o que não estiver definido anteriormente
	// o flag ira adicionar numa lista. Se não houver, uma string vazia é retornada
	f.url = s.Arg(0)

	if err := f.validate(); err != nil {
		fmt.Fprintln(s.Output(), err)
		s.Usage()
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

/*
//parsing command line args manually
type parseFunc func(string) error

//                       nomeando o retorno já define a variável
func (f *flags) parse() (err error) {
	parsers := map[string]parseFunc{
		"url": f.urlVar(&f.url),
		"n":   f.intVar(&f.n),
		"c":   f.intVar(&f.c),
	}

	for _, arg := range os.Args[1:] {
		n, v, ok := strings.Cut(arg, "=")
		if !ok {
			continue
		}
		parse, ok := parsers[strings.TrimPrefix(n, "-")]
		if !ok {
			continue
		}
		if err = parse(v); err != nil {
			err = fmt.Errorf("invalid value %q for flag %s: %w", v, n, err)
			break
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
*/
