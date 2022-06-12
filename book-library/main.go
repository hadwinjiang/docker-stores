package main

import (
	"log"
	"os"

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
	log.Println("Application running in environment: ", os.Getenv("RUNTIME_SETUP"), " and on port: ", os.Getenv("PORT"))

	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static")
	router.Run()

}
