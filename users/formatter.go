package users

type UserFormatResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func FormatUserResponse(user User, token string) UserFormatResponse {
	formatter := UserFormatResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return formatter
}
