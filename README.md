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
### **if ifelse else**
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
### **range**
### **make:**
### **anonymous struct**
### **embedded struct**
### **struct**
### **anonymous func**
### **package**
### **exported package**
### **callback**
### **defer**
### **func expression**
### **interfaces**
### **polymorphism**
### **methods**
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