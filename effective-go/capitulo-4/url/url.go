package url

import (
	"errors"
	"strings"
)

// A URL type represents a parsed URL from a String
type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse receives a string url and returns an instance of URL
func Parse(rawurl string) (*URL, error) {
	scheme, rest, ok := parseScheme(rawurl)
	if !ok {
		return nil, errors.New("missing scheme")
	}
	host, path := parseHostPath(rest)
	return &URL{scheme, host, path}, nil
}

func parseScheme(rawurl string) (scheme, rest string, ok bool) {
	// o começo da scheme signature deve estar pelo menos na posição da string
	// número 1, pois na 0 deve ter alguma scheme antes da assinatura para ser
	// valido
	return split(rawurl, "://", 1)
}

func parseHostPath(hostpath string) (host, path string) {
	host, path, ok := split(hostpath, "/", 0)
	if !ok {
		host = hostpath
	}
	return host, path
}

// Hostname returns u.Host, stripping any port number if present.
func (u *URL) Hostname() string {
	host, _, ok := split(u.Host, ":", 0)
	if !ok {
		host = u.Host
	}
	return host
}

// Port returns the port part of u.Host, without the leading colon.
// If u.Host doesn't contain a port, Port returns an empty string.
func (u *URL) Port() string {
	_, port, _ := split(u.Host, ":", 0)
	return port
}

// split s by sep.
//
// split returns empty strings if it couldn't find sep in s at index n.
func split(s, sep string, n int) (a, b string, ok bool) {
	i := strings.Index(s, sep)
	if i < n {
		return "", "", false
	}
	return s[:i], s[i+len(sep):], true
}

func (u *URL) String() string {
	/*
		u.String() é igual a (*url.URL).String(u), sendo que (*url.URL) é o
		tipo de u. Então se tivermos um URL nil, continua sendo um nil do tipo
		*URL, então o u é passado para String.
		Por isso podemos chamar o método String, mesmo para ponteiros nil
	*/
	if u == nil {
		return ""
	}
	// vai criando a URL na medida em que confirmamos que cada parte existe
	/*
		// concatenação de strings em go geram outras alocações de strings
		// se tivermos uma url que é muito chamada, teremos diversas alocações
		// de valores parecidos, fazendo com que muita memória seja gasta e que
		// o garbage collector deve ser chamado muitas vezes
		var s string
		if sc := u.Scheme; sc != "" {
			s += sc
			s += "://"
		}
		if h := u.Host; h != "" {
			s += h
		}

		if p := u.Path; p != "" {
			s += "/"
			s += p
		}
	*/

	// strings.Builder é um único buffer que permite a adição de strings e retorna uma lista delas
	var s strings.Builder
	if sc := u.Scheme; sc != "" {
		s.WriteString(sc)
		s.WriteString("://")
	}
	if h := u.Host; h != "" {
		s.WriteString(h)
	}
	if p := u.Path; p != "" {
		s.WriteByte('/')
		s.WriteString(p)
	}
	// o builder possui um método string para converter a lista numa só string
	return s.String()
}
