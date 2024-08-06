package main

import (
	"fmt"
	"reflect"
	"math"
)

// import "golang.org/x/tour/pic"

// func Pic(dx, dy int) [][]uint8 {
// }

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

type Vertex1 struct {
	Lat, Long float64
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

var m map[string]Vertex1

var m1 = map[string]Vertex1{
	"Bell Labs": Vertex1{
		40.68433, -74.39967,
	},
	"Google": Vertex1{
		37.42202, -122.08408,
	},
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	fmt.Println("Variable declaration")
	// variable declaration
	 var var1='z';
	 // checking the type of the variable
    fmt.Println(reflect.TypeOf(var1));
	var2 := 1

	//printing two variables in same println
	fmt.Println("var2" , var2);
	//checking the type of variable
    fmt.Println(reflect.TypeOf(var2))
	fmt.Println("\nLoops\nfor-loop");

	/* looping 
		statements */
	var i = 0
	//for loop
	for i=0;i<5;i++{
		fmt.Print("i-");fmt.Println(i);
	}
	s :=""
	fmt.Println("\nwhile-loop")
	//while loop wth sprintf
	for i > var2{
		s+= fmt.Sprintf("i is --- %d",i);
		s+="\n"
		i=i-1;
	}
	fmt.Print(s);

	//functions
	fmt.Println("\nFunctions")
	fmt.Print(sum(2,3))





	defer fmt.Println("world")



	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j




	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)



	pr := &v
	pr.X = 1e9
	fmt.Println(v)



	fmt.Println(v1, p, v2, v3)



	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)



	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var sw []int = primes[1:4]
	fmt.Println(sw)



	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	
	ar := names[0:2]
	br := names[1:3]
	fmt.Println(ar, br)

	br[0] = "XXX"
	fmt.Println(ar, br)
	fmt.Println(names)



	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	sa := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sa)



	
	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])


	fmt.Println(m)




	mq := make(map[string]int)

	mq["Answer"] = 42
	fmt.Println("The value:", mq["Answer"])

	mq["Answer"] = 48
	fmt.Println("The value:", mq["Answer"])

	delete(mq, "Answer")
	fmt.Println("The value:", mq["Answer"])

	v, ok := mq["Answer"]
	fmt.Println("The value:", v, "Present?", ok)





	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))





	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}


	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

func sum(num1 int ,num2 int) int{
	return num1+num2
}