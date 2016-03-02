//Packages
//Every Go program is made up of packages.
//Programs start running in package main.
//This program is using the packages with import paths "fmt" and "math/rand"
package main
//Imports
import (
    "fmt"
    "time"
    "math/rand"
    "math"
    "math/cmplx"    
)

//var c , python , java bool
//Functions 
//Notice that the type comes after the variable name.
//func add(x int, y int) int {.....
//When two or more consecutive named function parameters share a type, you can omit the type from all but the last. 
func add(x , y int) int {
    return x + y
}
//Multiple results
func swap(x, y string)(string , string){
    return y, x
}
//Named return values
func split(sum int)(x,y int){
    x = sum*4/9
    y = sum-x
    return
}
//Basic types
var (
    ToBe bool =false
    MaxInt64 uint64 = 1<<64-1
    MaxInt32 uint32 = 1<<32-1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)
const Pi=3.14
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
func needInt(x int) int { return x*10+1}
func needFloat(x float64) float64{
	return x*0.1
}
func main() {
        fmt.Println("This time is" , time.Now())
	fmt.Println("Hello, 世界")
        fmt.Println("My favorite number is" , rand.Intn(10))
        fmt.Println("Now you have %g problems." , math.Sqrt(7))
//Exported names
//When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package. 
        fmt.Println(math.Pi)
//Functions continued
        fmt.Println(add(42, 13))
//Multiple results
        a,b :=swap("hello", "world")
        fmt.Println(a,b)
//Named return values
        fmt.Println(split(17))
//Variables
//The var statement declares a list of variables; as in function argument lists, the type is last. 
//      var i int
//      var c, python, java bool 
//Variables with initializers
//Short variable declarations
//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.        
        var i, j int = 1, 2
        c,python,java :=true,false,"no!"
        k :=3
        fmt.Println(i, j , k , c , python , java)       
//Basic types
//bool
//string
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte // alias for uint8
//rune // alias for int32
     // represents a Unicode code point
//float32 float64
//complex64 complex128
        const f = "%T(%v)\n" 
        fmt.Printf(f, ToBe, ToBe)
        fmt.Printf(f, MaxInt64,MaxInt64)
        fmt.Printf(f, MaxInt32,MaxInt32)
        fmt.Printf(f,z,z)
//Zero values        
        var i1 int
        var f1 float64
        var b1 bool
        var s string
        fmt.Printf("%v %v %v %q\n",i1,f1,b1,s)
//Type conversions
//var i int = 42
//var f float64 = float64(i)
//var u uint = uint(f)
//i := 42
//f := float64(i)
//u := uint(f)
        x,y :=3,4
        var f2 float64=math.Sqrt(float64(x*x+y*y))
        var z uint=uint(f2)
        fmt.Println(x,y,z)
//Type inference
//var i int
//j := i // j is an int
//i := 42           // int
//f := 3.142        // float64
//g := 0.867 + 0.5i // complex128
        v:=42 // change me!
        fmt.Printf("v is of type %T\n",v)
        v1:=2.4
        fmt.Printf("v1 is of type %T\n",v1)
//Constants
//Constants cannot be declared using the := syntax. 
	const World = "世界"
        fmt.Println("Hello",World)
        fmt.Println("Hello",Pi,"Day")
        
        const Truth = true
        fmt.Println("Go rules?", Truth)

//Numeric Constants
        fmt.Println(needInt(Small))
        fmt.Println(needFloat(Small))
        fmt.Println(needFloat(Big))
}