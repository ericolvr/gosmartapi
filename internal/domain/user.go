package domain

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Document  string `json:"document"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	Completed bool   `json:"completed"`
}
