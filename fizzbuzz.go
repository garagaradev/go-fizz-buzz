package main
import "fmt"

func fizzBuzz(start, end int) {
  for i:=start;i<=end;i++{
    if i%3==0 && i%5==0 {
      fmt.Println("FizzBuzz")
    }else if i%3 == 0 {
      fmt.Println("Fizz")
    }else if i%5 == 0 {
      fmt.Println("Buzz")
    }else {
      fmt.Println(i)
    }
  } 
}

func main(){
  var start int
  var end int

  fmt.Print("Input the initial number:")
  fmt.Scan(&start)
  fmt.Print("Input the ending number:")
  fmt.Scan(&end)
  fizzBuzz(start,end)
}
