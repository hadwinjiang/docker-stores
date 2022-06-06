package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)

	// log.Println("Starting the application, listening on :8080...")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static")
	router.Run()

}
