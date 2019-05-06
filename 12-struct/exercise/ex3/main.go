package main

import "fmt"

type vehicle struct {
	doors   int
	colours string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	v1 := vehicle{
		doors:   4,
		colours: "black",
	}

	t1 := truck{
		fourwheel: true,
		vehicle: vehicle{
			doors:   4,
			colours: "black",
		},
	}

	s1 := sedan{
		luxury: true,
		vehicle: vehicle{
			doors:   4,
			colours: "black",
		},
	}

	fmt.Println(v1)
	fmt.Println(v1.doors)
	fmt.Println(t1.doors)
	fmt.Println(s1.doors)

}
