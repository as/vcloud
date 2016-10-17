package main

import (
	"flag"
	"fmt"
	"os"
)

import (
	"github.com/as/vcloud/login"
)

type Args struct {
	Socket, Org, User, Pass, Env *string
}

func main() {
	a := parseargs()

	token, err := login.Do(*a.Socket, *a.Org, *a.User, *a.Pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(token)
	fmt.Fprint(os.Stderr, "\n")
}

// parseargs parses command line args passed in to the program
func parseargs() *Args {
	var a Args

	a.Socket = flag.String("s", "", "server: ex, example.com:443")
	a.Org = flag.String("o", "", "org: ex, strickland-west")
	a.User = flag.String("u", "", "user: ex, hankhill")
	a.Pass = flag.String("p", "", "propane")

	flag.Parse()

	// Prompt on stdin if args are unset

	if *a.Socket == "" {
		*a.Socket = ask("server: ")
	}

	if *a.Org == "" {
		*a.Org = ask("org: ")
	}

	if *a.User == "" {
		*a.User = ask("user: ")
	}

	if *a.Pass == "" {
		*a.Pass = ask("pass: ")
	}

	return &a
}

// ask asks the user the question 'q' and returns
// the answer 'a'
func ask(q string) (a string) {
	fmt.Fprintf(os.Stderr, q)
	fmt.Scanln(&a)
	return a
}