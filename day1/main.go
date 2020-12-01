package main

import (
  "fmt"
  _ "fmt"
  "encoding/json"
  "io/ioutil"
)

func main()  {

  inputs := loadInput()

  firstNum, secondNum := solveChallange(inputs)
  result := firstNum * secondNum
  fmt.Println("the result is ",result)


}

func solveChallange(inputs []int) (int,int) {
  for _,s := range inputs {
    remainingHalf := 2020 - s


    if(itemExists(inputs,remainingHalf)) {

      return s,remainingHalf
    }

  }
  return 0,0
}


func itemExists(inputs []int, searchNum int) bool {
  for _,s := range inputs {

    if(s == searchNum) {
      return true
    }
  }
  return false
}

func loadInput() []int {
  byteValue, _ := ioutil.ReadFile("input")

  var inputs []int

  json.Unmarshal(byteValue, &inputs)


  return inputs
}
