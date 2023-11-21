package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/vitoraalmeida/learn-go/hit/hit"
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

func banner() string { return bannerText[1:] } //remove o primeiro \n
//func usage() string  { return usageText[1:] }  //remove o primeiro \n

func main() {
	if err := run(flag.CommandLine, os.Args[1:], os.Stdout); err != nil {
		os.Exit(1)
	}
}

func run(s *flag.FlagSet, args []string, out io.Writer) error {
	// padrão de execução caso não seja preenchido pelo usuário
	f := &flags{
		n: 100,
		c: runtime.NumCPU(),
	}

	if err := f.parse(s, args); err != nil {
		return err
	}

	fmt.Fprintln(out, banner())

	fmt.Fprintf(out, "Making %d requests to %s with a concurrency level of %d.\n",
		f.n, f.url, f.c)

	if f.rps > 0 {
		fmt.Fprintf(out, "(RPS: %d)\n", f.rps)
	}

	request, err := http.NewRequest(http.MethodGet, f.url, http.NoBody)

	if err != nil {
		return err
	}

	c := &hit.Client{
		C:   f.c,
		RPS: f.rps,
	}

	sum := c.Do(request, f.n)

	sum.Fprint(out)

	return nil
}
