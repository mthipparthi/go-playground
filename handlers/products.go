package handlers

import (
	"log"
	"net/http"

	"github.com/mthipparthi/go-playground/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// p.l.Println("In Product")
	// lp := data.GetProducts()

	// err := lp.ToJSON(w)
	// if err != nil {
	// 	http.Error(w, "Failed to Convert to Json data", http.StatusInternalServerError)
	// }

	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("In Product")
	lp := data.GetProducts()

	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Failed to Convert to Json data", http.StatusInternalServerError)
	}
}
