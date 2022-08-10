package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/mcustiel/go-blog/pkg/factory"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func main() {
	fmt.Println("Running...")
	var err error

	connectionManager := factory.CreateDbConnectionManager()
	connectionManager.Open()
	exitFunc := func() {
		log.Println("Closing db connection...")
		connectionManager.Close()
		log.Println("Closed...")
	}
	defer exitFunc()

	postHandler := factory.CreatePostHandler()

	rtr := factory.CreateRouter()
	err = postHandler.RegisterRoutes(rtr)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[DEBUG] - Request received: %s\n", r.URL.String())
		rtr.Route(w, r)
		PrintMemUsage()
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			exitFunc()
			log.Print("Aborted")
			os.Exit(1)
		}
	}()

	log.Println("Listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
