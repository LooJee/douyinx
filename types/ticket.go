package types

type JSBTicket struct {
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	ExpiresIn   int    `json:"expires_in"`
	Ticket      string `json:"ticket"`
}

type JSBTicketResp struct {
	Data  JSBTicket `json:"data"`
	Extra ExtraData `json:"extra"`
}
