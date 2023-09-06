package main

import "fmt"

//1 add . = add files
//2 git commit -m "message"  = commit files
//3 git push origin master = push files to github
func sub(a,b int) int{
	return a-b
}
func sum (a,b int) int{
	return a+b
}

func mul(a,b int)int{
	return a*b
}

func main() {
	fmt.Println("Hello Github")
	fmt.Println(sum(1,2))
	fmt.Println(sub(1,2))
	fmt.Println(mul(1,2))
}
