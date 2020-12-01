package main

import (
  "fmt"
  _ "fmt"
  "encoding/json"
  "io/ioutil"
)

func main()  {

  inputs := loadInput()

  firstNum, secondNum := solveChallange1(inputs)
  fmt.Println("the result for 1st puzzle is ",firstNum * secondNum)
  firstNum, secondNum,thirdNum := solveChallange2(inputs)
  fmt.Println("the result for third puzzle is ",firstNum * secondNum * thirdNum)


}

/**

Search two number to form a single given num
data [6,2,3,4]
e.g 10 = ^+4
 */
func searchTwoNumComposition(inputs []int, numberToSearch int) (int,int) {
  for _,s := range inputs {
    remainingHalf := numberToSearch - s


    if(itemExists(inputs,remainingHalf)) {

      return s,remainingHalf
    }
  }
  return 0,0
}
func solveChallange2(inputs []int) (int,int,int) {
  for _,sampleCase := range inputs {
    remainingHalf := 2020 - sampleCase

    if(!itemExists(inputs,remainingHalf)) {

      firstNum,secondNum := searchTwoNumComposition(inputs,remainingHalf)
      if firstNum > 0 && secondNum > 0 {
        return firstNum,secondNum,sampleCase
      }
    }

  }
  return 0,0,0
}


func solveChallange1(inputs []int) (int,int) {
  return searchTwoNumComposition(inputs,2020)
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
