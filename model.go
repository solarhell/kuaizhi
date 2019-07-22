package kuaizhi

type CommonResponse struct {
	Errno  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
}

type Job struct {
	CommonResponse
	Job struct {
		JobID  string        `json:"job_id"`
		Params []interface{} `json:"params"`
	} `json:"job"`
	Cache struct {
	} `json:"cache"`
}

type Card struct {
	Images []string `json:"images"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Url    string   `json:"url"`
	Video  string   `json:"video"`
}

type Message struct {
	JobID string `json:"job_id"`
	Cards []Card `json:"cards"`
}
