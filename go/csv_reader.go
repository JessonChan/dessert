package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("1.csv")
	r := csv.NewReader(strings.NewReader(string(b)))
	s, _ := r.ReadAll()
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i][0])
	}
}
