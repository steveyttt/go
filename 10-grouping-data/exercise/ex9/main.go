package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "Martinis`", "Women"},
		"moneypenny_miss": []string{"james Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	}

	m["Q"] = []string{"Gadgets", "cars", "science"}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
