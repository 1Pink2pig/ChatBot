package model

type NewChatRequest struct {
	UserId string 
	Balance  float64
}

type EndChatRequest struct {
	UserId
}

type ChatRequest struct {

}