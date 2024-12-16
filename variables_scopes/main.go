/* The top-level scope is the package scope. A scope can have
child scopes within it. There are a few ways a child scope gets defined; the easiest way to think about
this is that when you see {, you are starting a new child scope, and that child scope ends when you
get to a matching }. The parent-child relationship is defined when the code compiles, not when the
code runs. When accessing a variable, Go looks at the scope the code was defined in. If it can’t find a
variable with that name, it looks in the parent scope, then the grandparent scope, all the way until it
gets to the package scope. It stops looking once it finds a variable with a matching name or raises an
error if it can’t find a match. */

package main

import "fmt"

var level = "pkg"

func main() {
	fmt.Println("Main starts: ", level)
	//create a shadow variable
	level := 42
	if true {
		fmt.Println("Block starts: ", level)
		funcA()
	}
	fmt.Println("main ends: ", level)
}

func funcA() {
	fmt.Println("funcA starts: ", level)
}
