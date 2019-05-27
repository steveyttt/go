# A Repo to host all golang basics

The Repo is built based upon lessons provided by Todd Mcleod - https://www.udemy.com/course/learn-how-to-code/. It can be used as a cheat sheet for most common coding scenarios which can be done with the golang standard lib.

### **var**
```var``` can be used to create local & package wide variables. You can set local scope with ```:=```, this is the short hand declaration method.
```
    var z int = 10
```

### **type**
GO is all about ```TYPE``` Type can be set when declaring vars, you can configure functions to only allow parameters which are of a specific type. You can create new custom types, you can create custom types based upon pre-existig types too.
```
    type hotdog int

    var a int
    var b hotdog
    a = int(b) //convert b from hotdog to a int
```

### **bitshifting**
Shift the bits of a binary number left

### **bool**
Bool is bool true or false

### **constant**
Like vars but once defined cannot change
```
    const a = 42
```

### **int**
Ints are numbers.
```
    //int can be positive and negative
    //uint can only be positive
    //byte = uint8 (0-255)
    //rune = int32 (-2147483648 to 2147483647)
    // the number at the after int and uint dictates the number of bytes to assign the variable
    var x int
    var y float64
    var z int8 = 127   // int 8 can go up to 127, if you set this to 128 it will fail
    var zz int16 = 256 // int 16 can go up to 32767
```

### **iota**
```iota``` is a simple way to incrememt a number.

### **modulus**
defined as ```%``` used to divide numbers and provide the remainder
```
    y := 83 % 40 //83 goes into 40 twice and leaves a remainder of 3
```

### **runtime**
Simple package to provide system specific stats at runtime
```
    fmt.Println(runtime.GOOS)
```

### **string**
wrap them in ```"" or '' ```. You can do multi line ones too
```
    s := `"Hello,
        SteveyT"`
```

### **conditional**
you can use ```bool``` to set conditionals in code
```
    fmt.Println(true && true)  //true AND true make true
    fmt.Println(true && false) //true AND false together make false - both need to be true
    fmt.Println(true || true)  //|| is OR, either can be true
    fmt.Println(true || false) //|| is OR, either can be true
    fmt.Println(!true)         //not true (false)
    fmt.Println(!false)        //not false (true)
```

### **forloop**
loop over values and re-run code... INIT CONDITION POST
```
	//https://golang.org/ref/spec#IncDec_statements
	//for statement with for clause
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
```

### **loop**
Loops dont need to have conditions.
```
	//while x is less than true run
	//for statement with single condition
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for {
		//if y is greater than 9 break from the for loop
		if y > 9 {
			break
		}

		fmt.Println(y)
		y++
	}
```

You can break and continue loops
```
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue //If x divided by 2 has a remainder then continue to next loop
		}
		fmt.Println(x)
	}
```
### **if else if else**
if this run this ifelse this then run this, else run this.
```
	x := 42
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was", x)
	}
```

### **switch**
switch evaluates a conditional statement and executes on first match only if you add the FALLTHROUGH keyword it will allow all other matching case statements below to also execute avoid FALLTHROUGH though - it drops wonky logic everywhere.
```
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print2")
	case (3 == 3):
		fmt.Println("This should print cos 3 == 3")
		fallthrough
	case (4 == 4):
		fmt.Println("This should now print")
	default: // executed if no statements match
		fmt.Println("This is default") //if nothing matches run the DEFAULT
	}
```

### **array**
Make an array - create an array of maximum 50 ints and creates a slice structure with a length of 10 (Length and capacity) - this allows the slice to grow by 40 more values until the arra needs to expand
```
	x := make([]int, 10, 50)

```

### **slice**
Create a slice. slices are pointers to arrays and scale dynamically based on added values
```
    x := []int{4, 5, 7, 8, 42}
```

### **map**
Used to store values, needs defined type
```
	//Create a map with a KEY of string and a value type of INT
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 25,
	}
```

### **range**
You can range over a map as needed. You can range over lots of things.
```
	//print out the key value pairs of all entries in the map
	for k, v := range m {
		fmt.Println(k, v)
	}
```

### **make**
make is a built in key word used to create a map, slice or channel
```
    x := make([]int, 10, 50)
    ch := make(chan int)
```

### **struct**
Think of structs as a data structure (OBJECT). Can contain values of differnt types i.e. int, string. Type is defined in struct.

```
    type person struct {
        first string
        last  string
        age   int
    }
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   38,
	}
    fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
```

### **anonymous struct**
Structures of data. anonymous ones dont have a name.
```
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
```

### **embeddedstruct**
A struct nested into another struct
```
    type person struct {
        first string
        last  string
        age   int
    }

    //referencing the type person inside a new type called secretAgent
    type secretAgent struct {
        person
        first string
        ltk   bool
    }

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "collision", // works but ugly. Best not to have double key names.
		ltk:   true,
	}
```

### **package**
All code can be broken into packages, you can import packages using the ```import``` keyword

### **exported package**
Exported packages need capitol first letters to be accessible by other packages. For example function ```Sum()``` is exported but ```sum()``` is not.

### **callback**
passing a func as an argument

### **defer**
```defer``` is used to "defer" certain calls until the function exits. IE if you open a file in a function you can defer closing it until the function closes.
```
    func main() {
        defer foo() //foo will run after bar as it has been deferd to run at function exit
        bar()
}
``` 

### **function**
a set of code you can invoke as needed.  
```
    //func (r receiver) identifier (parameters) (return (s)) {function code}

    //basic function
    func foo() {
        fmt.Println("hello there matey moo")
    }

    //func mouse takes 2 string arguments and returns a string and a bool
    func mouse(fn string, ln string) (string, bool) {
        a := fmt.Sprint(fn, " ", ln, `, says "hello"`)
        b := false
        return a, b
    }

```

### **func expression**
Store a function as an expression.
```
	f := func() {
		fmt.Println("my first func expression")
	}
```

### **interfaces**
Interfaces say that if you can use these methods you are of my type. This works as values can be of MULTIPLE types.
```
    //an interface says hey baby if you can run method speak, youre my TYPE :)
    type human interface {
        speak()
    }
```

### **polymorphism**
Polymorphism is when a function which can take multiple types as parameters. A secret agent and a person perhaps?

### **methods**
A method is a function which belongs to a certain type.
```
    // method speak can only be called from type secretAgent
    func (s secretAgent) speak() {
        fmt.Println("I am", s.first, s.last)
    }
```

### **returnedfunc**


### **unfurling**
### **variadic function / parameter**
### **pointer address**
### **pointer value**
### **methodsets**
### **bcrypt**
### **marshall**
### **unmarshall**
### **atomic**
### **channel**
### **mutex**
### **racecondition**
### **waitgroup**
### **channels**
### **comma ok idiom**
### **directional channel**
### **fan-in**
### **fan-out**
### **select channel**
### **erorr type**
### **panic**
### **fatal**
### **recover**
### **composite literal**
### **testing**
### **coverage tests**