package handler

import (
	"strings"
	"net"
	"bufio"
	"os"
	"fmt"
	"NonRelDB/log"
	"NonRelDB/util/regex"
)

// SendRequest sends request to specified connection.
func SendRequest(req string, c net.Conn){
	fmt.Fprintf(c, req + "\n")
}

// HandleConnection handling communication with server.
func HandleConnection(c net.Conn){
	consoleReader := bufio.NewReader(os.Stdin)
	netReader := bufio.NewReader(c)
	for {
		fmt.Print("NonRelDB> ")
		req, err := consoleReader.ReadString('\n')
		req = strings.Trim(req, "\n")

		if err != nil {
			log.Error.Panicln(err.Error())
		}

		if regex.QueryReg.MatchString(req){
			SendRequest(req, c)
			resp, err := netReader.ReadString('\n')

			if err != nil {
				log.Error.Panicln(err.Error())
			}

			fmt.Println(resp)

		} else if regex.ExitReg.MatchString(req){
			fmt.Println("Good bye")
			SendRequest(req, c)
			return

		} else if regex.TopicReg.MatchString(req) {
			reqParts := strings.Split(req, " ")

			switch reqCtx := len(reqParts); reqCtx{
				case 2:{
					if strings.ToLower(reqParts[0]) == "subscribe"{
							SendRequest(req, c)
							HandleTopic(c, *netReader, reqParts[1])
					}
				}
				case 3:{
					if strings.ToLower(reqParts[0]) == "publish"{
						SendRequest(req, c)
					}
				}
			}
		} else {
			fmt.Println("Bad request")
			continue 
		}
	}
}