package main
import (
    "fmt"
)

type myType struct {i int}
type myChildType struct{
    myType // promoted field
    j int
}

type myInterface interface{
    Get() int
}

// myType Get
func (p *myType) Get() int{
    return p.i
}

func a_function_accept_myinterface(x myInterface){
}

func main(){
    mct := myChildType{j: 0}
    fmt.Printf("%d\n", mct.Get())
    fmt.Printf("%d\n", mct.i)
    fmt.Printf("%d\n", mct.myType.i)
    a_function_accept_myinterface(&mct)
}
