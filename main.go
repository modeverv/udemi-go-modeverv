package main

import (
	"fmt"
	//	"strings"
	//	"os/user"
	//	"time"
	//"strconv"
)

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func init() {
	fmt.Println("Init")
}

/*
 * コメント
 */
func buzz() {
	fmt.Println("Bazz")
}

func main() {
	/*
		// buzz()
		fmt.Println("Hello World", time.Now())
		buzz()
		fmt.Println(user.Current)
		var (
			i    int     = 1
			f64  float64 = 1.2
			s    string  = "test"
			t, f bool    = true, false
		)
		fmt.Println(i, f64, s, t, f)
		xi := 1
		xf64 := 1.2
		xt, xf := true, false
		fmt.Println(xi, xf64, xt, xf)

		fmt.Println(Pi, Username, Password)

		var (
			u8 uint8 = 255
		)
		fmt.Println(u8)
	*/
	/*
		fmt.Println("Hello World" + "aaa")
		fmt.Println("He"[0])
		fmt.Println(string("He"[0]))
		var s = "Hel"
		fmt.Println(strings.Replace(s,"H","X",1))
		fmt.Println(`test
					  test
		test`)
	*/
	/*
		t, f := true, false
		fmt.Printf("%T %v \n", t, t)
		fmt.Printf("%T %v \n", f, f)
		fmt.Println("aadsfaab")
	*/
	/*
	// ?
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %F\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)
	
	var s string = "14"
	//z = int(s)
	i,_ := strconv.Atoi(s)
    fmt.Printf("%T %v %d\n", i, i, i)
	
	h := "Hello World"
	fmt.Println(string(h[0]))
	*/
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b = [2]int{100,200}
	fmt.Println(b)

	var c = []int{100,200}
	c = append(c, 300)
	fmt.Println(c)
}
