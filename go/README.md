# Golang

## Cheatsheet 

## 1. Structure (Package, import, exports, variables, functions): 
>first line is name of the package
```go
package go.com/cheatsheet
```
>import statement 

```go
 import "fmt" 
 ```

>Factored import statement

```go 
import ( 
    "fmt"
    "math/rand"
)
```

>Exported var or fucn needs to start with Capital Letter
```go 
var A int =1 //is accessible

var a int =1 // not accessible
```

>Variables deaclaration:  var < name > < type > = < value > or infer by < name > := < value >
```go 
var a int = 1
var a,b int =1,2
var (
    a int =1
    b float32 =2
    c string ="1"
)
a := 1 //shorthand available only inside functions
```
>basic Types
```go
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32, represents a Unicode code point
float32 float64
complex64 complex128

const Pi=3.14  // Cannot be declared using :=
```

>Type casting, needs explicit casting ! / type inference
```go 
var a int =42
var f float64= float62(i)
var i int
j := i // j is an int
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

> Functions
```go 
func function_name( [parameter list] ) [return_types]
{
   body of the function
}

//naked fucntions : named return parameters with empty return statement, return named params
func name(params) (a int){
    return 
}

```
## 2. Flow Control (for, if, else, switch &defer): 

>For , the only looping construct
```go
for i := 0; i < 10; i++ {
	sum += i
}
for ;sum < 10; {  //init and post statement optional
	sum += sum
}
for sum < 10{    // with no semicolons, acts as while
    sum+=1
} 
for{             //infinite loop
    sum++
}
```

>if, else
```go
if a>100{
    a++
}else{
    a--
}
if a:=i/10;a>100{    //shorthand a, the walrus operator  kinda stuff 
    i++
}else{
    i-- //var a available in else too
} 
```
> Switch
```go
switch a:=i/10;a{
    case 1:

    case 2:
    default:
}

```

> defer : A defer statement defers the execution of a function until the surrounding function returns.
> The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns(called at end of current block/ after the outer function returns and before poping it ). 
```go
defer fmt.Println("world")
fmt.Println("hello")
//output hello \n world \n
//multi defered func called in stack kinda lifo order
//Defer is commonly used to simplify functions that perform various clean-up actions.
```
## 3. More types : ptr, Structs, slices and maps
> Pointers (ditto to C , * - dereference, & - returns pointer to the var, except no pointer arithmetic)
```go
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
```

> Structs
```go
type Vertex struct{
    X,Y int
}

var (
    v1=Vertex{1,2}
    v2=Vertex{X:1}
    p=&Vertex{1,2}
)
```

> Arrays
```go
var a [10]int
var b [2]string
```


>Slices : A slice does not store any data, are like references to arrays  ;append ;  len() cap()  , empty ->nil
```go
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4]             //slice a[low:high]  high not included
fmt.Println(s)                        // [3 5 7]
s=append(s,2) //func append(s []T, vs ...T) []T
fmt.Println(s)                        // [3 5 7 2]

b=[]bool{true, true, false}            // creates an array, then builds a slice that references it:

fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) 

a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

```

> Range
```go
for index, value := range pow {
	fmt.Printf("2**%d = %d\n", index, value)
}
for i := range pow {}  // only index

```

> Maps
```go
m = make(map[string]Vertex)

```

