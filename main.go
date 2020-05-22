package main

import "github.com/cdutwhu/n3-frontend-test/host"

func main() {
	done := make(chan string)
	go host.HTTPAsync()
	<-done
}
