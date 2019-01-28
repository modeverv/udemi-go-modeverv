package main

import (
	"fmt"
	"strings"
//	"os/user"
//	"time"
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
	fmt.Println("Hello World" + "aaa")
	fmt.Println("He"[0])
	fmt.Println(string("He"[0]))
	var s = "Hel"
	fmt.Println(strings.Replace(s,"H","X",1))
	fmt.Println(`test
				  test
	test`)
}
