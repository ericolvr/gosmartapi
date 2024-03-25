package domain

type Equipment struct {
	ID          int64  `json:"id"`
	Identifier  string `json:"identifier"`
	Uniorg      string `json:"uniorg"`
	Code        string `json:"code"`
	UsedCode    bool   `json:"used_code"`
	Environment string `json:"environment"`
	SystemName  string `json:"system_name"`
	Schedule    string `json:"schedule"`
	Users       string `json:"user"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
}
