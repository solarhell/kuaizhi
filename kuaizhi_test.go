package kuaizhi

import (
	"testing"
)

const tokenStr = "749379809e98cde019a55e52ba5ab060"

func TestBuildLink(t *testing.T) {
	linkStr := BuildLink("https://bing.com", "bing")
	if linkStr != `<a-link href="https://bing.com">bing</a-link>` {
		t.Error("BuildLink error")
	}
}

func TestGetJob(t *testing.T) {
	job, err := GetJob(tokenStr)
	if err != nil {
		t.Error(err)
	}

	t.Log(job.Job.JobID)
}

func TestPostMessage(t *testing.T) {
	card := Card{
		Title: "测试",
		Text:  "测试文本",
	}
	var cards []Card
	err := PostMessage(tokenStr, Message{
		Cards: append(cards, card),
	})

	if err != nil {
		t.Error(err)
	}
}
