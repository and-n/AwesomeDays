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
	count int
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	log.SetOutput(os.Stdout)
	count = funcs.LoadOldCount()

	Info = log.New(os.Stdout,
		"INFO: ",
		log.LstdFlags)
	Error = log.New(os.Stderr,
		"ERROR: ",
		log.LstdFlags|log.Lshortfile)
}

func main() {
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
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Methods", "GET, DELETE, OPTIONS")
	if request.Method == "GET" {
		_, err := fmt.Fprint(writer, strconv.Itoa(count))
		if err != nil {
			Error.Println("converter error", err)
		}
	} else if request.Method == "DELETE" && res == "sr321" {
		Info.Println("Days reset")
		count = 0
		funcs.SaveCountToFile(count)
		writer.WriteHeader(200)
	} else if request.Method == "OPTIONS" {
		_, _ = fmt.Fprint(writer, "GET, DELETE, OPTIONS")
	} else {
		Info.Printf("unknown request from %s \n", request.Host)
		writer.WriteHeader(400)
	}

}

func getHelp(w http.ResponseWriter, r *http.Request) {
	Info.Println("heeeeelp!")
	if r.Method == "GET" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		_, err := fmt.Fprint(w, "Days without Sergey's accident!")
		if err != nil {
			Error.Println("", err)
		}
	} else {
		w.WriteHeader(400)
	}
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
				//time.Sleep(2 * time.Second)
				count++
				funcs.SaveCountToFile(count)
			}
		}()
	}()
}
