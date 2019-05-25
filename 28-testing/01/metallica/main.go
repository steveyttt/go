package metallica

//Mymulti is used to multiply lots of different numbers
func Mymulti(x ...int) int {
	s := 0
	for _, v := range x {
		s *= v
	}
	return s
}
