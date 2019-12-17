package main

import (
	_ "github.com/crusj/git-auto-pull/init"
)

func init() {
}
func main() {
	block := make(chan bool, 1)
	<-block
}
