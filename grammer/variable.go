package main

import "fmt"

func main(){
	variable()
	Datatype() // 대문자 선언 해야 export 가능
}


func variable(){
		// variable
		var a int;
		var f float32 = 11.
	
		a = 10
		f = 12.0
	
		fmt.Println(a);
		fmt.Println(f);
	
		var var1 = 15  //  Type inference
		var var2 = "Hello"
	
		fmt.Println(var1)
		fmt.Println(var2)
	
		// constants
		const con string = "Hello World"
		const num int = 1
	
		fmt.Println(con)
		fmt.Println(num)
	
		const const1 = 10  //  Type inference
		const const2 = 20
	
		fmt.Println(const1 + const2)
	
	
		const (
			APPLE = iota
			GRAPE
			ORANGE
		)
	
		fmt.Println(GRAPE)	
		fmt.Println("------------variable------------")
}

// keyword
// break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var



