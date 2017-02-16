package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8080", Handler: http.DefaultServeMux}

	go func() {
		fmt.Println("Aperte enter para parar o servidor")
		fmt.Scanln()
		log.Println("Parando o servidor...")
		if err := srv.Shutdown(context.Background()); err != nil { // HL
			log.Fatalf("não foi possível parar: %v", err)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Feliz dia do Go 1.8")
	})
	if err := srv.ListenAndServe(); err != http.ErrServerClosed { // HL
		log.Fatalf("ouvindo: %s\n", err)
	}
}
