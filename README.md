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
```int```s are numbers.
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
wrap them in ```"" or '' or \`\`\` ```
### **conditional**
### **forloop**
### **if ifelse**
### **loop**
### **switch**
### **fallthrough**
### **slice**
### **array**
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