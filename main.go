package main

import (
	"fmt"
	"net/http"
)

var fullGolang = "Watch Nana's Super Go Full Course"
var crashGolang = "Watch Go Crash Course"
var rewardDessert = "Eat some chocolate"
var taskItems = []string{fullGolang, crashGolang, rewardDessert}

func main() {

	fmt.Println("####### Welcome to our TO DO list app #######")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTaks)

	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello User to our To Do List App!"
	fmt.Fprintln(writer, greeting)
}

func showTaks(writer http.ResponseWriter, request *http.Request) {

	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}
