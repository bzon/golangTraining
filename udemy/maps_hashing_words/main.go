package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	testUrl := "https://gist.githubusercontent.com/bzon/e47dde6907511603b5eaf8eeec43a714/raw/01844e3a055dae03cb09bf0c8336ed4d8a44595b/DATA_TEST"
	res, err := http.Get(testUrl)
	if err != nil {
		log.Fatalln(err)
	}
	bs, _ := ioutil.ReadAll(res.Body) // get []bytes
	fmt.Println(bs)
	str := string(bs) // a string is a slice of bytes!
	fmt.Println(str)

	// fmt.Println("Printing this")

	// words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)
	// for sc.Scan() {
	// 	words[sc.Text()] = ""
	// }
	// if err := sc.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading input:", err)
	// }
	// fmt.Println(words)
	// for key := range words {
	// 	fmt.Println(key)
	// }

}
