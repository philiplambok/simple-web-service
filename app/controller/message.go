package controller

type Message struct {
	Body string `json:"body"`
}

func (m *Message) Set(body string) {
	m.Body = body
	return
}
