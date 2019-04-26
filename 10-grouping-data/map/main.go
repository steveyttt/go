//maps are kv strings (JSON), un-ordered list

package main

import "fmt"

func main() {
	//Create a map with a KEY of string and a value type of INT
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 25, //even needs a trailing , (Not like JSON)
	}
	fmt.Println(m)
	fmt.Println(m["James"]) //acess by the key and get the value
	fmt.Println(m["Steve"]) //shows as 0 as value is not stored in map (returns ZERO value by default)
	//, ok idium -- Does value steve exist in the map m
	v, ok := m["Steve"]
	fmt.Println(v)
	fmt.Println(ok)

	//if steve is present in map set bool to true (ok) and run
	if v, ok := m["James"]; ok {
		fmt.Println("This is the if print", v)
	}
}
