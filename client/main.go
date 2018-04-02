package main

import (
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/tukangkod/go-gobok/tm"
	"strings"
	"github.com/tukangkod/go-gobok/utils"
)

const (
	address    = ":50051"
)

func tagMap(s string) (map[string]string) {
	var m map[string]string
	var ss []string

	m = make(map[string]string)
	ss = strings.Split(s, "&")
	for _, pair := range ss {
		z := strings.Split(pair, "=")
		m[z[0]] = z[1]
	}
	return m
}

func main() {
	utils.NewLog()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.Log.Fatalf("not connected: %v", err)
	}
	defer conn.Close()
	c := tm.NewTagMsgServiceClient(conn)

	// Contact the server and print out its response.
	var clientIp, serverIp, msg string
	var tags map[string]string

	if len(os.Args) > 1 {
		clientIp = os.Args[1]
		serverIp = os.Args[2]
		tags = tagMap(os.Args[3])
		msg = os.Args[4]
	}
	r, err := c.Put(context.Background(), &tm.PutRequest{ClientIp: clientIp, ServerIp: serverIp, Tags: tags, Message: msg})
	if err != nil {
		utils.Log.Fatalf("could not put: %v", err)
	}

	utils.Log.Infof("%s", r.ResponseMsg)
}
