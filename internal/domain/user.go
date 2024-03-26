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

type UserResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Document  string `json:"document"`
	Role      string `json:"role"`
	Photo     string `json:"photo"`
	Completed bool   `json:"completed"`
}

func (u *User) Response() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Document:  u.Document,
		Role:      u.Role,
		Photo:     u.Photo,
		Completed: u.Completed,
	}
}
