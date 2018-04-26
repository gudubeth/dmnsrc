package checker

import (
	"math/rand"
	"net"
)

// DefaultWhoisServers is list of predefined whois servers.
var DefaultWhoisServers = []string{
	"com.whois-servers.net",
}

// Whois gets whois information of domainName using whoisServer.
//
// It returns whois data and any error encountered. You can use
// SelectRandomWhoisServer() function for random whois server
// from default default list.
//
// Source: https://groups.google.com/forum/#!msg/golang-nuts/Wg_5BuaJruk/tqBXZLlCOjAJ
func Whois(domainName, whoisServer string) (string, error) {
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	conn.Write([]byte(domainName + "\r\n"))
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

	return string(res), nil
}

// SelectRandomWhoisServer selects random whois server from default whois
// servers list.
func SelectRandomWhoisServer() string {
	return DefaultWhoisServers[rand.Intn(len(DefaultWhoisServers))]
}
