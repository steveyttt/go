//Package acdc asks if you are ready to rock
package acdc

//Sum is used to add together a potentially unlimited number of integers
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
