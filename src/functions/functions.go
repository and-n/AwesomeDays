package functions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const fileName string = "count.val"

func LoadOldCount() int {
	res, err := ioutil.ReadFile(fileName)
	if err == nil {
		var val, _ = strconv.Atoi(string(res))
		return val
	}
	return 0
}

func SaveCountToFile(count int) {
	f, _ := os.Create(fileName)
	_, err := fmt.Fprint(f, count)
	if err != nil {
		fmt.Println(err)
	}
	_ = f.Close()
}
