package submod

import "fmt"

func SubPrint(){
	fmt.Print("sub1")
}

func SubCall(moduleName string){
	fmt.Print("Call from "+ moduleName +"submodule1")
}

