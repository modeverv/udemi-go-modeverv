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
	/*
		var a [2]int
		a[0] = 100
		a[1] = 200
		fmt.Println(a)

		var b = [2]int{100,200}
		fmt.Println(b)

		var c = []int{100,200}
		c = append(c, 300)
		fmt.Println(c)
	*/
	/*
		n := []int{1,2,3,4,5}
		fmt.Println(n)
		fmt.Println(n[2])
		fmt.Println(n[2:4])
		fmt.Println(n[:2])
		fmt.Println(n[2:])
		fmt.Println(n[:])
		n[2] = 100
		fmt.Println(n)
		var board = [][]int {
			{0,1,2},
			{3,4,5},
			{6,7,8},
		}
		fmt.Println(board)
		n = append(n,100,200,300,400)
		fmt.Println(n)
	*/
	/*
		n := make([]int,3,5)
		fmt.Printf("len=%d cap=%d value=%v\n",len(n),cap(n),n)
		n = append(n,0,0)
		fmt.Printf("len=%d cap=%d value=%v\n",len(n),cap(n),n)
		n = append(n,1,2,3,4,5)
		fmt.Printf("len=%d cap=%d value=%v\n",len(n),cap(n),n)

		a := make([]int,3)
		fmt.Printf("len=%d cap=%d value=%v\n",len(a),cap(a),a)

		b := make([]int,0)
		fmt.Printf("len=%d cap=%d value=%v\n",len(b),cap(b),b)
		var c []int
		fmt.Printf("len=%d cap=%d value=%v\n",len(c),cap(c),c)
	*/
	/*
	   c := make([]int,5)
	   //[0 0 0 0 0 0]
	   //[0 0 0 0 0 0 1]
	   //[0 0 0 0 0 0 1 2]
	   //[0 0 0 0 0 0 1 2 3]
	   //[0 0 0 0 0 0 1 2 3 4]
	   //	c := make([]int,0,5)
	   //[0]
	   //[0 1]
	   //[0 1 2]
	   //[0 1 2 3]
	   //[0 1 2 3 4]
	   	for i := 0 ; i < 5; i++ {
	   		c = append(c,i)
	   		fmt.Println(c)
	   	}
	*/
	/*
		m := map[string]int{"apple": 100, "banana": 200}
		fmt.Println(m)
		fmt.Println(m["apple"])
		m["banana"] = 300
		m["new"] = 500
		fmt.Println(m)
		fmt.Println(m["nothing"])

		v,ok := m["apple"]
		fmt.Println(v,ok);
		v2,ok2 := m["nothing"]
		fmt.Println(v2,ok2)

		m2 := make(map[string]int)
		m2["pc"] = 5000
		fmt.Println(m2)
		// ?????
		//var m3 map[string]int

		var s[]int
		if s == nil {
			fmt.Println("nil")
		}
	*/
	/*
		b:=[]byte{72,73}
		fmt.Println(b)
		fmt.Println(string(b))
		//http://www.acii-code.com
	*/
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(3)

	func(x int) {
		fmt.Println("inner func", x)
	}(4)

}

//func add(x,y int) int {
func add(x int, y int) (int, int) {
	fmt.Println("add function")
	fmt.Println(x + y)
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	// reurn	// naked return
	return result
}

func convert(price int) float64 {
	return float64(price)
}
