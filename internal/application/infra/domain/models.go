package domain

type HelpMessage struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Handler struct {
	Handler MessageHandler
}
