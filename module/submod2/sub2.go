package submod

import (
	"fmt"
	sb "github.com/rohitkolhe108/GoProjects/module/submod"
)

func SubPrint(){
	fmt.Print("sub")
}

func SubCall2(moduelName string){
	sb.SubCall("submodule2")
}