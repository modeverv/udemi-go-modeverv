package main

import(
	"fmt"
	"os/user"	
	"time"
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
	// buzz()
	fmt.Println("Hello World",time.Now())
	buzz()
	fmt.Println(user.Current)
	var i int = 1
	var f64 float64 = 1.2
	var s string = "test"
	var t,f bool = true,false
	
	fmt.Println(i,f64,s,t,f)
}
