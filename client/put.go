package main

import (
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/tukangkod/go-gobok/tm"
	"github.com/tukangkod/go-gobok/utils"
	"fmt"
	"github.com/spf13/viper"
)

// client put
// cli example:
// $ client/put.go 0.0.0.0 "" "a=b&c=d" "message 1"
// $ client/put.go "" 0.0.0.0 "e=f&g=h" "message 2"
func main() {
	utils.NewLog()
	utils.InitConfig()

	// Set up a connection to the server.
	address := viper.GetString("grpc.host") + ":" + viper.GetString("grpc.port")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.Log.Fatalf("not connected: %v", err)
	}
	defer conn.Close()

	c := tm.NewTagMsgServiceClient(conn)

	if len(os.Args) < 5 {
		println("parameter missing")
		os.Exit(1)
	}

	clientIp := os.Args[1]
	serverIp := os.Args[2]
	tags := utils.TagMap(os.Args[3])
	msg := os.Args[4]

	r, err := c.Put(context.Background(), &tm.PutRequest{ClientIp: clientIp, ServerIp: serverIp, Tags: tags, Message: msg})
	if err != nil {
		utils.Log.Fatalf("could not put: %v", err)
	}

	fmt.Println(r.ResponseMsg)
	os.Exit(0)
}
