package sayings

import "fmt"

//Greet is designed to send a nice greeting to someone
func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)

}
