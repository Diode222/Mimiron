package serviceServer

import (
	"context"
	pb "github.com/Diode222/Mimiron/proto_gen"
	"sync"
)

type splittedWordServer struct {}

var serversplittedWord *splittedWordServer
var serversplittedWordOnce sync.Once

func NewSplittedWordServer() *splittedWordServer {
	serversplittedWordOnce.Do(func() {
		serversplittedWord = &splittedWordServer{}
	})
	return serversplittedWord;
}

func (s *splittedWordServer) GetSplittedMessage(context context.Context, list *pb.ChatMessageList) (*pb.SplittedMessageList, error) {
	// TODO 接入分词算法
	return nil, nil
}
