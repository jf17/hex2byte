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

	if len(argsWithProg) > 2 {
		fmt.Println("argument count >1 exeptions !")
		os.Exit(4)
	} else if len(argsWithProg) == 0 {
		fmt.Println("argument count = 0 exeptions !")
		os.Exit(4)
	} else if len(argsWithProg) < 2 {
		fmt.Println("argument count < 2 exeptions !")
		os.Exit(4)
	}

	b, err := ioutil.ReadFile(argsWithProg[1]) // just pass the file name
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	str := string(b) // convert content to a 'string'

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
