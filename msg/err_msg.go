package msg

type ErrMsg struct {
	Info string `json:"info"`
}

type CodeMsg struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}
