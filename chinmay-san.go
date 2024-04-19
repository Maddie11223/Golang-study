package main
import "fmt"
/* For exercises uncomment the imports below */
// import "strconv"
// import "encoding/json"

func main() {

    // a normal variable  whose address the pointer will store 
    var intData = 20 
    
    //declaration of a pointer 
    var intPointer *int

    //intPointer now points towards intData    
    intPointer = &intData 

    fmt.Println("what intData stores:", intData)
    fmt.Println("address of intData:", &intData)
    fmt.Println("what intPointer stores:", intPointer)


    //this updates the value of intData using dereferncing operator 
    *intPointer = 30

    fmt.Println("what intData now stores:", intData)
