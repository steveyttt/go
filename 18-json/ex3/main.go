package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "MoneyPenny",
		Age:   27,
		Sayings: []string{
			"Bond is a dick",
			"welcome home james",
		},
	}

	u3 := user{
		First: "m",
		Last:  "Hmmmmm",
		Age:   54,
		Sayings: []string{
			"Bond youre too cocky",
			"I'm too old for this shit",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Something went tits up", err)
	}

}
