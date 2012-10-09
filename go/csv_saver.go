package main

import (
	"bytes"
	"encoding/csv"
	"os"
	"strconv"
)

type Data struct {
	saveInt    int
	saveString string
}

func dToStrArr(d Data) []string {
	s := make([]string, 2)
	s[0] = strconv.Itoa(d.saveInt)
	s[1] = d.saveString
	return s
}

func saveCsv(d []Data, filePath string) {
	buf := new(bytes.Buffer)
	csvWtr := csv.NewWriter(buf)
	for i := 0; i < len(d); i++ {
		csvWtr.Write(dToStrArr(d[i]))
		csvWtr.Flush()
	}
	file, _ := os.Create(filePath)
	defer file.Close()
	file.WriteString(buf.String())
}

func main() {
	d := [2]Data{}
	d[1].saveInt = 1
	d[1].saveString = "b"
	d[0].saveInt = 0
	d[0].saveString = "a"
	saveCsv(d[0:2], "d.csv")
}
