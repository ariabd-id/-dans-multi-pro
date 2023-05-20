package params

type CreateUser struct {
	ID       int    `json:"int"`
	Password string `json:"password,omitempty"`
	Username string `json:"username"`
}
