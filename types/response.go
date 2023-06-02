package types

type Response struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
