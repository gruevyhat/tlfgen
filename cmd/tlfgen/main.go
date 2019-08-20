// Package main implements a simple CLI for tlfgen.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/gruevyhat/tlfgen"
)

var usage = `The Laundry Files Character Generator

Usage: tlfgen [options]
       tlfgen list

Options:
  -n, --name=<str>          The character's full name.
  -g, --gender=<str>        The character's gender.
  -a, --age=<int>           The character's age.
  -p, --personality=<str>   The character's personality type.
  -A, --assignment=<str>    The character's assignment.
  -P, --profession=<str>    The character's profession.
  -S, --skill-points=<str>  The character's bonus skill points. [default: 0]
  -b, --attr-bonus=<str>    One of {"smart", "tough", "mystical"}.
  -s, --seed=<hex>          Character generation signature.
  --log-level=<str>         One of {INFO, ERROR}. [default: ERROR]
  -h --help
  --version
`

func main() {
	opts := tlfgen.Opts{}
	optFlags, _ := docopt.ParseArgs(usage, nil, tlfgen.VERSION)
	optFlags.Bind(&opts)
	if opts.List {
		fmt.Println("Personality Types:",
			strings.Join(tlfgen.ListPersonalityTypeKeys(), ", "))
		fmt.Println("Professions:",
			strings.Join(tlfgen.ListProfessionKeys(), ", "))
		fmt.Println("Assignments",
			strings.Join(tlfgen.ListAssignmentKeys(), ", "))
	} else {
		c, err := tlfgen.NewCharacter(opts)
		if err != nil {
			fmt.Println("An error has occurred. Aborting.")
			os.Exit(1)
		}
		fmt.Println(c.ToJSON(true))
	}
}
