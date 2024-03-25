package domain

type Equipment struct {
	ID          int64  `json:"id"`
	Identifier  string `json:"identifier"`
	Uniorg      string `json:"uniorg"`
	Code        string `json:"code"`
	UsedCode    bool   `json:"used_code"`
	Environment string `json:"environment"`
	System      string `json:"system"`
	Schedule    string `json:"schedule"`
	Users       []User `json:"users"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
}
