package jieba

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
	"sync"
	"unicode/utf8"
)

type jieba struct {
	segmenter *gojieba.Jieba
}

var j *jieba
var jiebaOnce sync.Once

func GetJieba() *jieba {
	jiebaOnce.Do(func() {
		j = &jieba{
			segmenter: gojieba.NewJieba(),
		}
	})
	return j
}

// 句子分为多个词语切片，切片里有两个元素，分别是词语和词性。目前逻辑是抛弃掉单个字
func (j *jieba) Cut(str string) [][]string {
	wordAndPosList := [][]string{}
	wordInfoList := j.segmenter.Tag(str)
	for _, wordInfo := range wordInfoList {
		wordAndPos := strings.Split(wordInfo, "/")
		if len(wordAndPos) <= 1 || utf8.RuneCountInString(wordAndPos[0]) <= 1 {
			fmt.Println("filtered word: ", wordAndPos[0])
			continue
		}
		fmt.Println("remains word: ", wordAndPos[0])

		wordAndPosList = append(wordAndPosList, wordAndPos)
	}
	return wordAndPosList
}
