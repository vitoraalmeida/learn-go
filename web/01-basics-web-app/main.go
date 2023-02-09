package main

import (
	"fmt"
	"net/http"
)

// A função que usaremos precisa ser compatível com o tipo HandlerFunc definido
// no pacote http do go
// A função que será usada para tratar a requisição precisa receber
// um argumento que implemente a interface http.ResponseWriter,
// que define um conjunto de métodos para formar uma resposta,
// como definição de status code, headers, cookies etc.
// Quem chama a função passando uma implementação concreta da interface é o
// http.HandleFunc internamente. É uma interface pelo fato de poder usar
// diferentes implementações, sendo uma delas a response do pacote de testes
// http.

// Recebe também um ponteiro para a representação
// da requisição que chegou no servidor o *http.Request
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// por padrão, o servidor do go vai definir como html, se o texto for html
	// para mudar
	//w.Header().Set("Content-Type", "text/plain")

	w.Header().Set("Content-Type", "text/html charset=UTF-8")

	/* Fprint espera um local onde será escrito o texto, o local de saída,
	   que precisa implementar a interface io.Writer
	   Nesse caso, estamos escrevendo o html no escritor da resposta http
	   fmt.Println chama fmt.Fprintln passando o os.Stdout como saída

	   fmt.Fprintln(os.Stdout, "texto")
	*/
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-7")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:pro.vitoralmeida@gmail.com\">pro.vitoralmeida@gmail.com</a>.</p>")
}

// handler para lidar de forma centralizada com os paths
func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

// Criação de um router que implementa a interface Handler que é esperada pelo
// http.ListenAndServer
// Útil para incluir outras informações que são compartilhadas por toda a aplicação
// como a conexão com o banco de dados
type Router struct{}

// O contrato ta interface Handler exige implementação da função abaixo
func (router Router) ServeHttp(w http.ResopndeWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	// registra funções que lidarão com requisições para caminhos específicos
	// qualquer caminho que não tenha sido especificado será respondido pela função
	// associada ao /
	//http.HandleFunc("/", handlerFunc) // inclusão de uma rota
	//http.HandleFunc("/contact", contactHandler) // inclusão de outra rota

	// inclusão de uma rota passando função que trata outros paths
	//http.HandleFunc("/", pathHandler)

	//fmt.Println("Starting ther server on :3000...")
	//http.ListenAndServe(":3000", nil) // nil significa que está usando o Handler DefaultServeMux, que é usado pelo http.HandleFunc
	// Começa a escutar requisições na porta 3000 e as que forem para o / serão
	// passadas para a handlerFunc

	//var router Router // inicialização do handler personalizado
	//http.ListenAndServe(":3000", router)

	var router http.HandlerFunc = pathHandler // routerConv foi declarado como um HandlerFunc
	// que é um tipo função que espera um ResponseWriter e um *Request que possui
	// um método associado ServeHTTP. O método chama o prório valor (que é uma função)
	// passando o writer e o request, de forma que a função que foi atribuída
	// a variável pode ser usada como um Handler automaticamente

	/*  código do pacote http do go
	type HandlerFund func(ResponseWriter, *Request)
	func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
		f(w,r)
	}
	*/
	//http.Handle("/contact", http.HandlerFunc(contactHandler)) // Handle é chamado internamente por HandlerFunc
	// passando a função que passamos para ele
	http.ListenAndServe(":3000", router)
	//http.ListenAndServe(":3000", http.HandlerFunc(pathHandler)) // convertendo diretamente
	// igual converter um var int32 em int64 int64(var)
}

/*
//exemplo sem comentários
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=UTF-8")
	fmt.Fprint(w, "<h1>Welcome to my fucking awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-7")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:pro.vitoralmeida@gmail.com\">pro.vitoralmeida@gmail.com</a>.</p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	//http.HandleFunc("/", homeHandler) // inclusão de uma rota
	//http.HandleFunc("/contact", contactHandler) // inclusão de outra rota
	//fmt.Println("Starting ther server on :3000...")
	//http.ListenAndServe(":3000", nil)

	//http.HandleFunc("/", pathHandler) // inclusão de uma rota passando função que trata outros paths
	//fmt.Println("Starting ther server on :3000...")
	//http.ListenAndServe(":3000", nil)

	//var router Router
	//fmt.Println("Starting ther server on :3000...")
	//http.ListenAndServe(":3000", router)

	//fmt.Println("Starting ther server on :3000...")
	//http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
*/
