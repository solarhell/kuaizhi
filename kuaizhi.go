package kuaizhi

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

func BuildLink(href, name string) string {
	return fmt.Sprintf(`<a-link href="%s">%s</a-link>`, href, name)
}

func GetJob(token string) (Job, error) {
	r, err := req.Get(fmt.Sprintf("https://kuaizhi.app/bot/%s/getJob", token))
	if err != nil {
		panic(err)
	}
	var job Job
	err = r.ToJSON(&job)
	return job, err
}

func PostMessage(token string, msg Message) error {
	req.SetJSONEscapeHTML(false)
	r, err := req.Post(fmt.Sprintf("https://kuaizhi.app/bot/%s/pushMessage", token), req.BodyJSON(msg))
	if err != nil {
		return err
	}
	var res CommonResponse
	err = r.ToJSON(&res)
	if err != nil {
		return err
	}
	if res.Errno != 0 {
		return errors.New(res.ErrMsg)
	}
	return nil
}
