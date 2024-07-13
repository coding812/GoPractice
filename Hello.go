package main

import (
	"fmt" 
   "time"
   "math"
   "math/cmplx"
	"rsc.io/quote"
)
var c, python, java = true , false , "Hello!"

func add(x , y int) int{
   return x+y
}
const Pi float32 = 3.14

func swap(x , y string) (string, string){
   return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

type Vertex struct {
	X int
	Y int
}

func main() {
   defer fmt.Print(" World") // This will only run when the surrounding function returns
      
   // These call quotes from the Quotes module
   fmt.Println(quote.Hello())
   fmt.Println(quote.Go())
   fmt.Println(quote.Glass())
   fmt.Println(quote.Opt())

   fmt.Println(add(100,240))

   a, b := swap("hello", "world")
   fmt.Println(a, b);

   fmt.Println(split(20))

   var i int
   k := 3
   fmt.Println(k, i, c, python, java)

   fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

   var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

   flt := 3.44
   var badConvert int = int(flt) // apparently this isn't bad
   fmt.Println(badConvert)

   integer := 5
   var tryAgain float32 = float32(integer)
   fmt.Println(tryAgain)

   var isNull = [5]int{}
   slice := isNull[1:4]

   fmt.Println(isNull)
   fmt.Println(slice)

   // sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

   sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
   fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
   t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

   
   v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)


   fmt.Print("Hello") // keep this at the bottom of the main function
}
