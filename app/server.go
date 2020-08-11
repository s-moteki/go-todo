package main

import (
	"github.com/s-moteki/go-todo/app/infrastructure"
)

func main() {
	_ = infrastructure.Router.Run()
}
