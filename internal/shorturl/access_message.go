package shorturl

// AccessMessage is message sent to indicate short url access from IP address
type AccessMessage struct {
	ShortId string `json:"shortId"`
	Ip      string `json:"ip"`
}
