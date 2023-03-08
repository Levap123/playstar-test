package entities

type LogMessage struct {
	Level   string `json:"level"`
	Service string `json:"service"`
	Error   string `json:"error"`
	Time    string `json:"time"`
	Caller  string `json:"caller"`
	Message string `json:"message"`
}
