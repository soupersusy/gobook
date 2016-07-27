package main
import "fmt"
func main()  {
  elements := map[string]map[string]string{
    fmt.Print("Enter an element symbol:")
    var input string
    fmt.Scanf ("%s", &input)
    "H": map[string]string{
      //note commas
      "name":"Hydrogen",
      "state":"gas",
    } ,
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
    "Li": map[string]string{
      "name":"Lithium",
      "state":"gas",
    },
  }
  ele := input
//  if el, ok := elements[ele]; ok {
    fmt.Println(ele["name"], ele["state"])
    }
}
