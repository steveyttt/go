package metallica

//Mymulti is used to multiply lots of different numbers
func Mymulti(xi ...int) int {
	s := 0
	for i, v := range xi {
		if i == 0 { //if index = 0 then add it to s
			s += v
		} else {
			s *= v
		}
	}
	return s
}
