package main

import (
	"fmt"

	"github.com/steveyttt/go/28-testing/01/acdc"
	"github.com/steveyttt/go/28-testing/01/metallica"
	"github.com/steveyttt/go/28-testing/01/soad"
)

func main() {

	fmt.Println("the numbers add up to", acdc.Sum(1, 2, 3, 4))
	fmt.Println("the numbers multiply to be", metallica.Mymulti(2, 2, 2))
	fmt.Println("The numbers divided equal", soad.MyDivision(4, 2))

}
