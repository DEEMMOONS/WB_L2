package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func Parse(str string) (string, error) {
  var res []rune
  if len(str) == 0 {
    return "", nil
  }
  if unicode.IsDigit(rune(str[0])) {
		return "", fmt.Errorf("start with num")
	}

  for _, char := range []rune(str) {
    if !unicode.IsDigit(char) {
      res = append(res, char)
    } else {
      if res[len(res) - 1] == '\u005C' {
        res[len(res) - 1] = rune(0)
        res = res[:len(res) - 1]
        if len(res) < 2 || res[len(res) - 1] != '\u005C' {
          res = append(res, char)
          continue
        }
      }
      num, _ := strconv.Atoi(string(char))
      for i := 1; i < num; i++ {
        res = append(res, res[len(res) - 1])
      }
    }
  }
  return string(res), nil
}

func main() {
  var str string
  fmt.Scan(&str)
  ans, _ := Parse(str)
  fmt.Println(ans)
}
