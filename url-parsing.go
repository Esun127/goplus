package main

import "fmt"

import "net/url"


import "strings"


func main() {
	p := fmt.Println
	s := "postgres://user:passwd@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u)

	p(u.User)
	p(u.User.Username())
	pass, _ := u.User.Password()
	p(pass)

	p(u.Host)
	h := strings.Split(u.Host, ":")
	p(h[0])
	p(h[1])

	p(u.Path)
	p(u.Fragment)

	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])

}
