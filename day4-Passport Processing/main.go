package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "regexp"
  "strconv"
  "strings"
)

func main()  {

  solveChallange1()
  solveChallange2()
}

func solveChallange2() {
  file, err := os.Open("input")
  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)

  validPass := 0
  tempVal := []string{}

  for scanner.Scan() {

    if scanner.Text() != "" {
      lineFragments := strings.Split(scanner.Text(), " ")

      for _, val := range lineFragments {
        key := strings.Split(val, ":")

        fmt.Println("key1",key[1])
        passportVal := key[1]
        if(key[0] == "byr" ) {
          year,_ := strconv.Atoi(passportVal)
          if(year < 1920 || year > 2002) {
            continue
          }
        }

        if(key[0] == "iyr" ) {
          year,_ := strconv.Atoi(passportVal)
          if(year < 2010 || year > 2020) {
            continue
          }
        }

        if(key[0] == "eyr" ) {
          year,_ := strconv.Atoi(passportVal)
          if(year < 2020 || year > 2030) {
            continue
          }
        }

        if(key[0] == "hgt" ) {
          units := passportVal[len(passportVal)-2:]
          n, _ := strconv.Atoi(passportVal[:len(passportVal)-2])

          if(units == "cm" && 150 < n && n > 193) {

            continue
          }

          if(units == "in" && 59  < n && n > 76) {
            continue
          }

        }

        if(key[0] == "hcl" ) {
          match, _ := regexp.MatchString(`^#[\da-f]{6}$`, passportVal)
          if(!match) {
          continue
          }
        }

        if(key[0] == "ecl"  && passportVal != "amb" &&  passportVal != "blu" && passportVal != "brn" && passportVal != "gry" && passportVal != "hzl" && passportVal != "oth"  ) {
          //passportVal
          continue
        }

        if(key[0] == "pid") {
          match, _ := regexp.MatchString(`^\d{9}$`, passportVal)
          if(!match) {
            continue
          }
        }
        tempVal = append(tempVal, key[0])

      }
    }

    if scanner.Text() == "" {

      a := []string{"cid","iyr","eyr","hgt","hcl","ecl","pid","byr"}
      diff :=compare(a,tempVal)
      if(len(diff) == 0) {
        validPass++
      }

      if(len(diff) ==1) {
        fmt.Println(diff,"with",diff[0])

        if(diff[0] == "cid") {
          validPass++
        }
      }

      tempVal = []string{}

    }

  }

  fmt.Println("result for 2",validPass)
}


func solveChallange1() {

  file, err := os.Open("input")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  validPass := 0
  tempVal := []string{}

  for scanner.Scan() {

    data := scanner.Text()

    if scanner.Text() != "" {
      lineFragments := strings.Split(data, " ")

      for _, val := range lineFragments {
        key := strings.Split(val, ":")
        tempVal = append(tempVal, key[0])

      }
    }

    if scanner.Text() == "" {

      a := []string{"cid","iyr","eyr","hgt","hcl","ecl","pid","byr"}
      diff :=compare(a,tempVal)
      if(len(diff) == 0) {
        validPass++
      }

      if(len(diff) ==1) {
        fmt.Println(diff,"with",diff[0])

        if(diff[0] == "cid") {
          validPass++
        }
      }

      tempVal = []string{}

    }

  }

  fmt.Println("result for 1",validPass)
}

//
func compare(a, b []string) []string {
  for i := len(a) - 1; i >= 0; i-- {
    for _, vD := range b {
      if a[i] == vD {
        a = append(a[:i], a[i+1:]...)
        break
      }
    }
  }
  return a
}
