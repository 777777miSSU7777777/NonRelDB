package handler

import (
	"net"
	"NonRelDB/server/log"
)

func HandleConnection(c net.Conn){
	defer c.Close()
	for {
		req := make([]byte,1024)
		lenReq, err := c.Read(req)

		if err != nil {
			log.Error.Println(err.Error())
			return
		}

		query := string(req[:lenReq]) 
		log.Info.Printf("Received request from %s -> %s", c.RemoteAddr().String(), query)

		resp_str := HandleQuery(query)
		resp := []byte(resp_str)
		log.Info.Printf("Sent response to %s -> %s", c.RemoteAddr().String(), resp_str)
		c.Write(resp)
	}
}