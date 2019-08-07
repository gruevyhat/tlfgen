package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"

	"github.com/docopt/docopt-go"
	"github.com/gorilla/mux"
	"github.com/gruevyhat/sotdlgen"
)

var mutex sync.Mutex

var usage = `M6IK Character Generation Service

Usage: m6ikserv [options]

Options:
  --port PORT	  The listening port. [default: 8080]
  -h --help
  --version
`

var cmdOpts struct {
	Port string `docopt:"--port"`
}

func generate(w http.ResponseWriter, r *http.Request) {
	charOpts := sotdlgen.Opts{
		Name:       r.URL.Query().Get("name"),
		Gender:     r.URL.Query().Get("gender"),
		Level:      r.URL.Query().Get("level"),
		Ancestry:   r.URL.Query().Get("ancestry"),
		ExpertPath: r.URL.Query().Get("expert-path"),
		MasterPath: r.URL.Query().Get("master-path"),
		NovicePath: r.URL.Query().Get("novice-path"),
		Seed:       r.URL.Query().Get("seed"),
		LogLevel:   "ERROR",
	}
	mutex.Lock()
	c, err := sotdlgen.NewCharacter(charOpts)
	if err != nil {
		fmt.Println("An error occurred:", err)
	}
	mutex.Unlock()
	json.NewEncoder(w).Encode(c)
}

func main() {
	optFlags, _ := docopt.ParseDoc(usage)
	optFlags.Bind(&cmdOpts)

	fmt.Printf("SotDL Character Generation Service started at <http://localhost:%s>\n", cmdOpts.Port)

	runtime.GOMAXPROCS(runtime.NumCPU())
	router := mux.NewRouter()
	router.HandleFunc("/", generate).Methods("GET")
	router.HandleFunc("/generate", generate).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+cmdOpts.Port, router))
}
