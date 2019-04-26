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
	//comma OK IDIUM
	v, ok := m["Steve"]
	fmt.Println(v)
	fmt.Println(ok)

	//if steve is present in map set bool to true (ok) and run
	if v, ok := m["James"]; ok {
		fmt.Println("This is the if print", v)
	}

	//add a new kv to the map
	m["todd"] = 33
	fmt.Println(m)

	//print out the key value pairs of all entries in the map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//delete from a map
	delete(m, "todd")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("Removing", v)
		delete(m, "Miss Moneypenny")
	}

	//delete from a map - even if the user oesnt exist in map the delete does not fail
	delete(m, "bob")
	fmt.Println(m)

	//range through a
	xi := []int{1, 2, 3, 4, 5}
	for i, v := range xi {
		fmt.Println(i, v)
	}

}
