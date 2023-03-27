package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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

/* utilizando usage do flag
	usageText = `
Usage:
  -url
       HTTP server URL to make requests (required)
  -n
       Number of requests to make
  -c
       Concurrency level`
*/
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

	return nil
}
