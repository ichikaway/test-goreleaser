package main

import "fmt"

var (
    version = "v"
    commit  = "none"
    date    = "unknown"
    builtBy = "unknown"
)

func main() {
  fmt.Printf("my app %s, commit %s, built at %s by %s \n", version, commit, date, builtBy)
}

