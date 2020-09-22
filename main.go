package main

import "fmt"

var (
    version = "v0.0.1"
    commit  = "none"
    date    = "unknown"
    builtBy = "unknown"
)

func main() {
  fmt.Printf("my app %s, commit %s, built at %s by %s", version, commit, date, builtBy)
}

