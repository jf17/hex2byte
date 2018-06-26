package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	argsWithProg := os.Args

	if len(argsWithProg) == 2 {

	} else {
		fmt.Println("argument count exeptions !")
		os.Exit(4)
	}

	b, err := ioutil.ReadFile(argsWithProg[1]) // just pass the file name
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	rawStr := string(b) // convert content to a 'string'

	str := strings.ToUpper(rawStr)

	re := regexp.MustCompile(`\r?\n`)
	str2 := re.ReplaceAllString(str, "")

	arrayStr := strings.Split(str2, " ")

	strTobyte := ""

	for _, c := range arrayStr {
		length := len(c)

		if length == 2 {
			strTobyte = strTobyte + c
		}

	}

	decoded, err := hex.DecodeString(strTobyte)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	l := strings.TrimSuffix(argsWithProg[1], path.Ext(argsWithProg[1]))

	f, err := os.Create(l + ".bin")
	check(err)

	defer f.Close()

	n2, err := f.Write(decoded)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

}
