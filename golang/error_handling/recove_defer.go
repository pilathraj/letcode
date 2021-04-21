package main
import "fmt"

func main(){
  defer func(){
    if r := recover(); r!=nil{
      fmt.Println(r)
    }
  }()
  a, b := 10, 0
  c := a/b
  fmt.Println(c)
}
