package datatransfers

type ApiResponse struct {
	Code          int         `json:"code"`
	Status        string      `json:"status"`
	Message       string      `json:"message"`
	MessageDetail string      `json:"message_detail"`
	Data          interface{} `json:"data"`
}
