// Package main implements a simple CLI for tlfgen.
package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/gruevyhat/tlfgen"
)

var usage = `The Laundry Files Character Generator

Usage: tlf [options]

Options:
  -n, --name=<str>          The character's full name.
  -g, --gender=<str>        The character's gender.
  -a, --age=<int>           The character's age.
  -p, --personality=<str>   The character's personality type.
  -A, --assignment=<str>    The character's assignment.
  -P, --profession=<str>    The character's profession.
	-S, --skill-points=<str>  The character's bonus skill points. [default: 0]
	-b, --attr-bonus=<str>    The character's attribute bonuses ("smart", "tough", or "mystical".
  -s, --seed=<hex>          Character generation signature.
  --log-level=<str>         One of {INFO, WARNING, ERROR}. [default: ERROR]
  -h --help
  --version
`

func main() {
	opts := tlfgen.Opts{}
	optFlags, _ := docopt.ParseArgs(usage, nil, tlfgen.VERSION)
	optFlags.Bind(&opts)
	c, err := tlfgen.NewCharacter(opts)
	if err != nil {
		fmt.Println("An error has occurred. Aborting.")
		os.Exit(1)
	}
	fmt.Println(c.ToJSON(true))
}
