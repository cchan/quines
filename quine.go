package main
import (
  "fmt"
  "strings"
)
func main () {
  str := "package main\nimport (\n  \"fmt\"\n  \"strings\"\n)\nfunc main () {\n  str := \"{CODE}\"\n  strEsc := strings.Replace(str, \"\\\\\", \"\\\\\\\\\", -1)\n  strEsc = strings.Replace(strEsc, \"\\n\", \"\\\\n\", -1)\n  strEsc = strings.Replace(strEsc, \"\\\"\", \"\\\\\\\"\", -1)\n  fmt.Println(strings.Replace(str, \"{CODE}\", strEsc, 1))\n}"
  strEsc := strings.Replace(str, "\\", "\\\\", -1)
  strEsc = strings.Replace(strEsc, "\n", "\\n", -1)
  strEsc = strings.Replace(strEsc, "\"", "\\\"", -1)
  fmt.Println(strings.Replace(str, "{CODE}", strEsc, 1))
}
