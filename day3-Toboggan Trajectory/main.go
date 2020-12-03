package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "strings"
)


func main() {

    solveChallange1()
  solveChallange2()
  }


  func solveChallange1()  {

    trees := traverseAndFindTrees(3,1)
    fmt.Println("total trees in challange 1 are ",trees)
  }

func solveChallange2()  {

  firstCasetrees := traverseAndFindTrees(1,1)
  secondCasetrees := traverseAndFindTrees(3,1)
  thirdCasetrees := traverseAndFindTrees(5,1)
  fourthCasetrees := traverseAndFindTrees(7,1)
  fifthCasetrees := traverseAndFindTrees(1,2)
  fmt.Println("total trees in challange 2 are ",firstCasetrees * secondCasetrees * thirdCasetrees * fourthCasetrees * fifthCasetrees)
  //fmt.Println("total trees in challange 2 are ",fifthCasetrees)
}


  func traverseAndFindTrees(right int,down int) int {
    file, err := os.Open("input")
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    iteration := 0
    trees := 0
    counter := 0
    for scanner.Scan(){

      if(counter > 0  &&  counter % down > 0 ) {

        counter ++;
        continue
      }
      fragments := strings.Split(scanner.Text(),"")
      index :=iteration % len(scanner.Text());
      val := fragments[index]
      if(val == "#") {

        trees++
      }

      counter ++;

      iteration += right
    }

    return trees



  }

