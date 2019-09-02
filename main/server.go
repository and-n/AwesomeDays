package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var count = 0

func main() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("need port")
	port, _ := scanner.ReadString('\n')

	fmt.Println("Server started!!")

	http.HandleFunc("/help", getHelp)
	http.HandleFunc("/days", getDays)
	http.HandleFunc("/reset", resetHandler)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			count++
		}
	}()
	err := http.ListenAndServe(":"+strings.Trim(port, "\n"), nil)

	if err != nil {
		fmt.Println("server error ", err)
	}
}

func resetHandler(writer http.ResponseWriter, request *http.Request) {
	count = 0
}

func getDays(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprint(writer, strconv.Itoa(count))
	if err != nil {
		fmt.Println("converter error", err)
	}
}

func getHelp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Days without Sergey's accident!")
}
