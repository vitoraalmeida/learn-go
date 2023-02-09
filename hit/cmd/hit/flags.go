package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type flags struct {
	url  string
	n, c int //n = número de requests | c = nível de concorrência
}

func (f *flags) parse() (err error) {
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
		return fmt.Errorf("invalid value %q for flag -url: %w", f.url, err)
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
