package serviceServer

import (
	"context"
	"github.com/Diode222/Mimiron/jieba"
	pb "github.com/Diode222/Mimiron/proto_gen"
	"github.com/Diode222/Mimiron/utils"
	"strings"
	"sync"
)

type wordSplitServer struct {}

var serverWordSplit *wordSplitServer
var serverWordSplitOnce sync.Once

func NewWordSplitServer() *wordSplitServer {
	serverWordSplitOnce.Do(func() {
		serverWordSplit = &wordSplitServer{}
	})
	return serverWordSplit;
}

func (s *wordSplitServer) GetWordSplittedMessageList(context context.Context, list *pb.ChatMessageList) (*pb.ChatMessageList, error) {
	splittedChatMessageList := &pb.ChatMessageList{
		ChatMessages:         []*pb.ChatMessage{},
	}
	segmentor := jieba.GetJieba()
	sourceChatMessages := list.GetChatMessages()
	for _, sourceChatMessage := range sourceChatMessages {
		sourceChatMessage.WordAndPosList = []*pb.WordAndPos{}
		wordAndPosList := segmentor.Cut(sourceChatMessage.GetMessage())
		if len(wordAndPosList) <= 0 {
			continue
		}

		for _, wordAndPos := range wordAndPosList {
			word := wordAndPos[0]
			posType := pb.PartOfSpeech_UNKNOWN
			if strings.HasPrefix(wordAndPos[1], "n") {
				posType = pb.PartOfSpeech_NOUN
			} else if strings.HasPrefix(wordAndPos[1], "v") {
				posType = pb.PartOfSpeech_VERB
			} else if strings.HasPrefix(wordAndPos[1], "a") {
				posType = pb.PartOfSpeech_ADJECTIVE
			} else if utils.StringContains([]string{"e", "f", "i", "j", "l"}, wordAndPos[1]) {
				posType = pb.PartOfSpeech_PHRASE
			}

			sourceChatMessage.WordAndPosList = append(sourceChatMessage.WordAndPosList, &pb.WordAndPos{
				Word:                 &word,
				Pos:                  &pb.PartOfSpeech{
					Type:                 &posType,
				},
			})
		}

		// 只返回具有有效分词的消息
		if len(sourceChatMessage.GetWordAndPosList()) > 0 {
			splittedChatMessageList.ChatMessages = append(splittedChatMessageList.ChatMessages, sourceChatMessage)
		}
	}

	return splittedChatMessageList, nil
}
