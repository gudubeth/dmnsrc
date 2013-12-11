package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

//source: 
//https://groups.google.com/forum/#!msg/golang-nuts/Wg_5BuaJruk/tqBXZLlCOjAJ
func whois(dom, server string) string {
	conn, err := net.Dial("tcp", server+":43")
	if err != nil {
		fmt.Println("Error")
	}
	conn.Write([]byte(dom + "\r\n"))
	buf := make([]byte, 1024)
	res := []byte{}
	for {
		numbytes, err := conn.Read(buf)
		sbuf := buf[0:numbytes]
		res = append(res, sbuf...)
		if err != nil {
			break
		}
	}
	conn.Close()
	return string(res)
}

func main() {
	suffix := flag.String("suffix", "", "suffix for domain")
	prefix := flag.String("prefix", "", "prefix for domain")
	test := flag.Bool("test", false, "if true, only prints domain names. no check is done")

	flag.Parse()

	filename := "words.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("err in reading file")
		return
	}
	words := string(content)
	words = strings.Replace(words, " ", "", -1)
	wordarr := strings.Split(words, ",")

	for i := range wordarr {
		domain := *prefix + wordarr[i] + *suffix + ".com"

		if *test == true {
			fmt.Println(domain)
		} else {
			resp := whois(domain, "com.whois-servers.net")
			if strings.Contains(resp, "No match") {
				fmt.Println(domain)
			}
		}
	}
}
