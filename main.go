package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mthipparthi/go-playground/handlers"
)

func main() {

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
	// 	// log.Println("Hello world")
	// 	// data, err := ioutil.ReadAll(r.Body)
	// 	// if err != nil {
	// 	// 	http.Error(rw, "Oops", http.StatusBadRequest)
	// 	// 	// fmt.Println("Failed to read", data)
	// 	// 	// rw.WriteHeader(http.StatusBadRequest)
	// 	// 	// rw.Write([]byte("oops"))
	// 	// 	return
	// 	// }

	// 	// fmt.Fprintf(rw, "Hello %s", data)

	// })

	// http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request){
	// 	log.Println("Good Bye World")
	// })

	// log.Println("Starting Server")
	// err := http.ListenAndServe(":9090", sm)
	// if err != nil {
	// 	log.Println("Failed to Start Server")
	// }

	l := log.New(os.Stdout, "my-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	hg := handlers.NewGoodbye(l)
	hp := handlers.NewProduct(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodbye", hg)
	sm.Handle("/products", hp)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Got signal:", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
