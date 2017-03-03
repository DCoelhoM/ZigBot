package main

import (
  "strings"
  "regexp"
  "net/url"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func calc(expression string) string {
  const url_calc = "https://newton.now.sh/simplify/"
  const error_msg = "Something went wrong!"

  match, _ := regexp.MatchString("^[0-9+-/*^() ]+$", expression)
  if (!match) {
    return error_msg
  }

  r := strings.NewReplacer(" ", "","(", "%28", ")", "%29", "*", "%2A", "+", "%2B", "-", "%2D", "/", "%2F", "^", "%5E")
  u, err := url.Parse(url_calc + r.Replace(expression))
  if err != nil {
    return error_msg
  }

  resp, err := http.Get(u.String())
  if err != nil {
    return error_msg
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  type Simplify struct {
    Operation string `json:"operation"`
    Expression string `json:"expression"`
    Result string `string:"result"`
  }
  var simp Simplify
  err = json.Unmarshal(body, &simp)
  if err != nil {
    return error_msg
  }
  return simp.Result
}
