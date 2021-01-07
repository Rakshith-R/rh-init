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

>Exported var or func needs to start with Capital Letter
```go 
var A int =1 //is accessible

var a int =1 // not accessible
```

>single and multi line commments
```go
//single line comment
/*
multi line
comment
*/
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
>basic Types + const + typedef ->type
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
type Myint int
type Name struct{
    first_name,last_name string
}
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

//naked fucntions : named return parameters with empty return(naked) statement, return named params
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
    p=&Vertex{1,2} //gets a *Vertex
)
(*p).X=2
p.X=2  // or directly permitted
```

> Arrays - Fixed length 
```go
var a [10]int
var b [2]string
```


>Slices : A slice does not store any data, are like references to arrays  ;append() ;  len() cap()  , empty ->nil , make
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
var m = make(map[string]Vertex)
var n= map[string]int{}
var n map[string]int          //needs to be initialized
k:=map[int]int{ 1:2,3:3}      //only inside func
k[4]=4                        // add key-value 
delete(k,4)                   //delete key-value
k=k[4]                        // returns default type value, no error thrown
val,ok :=k[4]                 // ok true if present else false(val=0/""/false)

```


## 4. Methods and interfaces : 

> Methods - kinda like operator overloading on types
```go
type Vertex struct {
	X, Y float64
}
//func (special rec arg) name([args list]) [return type]{}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//declare with ptr to be able to modify value, convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver. 
func (v *Vertex) Scale(f float64) {	
    v.X = v.X * f
	v.Y = v.Y * f
}
v := Vertex{3, 4}
fmt.Println(v.Abs())        //variable.method_name()


/*
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int). 
*/
```

>Interface :as a set of method signatures. 
```go
type Vertex struct{var x,y float64}
type myfloat float64
func (v Vertex) abs() float64{
    return math.Sqrt(v.x*v.x+v.y*v.y)
}

func (f myfloat) abs() float64{
    if f<0{return -f}
    return f
}
type Abser interface{
    abs() float64 
}
var a Abser
a=myfloat{6} 
fmt.Println(a.abs())
a=Vertex{2,3}
fmt.Println(a.abs())
/*
#interface -> set of method signature matching args and return types, it is implicit(no implements required)
#Can contain any type which implements all method signatures
#Can only be used to call that method, nothing else can be accessed
#holds (value,type)
# calling with nil value calls function with nil value (but has concrete type :- ptr)
#calling nil interface method is runtime error ! (
    var a abser
    a.abs() // error
    )  
# can have zero method empty interface
*/
//Type assertion : access underlying value 
val:=i.(int) // if i doesnt have that concrete type, trigger a panic !
t,ok:=i.(int) // ok=true if ctype found else false(with t= 0 value)

//Type Switch : i.(type)can only be used in type switch,. v will have the value 
switch v := i.(type) {
    case T:
        // here v has type T
    case S:
        // here v has type S
    default:
        // no match; here v has the same type as i
}
//Stringer interface
func (ip IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v",ip[0],ip[1],ip[2],ip[3])
}
```



## 5. Concurrency : Goroutines, Channels, Buffered Chn, Range&Close, Select, Def Select, Sync Mutex

>Go routines :  lightweight thread managed by the Go runtime. evaluation in curr goroutine, exec in new
```go
go f(x,y,z)
```

>Channels : sends and receives block until the other side is ready. 
```go
ch := make(chan int)
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty. 
ch := make(chan int, 100)
close(ch)
//Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
for i := range c {}
v, ok := <-ch // ok indicates whether channel closed or open
```
>Select : The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. 
The default case in a select is run if no other case is ready.
```go
//channels -> c & quit
select {
    case c <- x:
    	x, y = y, x+y
	case <-quit:
		fmt.Println("quit")
        return
    default:
        //try a send or receive without blocking: 
}
```
>sync.Mutex
```go
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}
c.mu.Lock()
// Lock so only one goroutine at a time can access the map c.v.
c.v[key]++
c.mu.Unlock()
```
