package main

import (
	"fmt"
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
	// padrão de execução caso não seja preenchido pelo usuário
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
