package requests

type UserStoreRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

type UserAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
