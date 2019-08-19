// Package main implements a simple web service for tlfgen.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"

	"github.com/docopt/docopt-go"
	"github.com/gorilla/mux"
	"github.com/gruevyhat/tlfgen"
)

var mutex sync.Mutex

var usage = `The Laundry Files Character Generation Service

Usage: tlfserv [options]

Options:
  --port PORT	  The listening port. [default: 8080]
  -h --help
  --version
`

var cmdOpts struct {
	Port string `docopt:"--port"`
}

func generate(w http.ResponseWriter, r *http.Request) {
	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	skillPoints, err := strconv.Atoi(r.URL.Query().Get("skill-points"))
	charOpts := tlfgen.Opts{
		Name:            r.URL.Query().Get("name"),
		Age:             age,
		Gender:          r.URL.Query().Get("gender"),
		PersonalityType: r.URL.Query().Get("personality-type"),
		Assignment:      r.URL.Query().Get("assignment"),
		Profession:      r.URL.Query().Get("profession"),
		SkillPoints:     skillPoints,
		AttributeBonus:  r.URL.Query().Get("attribute-bonus"),
		Seed:            r.URL.Query().Get("seed"),
		LogLevel:        "INFO",
	}
	mutex.Lock()
	c, err := tlfgen.NewCharacter(charOpts)
	if err != nil {
		fmt.Println("An error occurred:", err)
	}
	mutex.Unlock()
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(c)
}

func main() {
	optFlags, _ := docopt.ParseDoc(usage)
	optFlags.Bind(&cmdOpts)

	fmt.Printf("Service started at <http://localhost:%s>\n", cmdOpts.Port)

	runtime.GOMAXPROCS(runtime.NumCPU())
	router := mux.NewRouter()
	router.HandleFunc("/", generate).Methods("GET")
	router.HandleFunc("/generate", generate).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+cmdOpts.Port, router))
}
