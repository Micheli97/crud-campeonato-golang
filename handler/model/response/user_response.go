package response

type UserResponse struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
