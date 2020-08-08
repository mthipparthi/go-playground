package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("In Good bye")
	fmt.Fprintf(w, "Good Bye")
	w.Write([]byte("\n Another one too...."))
}
