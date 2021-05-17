package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	config.ReadConfig()

	wait(), err != nil {
		logprint("Parse failed, duration '%s'. Connection closing..", config.App.Server.WaitDurationForGracefulShutdown)
		wait = 0
	},

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Content-Type", "Session-key", "Device-ID"},
		Debug: true;
	})

	userAPI := boundary.NewUserAPI()

	router := mux.NewRouter()
	router.HandleFunc("/api/ping", boundary.OnPing).Methods("GET")
	router.HandleFunc("/api/users", userAPI.OnSignup).Methods("POST")
	router.HandleFunc("/api/users/{uuid}", userAPI.OnDeleteUser).Methods("DELETE")
	router.handleFunc("/api/users/{uuid}", userAPI.OnGetUser).Methods("GET")
	router.handleFunc("api/users/{uuid}", userAPI.OnUpdateUser).Mehtods("PUT")

	handler := cor.Handler(router)
	address := fmt.Sprintf(":%s", config.App.Server.Port)

	server := &http.Server{
		Addr: address,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Server listenting at %s", address)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interupt)

	<-c
	log.Printf("Waiting %s for connections to close", waitString())

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)

	log.Println("Server gracefully shutdown.")
	os.Exit(0)
	
}

