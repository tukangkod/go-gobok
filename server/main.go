package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"encoding/json"
	"github.com/tukangkod/go-gobok/tm"
	"github.com/tukangkod/go-gobok/utils"
)

type putServer struct{}

func newTagMsgServer() tm.TagMsgServiceServer {
	return new(putServer)
}

func (s *putServer) Put(ctx context.Context, msg *tm.PutRequest) (*tm.PutResponse, error) {
	utils.Log.Infof(utils.LogTemplate(), "main.Put", "Run")

	marshalMsg, err := json.Marshal(msg)
	if (err != nil) {
		utils.Log.Panicf(utils.LogTemplate(), utils.GetFunctionName(putServer{}), err)
	}
	utils.Log.Infof(utils.LogTemplate(), "main.Put", "Data: " + string(marshalMsg))

	return &tm.PutResponse{ResponseMsg: string(marshalMsg)}, nil
}

func Run() error {
	fnName := utils.GetFunctionName(Run)
	utils.Log.Infof(utils.LogTemplate(), fnName, "RUN")

	listen, err := net.Listen("tcp", ":50051")
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

func main() {
	utils.NewLog()
	utils.Log.Infof(utils.LogTemplate(), utils.GetFunctionName(main), "START")

	if err := Run(); err != nil {
		utils.Log.Errorf(utils.ErrTemplate(), utils.GetFunctionName(main), err)
	}
}