package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"github.com/tukangkod/go-gobok/tm"
	"github.com/tukangkod/go-gobok/utils"
	"github.com/spf13/viper"
)

type tagmsgServer struct{}

func newTagMsgServer() tm.TagMsgServiceServer {
	return new(tagmsgServer)
}

// Put TagMsg method
func (s *tagmsgServer) Put(ctx context.Context, msg *tm.PutRequest) (*tm.PutResponse, error) {
	utils.Log.Infof(utils.LogTemplate(), "main.Put", "Run")

	marshalMsg := utils.MarshalMsg(msg)
	utils.Log.Infof(utils.LogTemplate(), "main.Put", "Data: " + marshalMsg)

	err := SaveTagMsg(msg)
	if err != nil {
		return &tm.PutResponse{ResponseMsg: false}, err
	}

	return &tm.PutResponse{ResponseMsg: true}, nil
}

// Search TagMsg via SearchRequest
func (s *tagmsgServer) Search(ctx context.Context, msg *tm.SearchRequest) (*tm.SearchResponse, error) {
	tagmsg, err := SearchTagMsg(msg)
	if err != nil {
		utils.Log.Errorf(utils.ErrTemplate(), "Search", err)
	}

	result := make([]*tm.SearchResult, 0)
	for _, data := range tagmsg {
		result = append(result, &tm.SearchResult{ClientIp: data.ClientIp, ServerIp: data.ServerIp, Tags: data.Tags, Message: data.Message})
	}

	return &tm.SearchResponse{SearchResult: result}, nil
}

// run service
func Run() error {
	fnName := utils.GetFunctionName(Run)
	utils.Log.Infof(utils.LogTemplate(), fnName, "RUN")

	address := viper.GetString("grpc.host") + ":" + viper.GetString("grpc.port")
	listen, err := net.Listen(viper.GetString("grpc.network"), address)
	if err != nil {
		utils.Log.Panicf(utils.ErrTemplate(), fnName, err)
	}

	utils.Log.Infof(utils.LogTemplate(), fnName, "Connecting...")
	server := grpc.NewServer()
	utils.Log.Infof(utils.LogTemplate(), fnName, "New Server Success ")

	utils.Log.Infof(utils.LogTemplate(), fnName, "RegisterPutServiceServer")
	tm.RegisterTagMsgServiceServer(server, newTagMsgServer())

	utils.Log.Infof(utils.LogTemplate(), fnName, "Serve")
	if err := server.Serve(listen); err != nil {
		utils.Log.Errorf(utils.ErrTemplate(), utils.GetFunctionName(server.Serve), err)
	}
	return nil
}

// method entry point
func main() {
	utils.NewLog()
	utils.Log.Infof(utils.LogTemplate(), utils.GetFunctionName(main), "START")

	utils.InitConfig()
	DBConnect()
	defer db.Close()
	CreateTable()

	if err := Run(); err != nil {
		utils.Log.Errorf(utils.ErrTemplate(), utils.GetFunctionName(main), err)
	}
}