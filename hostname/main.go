// hostname is a web server that announces the hostname of the server it's being served from

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Command-line flags.
var (
	httpAddr = flag.String("http", ":8080", "Listen address")
)

func main() {
	flag.Parse()
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}
	http.Handle("/", NewServer(hostName))
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

// Server implements the hostname server.
// It serves the user interface (it's an http.Handler).
type Server struct {
	hostname string
}

// NewServer returns an initialized hostname server.
func NewServer(hostname string) *Server {
	s := &Server{hostname: hostname}
	return s
}

// ServeHTTP implements the HTTP user interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := struct {
		HOSTNAME string
	}{
		s.hostname,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}

// tmpl is the HTML template that drives the user interface.
var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html>
    <head><title>Hej!</title></head>
    <body>
	<h2>What's my hostname?</h2>
	<h1>{{.HOSTNAME}}</h1>
</body></html>
`))
