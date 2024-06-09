package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Server will listen at 0.0.0.0:8888, input http://localhost:8888 to browser address")

		mux := http.NewServeMux()
		files := http.FileServer(http.Dir("."))
		mux.Handle("/", files)

		server := &http.Server{
			Addr:    "0.0.0.0:8888",
			Handler: mux,
		}
		server.ListenAndServe()
	} else if len(os.Args) != 3 || os.Args[1] != "-f" {
		fmt.Println("Usage: ", os.Args[0], " [-f <raw.txt>]")
		fmt.Println("Use option -f and parmameter (file name), format text as standard format")
		fmt.Println("If no option or parameter provided, a server will be started")
	} else {
		var readBuf []byte
		var err error
		if readBuf, err = os.ReadFile(os.Args[2]); err != nil {
			log.Fatal("ReadFile error, ", os.Args[2], " is not a file?")
		}
		text := strings.Trim(string(readBuf), " \t")
		re := regexp.MustCompile(`(\s*\n)+`) // 可能一些空白再换行
		text = re.ReplaceAllString(text, "\n")
		re = regexp.MustCompile(`\n\s+`) // 行首空白
		text = re.ReplaceAllString(text, "\n")
		re = regexp.MustCompile(`(\w+\)?)\n`) // 单词或单词+右括号再换行
		text = re.ReplaceAllString(text, "$1 -> ")
		fmt.Print(text)
	}
}
