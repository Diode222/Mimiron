package serviceServer

import (
	"context"
	"github.com/Diode222/Mimiron/jieba"
	pb "github.com/Diode222/Mimiron/proto_gen"
	"github.com/Diode222/Mimiron/utils"
	"strings"
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
	segmentor := jieba.GetJieba()
	splittedMessageList := &pb.SplittedMessageList{}
	chatMessages := list.GetChatMessages()
	splittedMessages := []*pb.SplittedMessage{}
	for _, chatMessage := range chatMessages {
		wordAndPosList := segmentor.Cut(chatMessage.GetMessage())
		if len(wordAndPosList) <= 0 {
			continue
		}

		splittedMessage := &pb.SplittedMessage{
			WordAndPosList: []*pb.WordAndPos{},
			Time:                 chatMessage.Time,
			ChatPerson:           chatMessage.ChatPerson,
			Message:              chatMessage.Message,
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

			splittedMessage.WordAndPosList = append(splittedMessage.WordAndPosList, &pb.WordAndPos{
				Word:                 &word,
				Pos:                  &pb.PartOfSpeech{
					Type:                 &posType,
				},
			})
		}

		if len(splittedMessage.WordAndPosList) > 0 {
			splittedMessages = append(splittedMessages, splittedMessage)
		}
	}
	splittedMessageList.SplittedMessages = splittedMessages
	return splittedMessageList, nil
}
