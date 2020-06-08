package main

import "github.com/gin-gonic/gin"

var DefaultEngine = gin.Default()

func main() {
	InitAuthApi().InitApi()
	DefaultEngine.Run()
}
