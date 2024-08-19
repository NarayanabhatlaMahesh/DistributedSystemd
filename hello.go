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

func (v Vertex) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
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

type MyFloat float64

func (f MyFloat) Abs1() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
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

	v = Vertex{3, 4}
	fmt.Println(Abs(v))



	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs1())




	v = Vertex{3, 4}
	v.Scale(10)
	fmt.Println(Abs2())




	Scale(&v, 10)
	fmt.Println(Abs2(v))





	v = Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p = &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)





	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p = &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))





	v = &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs2())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs2())





	var a Abser
	f := MyFloat(-math.Sqrt2)
	v = Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())



	var i I = T{"hello"}
	i.M()



	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()




	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()



	describe(i)
	i.M()


	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)








	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)






	do(21)
	do("hello")
	do(true)
	
	



	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)




	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))





	go say("world")
	say("hello")





	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)





	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)





	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}





	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci1(c, quit)






	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}







	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))






	
}


func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func sum(num1 int ,num2 int) int{
	return num1+num2
}