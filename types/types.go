package types

type APIResponse struct {
	Status string `json:"status"`
	Result any    `json:"result"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type RootResult struct {
	Message  string `json:"message"`
	Version  string `json:"version"`
	DateTime string `json:"datetime"`
}

type MorseRequest struct {
	Input string `json:"input"`
}

type ConversionResult struct {
	Input  string `json:"input"`
	Output string `json:"output"`
	Length int    `json:"length"`
}
