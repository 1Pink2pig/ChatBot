package model

type NewChatRequest struct {
	UserId  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

type EndChatRequest struct {
	UserId string `json:"user_id"`
	ChatId string `json:"chat_id"`
}

type ChatRequest struct {
	UserId   string `json:"user_id"`
	ChatId   string `json:"chat_id"`
	Question string `json:"question"`
}

type ChatReply struct {
	Response string `json:"response"`
}