package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

// recebe uma url em string no formato "schema://host/path" e retorna
// um struct em que cada campo contem as partes da url separadas
func Parse(rawurl string) (*URL, error) {
	i := strings.Index(rawurl, "://")
	if i < 0 {
		return nil, errors.New("missing scheme")
	}

	scheme, rest := rawurl[:i], rawurl[i+3:]
	host, path := rest, ""
	// fazendo a atribuição no if, faz com que a variável
	// se restrinja ao escopo do if, não poluindo o namespace
	// da função
	if i := strings.Index(rest, "/"); i >= 0 {
		host, path = rest[:i], rest[i+1:]
	}

	return &URL{scheme, host, path}, nil
}

func (u *URL) Port() string {
	i := strings.Index(u.Host, ":")
	if i < 0 {
		return ""
	}
	return u.Host[i+1:]
}

// Hostname returns u.Host, stripping any port number if present.
func (u *URL) Hostname() string {
	i := strings.Index(u.Host, ":")
	if i < 0 {
		return u.Host
	}
	return u.Host[:i]
}
