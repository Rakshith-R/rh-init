# Golang

## Cheatsheet 

### 1. Structure (Package, import, exports, variables, functions): 
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

>Exported var or fucn needs to start with Capital name
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
```
### 1. Flow Control (for, if, else, switch &defer): 

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


## Driver 
