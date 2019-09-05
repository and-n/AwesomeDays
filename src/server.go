package main

import (
	funcs "../src/functions"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	count = 0
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	count = funcs.LoadOldCount()
	Info = log.New(os.Stdout,
		"INFO: ",
		log.LstdFlags)
	Error = log.New(os.Stderr,
		"ERROR: ",
		log.LstdFlags|log.Lshortfile)
}

func main() {
	//scanner := bufio.NewReader(os.Stdin)
	//fmt.Println("need port")
	//port, _ := scanner.ReadString('\n')
	port := "8088"
	Info.Println("Server started!!")
	Info.Println("Days without Sergey's accident: ", count)
	startCounters()

	http.HandleFunc("/help", getHelp)
	http.HandleFunc("/days", getDays)

	err := http.ListenAndServe(":"+strings.Trim(port, "\n"), nil)

	if err != nil {
		Error.Fatal("server error ", err)
	}
}

func getDays(writer http.ResponseWriter, request *http.Request) {
	res := request.URL.Query().Get("hash")
	if request.Method == "GET" {
		_, err := fmt.Fprint(writer, strconv.Itoa(count))
		if err != nil {
			Error.Println("converter error", err)
		}
	} else if request.Method == "DELETE" && res == "sr321" {
		count = 0
		writer.WriteHeader(200)
	} else {
		writer.WriteHeader(400)
	}

}

func getHelp(w http.ResponseWriter, r *http.Request) {
	Info.Println(w, "Days without Sergey's accident!")
}

func startCounters() {
	go func() {
		for {
			hour, min, _ := time.Now().Clock()
			if hour == 0 && min == 0 {
				count++
				break
			} else {
				time.Sleep(1 * time.Minute)
			}
		}
		go func() {
			for {
				time.Sleep(24 * time.Hour)
				count++
				funcs.SaveCountToFile(count)
			}
		}()
	}()
}
