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

// client search - cli
// client/search.go <ClientIp> <ServerIp> <tag>
// client/search.go // search everything
// client/search.go "" "" "a=b"
// client/search.go "" "" "a=b&c=d"
// client/search.go "" "" "a=b&c=d"
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

	var clientIp, serverIp string
	var tags map[string]string

	if len(os.Args) == 4 {
		if os.Args[1] != "" {
			clientIp = os.Args[1]
		}
		if os.Args[2] != "" {
			serverIp = os.Args[2]
		}
		if os.Args[3] != "" {
			tags = utils.TagMap(os.Args[3])
		}
	}

	r, err := c.Search(context.Background(), &tm.SearchRequest{ClientIp: clientIp, ServerIp: serverIp, Tags: tags})
	if err != nil {
		utils.Log.Fatalf("could not put: %v", err)
	}

	fmt.Printf(utils.MarshalMsg(r.SearchResult))
	os.Exit(0)
}
