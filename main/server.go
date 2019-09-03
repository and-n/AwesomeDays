package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var count = 0

func main() {
	//scanner := bufio.NewReader(os.Stdin)
	//fmt.Println("need port")
	//port, _ := scanner.ReadString('\n')
	port := "8088"
	fmt.Println("Server started!!")

	http.HandleFunc("/help", getHelp)
	http.HandleFunc("/days", getDays)

	go func() {
		for {
			hour, min, _ := time.Now().Clock()
			if hour == 0 && min == 0 {
				time.Sleep(1 * time.Minute)
			} else {
				break
			}
		}
		go func() {
			for {
				time.Sleep(24 * time.Hour)
				count++
			}
		}()
	}()

	err := http.ListenAndServe(":"+strings.Trim(port, "\n"), nil)

	if err != nil {
		fmt.Println("server error ", err)
	}
}

func getDays(writer http.ResponseWriter, request *http.Request) {
	res := request.URL.Query().Get("hash")
	if request.Method == "GET" {
		_, err := fmt.Fprint(writer, strconv.Itoa(count))
		if err != nil {
			fmt.Println("converter error", err)
		}
	} else if request.Method == "DELETE" && res == "sr321" {
		count = 0
		writer.WriteHeader(200)
	} else {
		writer.WriteHeader(400)
	}

}

func getHelp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Days without Sergey's accident!")
}
