# Golang

## Cheatsheet 
### Resources :
- <a href="https://www.youtube.com/watch?v=sX8r6zATHGU">Philosophy of go(Yt)</a> 
- <a href="https://tour.golang.org/welcome/1">A tour of go</a>
- <a href="https://blog.golang.org/slices-intro">Slices</a>
- <a href="https://golang.org/doc/effective_go.html#introduction">Effective go code</a>
- <a href="https://golang.org/pkg/">Packages</a>

- <a href="https://blog.alexellis.io/golang-writing-unit-tests/">Unit Tests</a>

- <a href="https://blog.golang.org/using-go-modules">Dependency Management : Modules</a>

<details>
<summary>

## 1. Structure (Package, import, exports, variables, functions): 
</summary>

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
    "math"
)
//Import declaration          Local name of Sin
import   "lib/math"         //math.Sin
import m "lib/math"         //m.Sin
import . "lib/math"         //Sin

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
var a=1 //auto inference 
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
func function_name( [parameter list] ) [return_types]{
   body of the function
}

//naked fucntions : named return parameters with empty return(naked) statement, return named params
func name(params) (a int){
    return 
}

```

</details><br>

<details>
<summary>

## 2. Flow Control (for, if, else, switch &defer): 
</summary>

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
        break
    case 2,3,4:
    break         //multiple cases sep by comma
    default:
}
switch {            // if left empty, chooses true case top to bottom
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
}
//type switch refer interfaces
```

> defer : the deferred function is scheduled to be run immediately before the function executing the defer returns. The deferred call's arguments are evaluated immediately.
```go
defer fmt.Println("world")
fmt.Println("hello")
//output hello \n world \n
//multi defered func called in stack kinda lifo order
//Defer is commonly used to simplify functions that perform various clean-up actions.
```

</details><br>

<details>
<summary>

## 3. More types : ptr, Structs, slices and maps
</summary>

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

</details><br>
<details>
<summary>

## 4. Methods and interfaces : 
</summary> 

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

</details><br>

<details>
<summary>

## 5. Concurrency : Goroutines, Channels, Buffered Chn, Range&Close, Select, Def Select, Sync Mutex
</summary>

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
for i := range ch {}
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
    case <-quit:{
        fmt.Println("quit")
        return
    }
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
func(){
	c.mu.Lock()
	defer c.mu.Unlock() //clean in case u need to return value
    return c.v[key]
}
```

</details><br>
<details>
<summary>

## 6. Standard Packages : Go 
</summary>
<a href="https://golang.org/pkg/fmt/">"fmt"</a></br>
<a href="https://golang.org/pkg/sync/">"sync"</a></br>
<a href="https://golang.org/pkg/fmt/">"fmt"</a></br>
<a href="https://golang.org/pkg/strings/">"strings"</a></br>
<a href="https://golang.org/pkg/strconv/">"strconv"</a></br>
<a href="https://golang.org/pkg/os/">"os"</a></br>
<a href="https://golang.org/pkg/http/">"http"</a>
</details>
<br>

<details>
<summary>

## 7. Unit Tests : has testing package
</summary>
 
https://blog.alexellis.io/golang-writing-unit-tests/
```go
import testing //built in
// name the file name_test.go
func Test<name>(t * testing.T){
    var tests = []struct { //tabular tests
        a, b,want int
    }{
        {0, 1, 0},
        {1, 0, 0},
    }
//t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.
    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {  //naming tests 
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

/*
go test -v
//to check coverage
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out //gives html file to visualize
*/
```
</pre>
</details>

<br>


<details>
<summary>

## 8. Dependency Management : Modules 
</summary>

Modules : https://blog.golang.org/using-go-modules , 

project structure https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/ 
 
```go

```
</details><br>

<details open>
<summary>

## 9. HTTP & gPRC server/client 
</summary>

Reason : https://www.youtube.com/watch?v=u4LWEXDP7_M  

protoc buff :https://tutorialedge.net/golang/go-protocol-buffer-tutorial/ , 

grpc : https://tutorialedge.net/golang/go-grpc-beginners-tutorial/#video-tutorial 

>.proto file  : cmd to generate .pb.go files

>$protoc --go_out=. --go-grpc_out=. filename.proto
```go
syntax="proto3";
package main;

message HelloRequest{
    string name=1; //type name= size or ex from which it is inferred 
}
message HelloReply {
    string message = 1; //message can be nested
}
  
service Greeter{
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}
```
>Server: define type server, implement function(rpc) ; create, register new server, listen on tcp port. 
```go
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement GreeterServer.
type server struct {
	UnimplementedGreeterServer
}

// SayHello implements GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

```
>client : dial the address, create hrpc connection, create context, call rpc
```go
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "Rakshith"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

```
</details><br>
