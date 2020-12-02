package main

import (
  "fmt"
  "log"
  "os"
  "bufio"
  "regexp"
  "strings"
  "strconv"
)

func checkPasswordValidityForChallange1(password string,searchkey string,minOccurance int,maxOccurence int) bool {

  regx := regexp.MustCompile(searchkey)

  matches := regx.FindAllStringIndex(password, -1)
  totalOccurance := len(matches)
  if(totalOccurance >= minOccurance && totalOccurance <= maxOccurence) {
    return true;
  }

  return false
}

func checkPasswordValidityForChallange2(password string,searchkey string,minOccurance int,maxOccurence int) bool {

    charArray :=  strings.Split(password,"")
  fmt.Println("check length",searchkey)
  occuranceAtMinPos := charArray[minOccurance]
  occuranceAtMaxPos := charArray[maxOccurence]
  //
  if(occuranceAtMinPos == searchkey && occuranceAtMaxPos != searchkey || occuranceAtMinPos != searchkey && occuranceAtMaxPos == searchkey  ) {
   return true
  }

  return false
}


func parseMinMaxOccurance(data string) (int,int) {
  splitMinMax := strings.Split(data,"-")
  minStrToIn,_ := strconv.Atoi(splitMinMax[0])
  maxStrToIn,_ := strconv.Atoi(splitMinMax[1])
  return minStrToIn,maxStrToIn
}

func main() {

  solveChallange1()
  solveChallange2()
}

func solveChallange2() {
  file, err := os.Open("input")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  counter := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    parsedString := strings.Split(scanner.Text(), ":")

    rules := strings.Split(parsedString[0]," ")
    min,max := parseMinMaxOccurance(rules[0])
    searchKey := rules[1]
    if(checkPasswordValidityForChallange2(parsedString[1],searchKey,min,max)) {
      counter++
    }

  }
  fmt.Println("total correct passwords for challange 2 are",counter)

}

func  solveChallange1()  {
  file, err := os.Open("input")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  counter := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    parsedString := strings.Split(scanner.Text(), ":")

    rules := strings.Split(parsedString[0]," ")
    min,max := parseMinMaxOccurance(rules[0])
    searchKey := rules[1]
    if(checkPasswordValidityForChallange1(parsedString[1],searchKey,min,max)) {
      counter++
    }
  }

  fmt.Println("total correct passwords for challange 1 are",counter)

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
