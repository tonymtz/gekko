package user

// User ...
type User struct {
	ID             int64  `json:"id"`
	IDProvider     string `json:"id_provider"`
	DisplayName    string `json:"display_name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	Role           int64  `json:"role,omitempty"`
	Token          string `json:"token,omitempty"`
	JWT            string `json:"jwt,omitempty"`
}
