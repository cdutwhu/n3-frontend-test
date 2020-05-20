package main

import "github.com/nsip/n3-frontend-test/host"

func main() {
	done := make(chan string)
	go host.HTTPAsync()
	<-done
}
