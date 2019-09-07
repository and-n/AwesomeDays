package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName string = "count.val"

func LoadOldCount() int {
	log.Printf("load count from file %s", fileName)
	res, err := ioutil.ReadFile(fileName)
	if err == nil {
		log.Printf("loaded count %s", string(res))
		var val, ei = strconv.Atoi(strings.TrimSpace(string(res)))
		if ei != nil {
			log.Println(ei)
		}
		return val
	} else {
		log.Print(err.Error())
	}
	return 0
}

func SaveCountToFile(count int) {
	log.Printf("save count %d to file %s \n", count, fileName)
	f, _ := os.Create(fileName)
	_, err := fmt.Fprint(f, count)
	if err != nil {
		log.Println(err)
	}
	_ = f.Close()
}
