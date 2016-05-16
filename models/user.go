package models

// USER ROLES:
// 0 - admin
// 1 - registered

type User struct {
	Id             int64   `json:"id"`
	IdProvider     string  `json:"id_provider"`
	DisplayName    string  `json:"display_name"`
	Email          string  `json:"email"`
	ProfilePicture string  `json:"profile_picture"`
	Role           int64   `json:"role,omitempty"`
	Token          string  `json:"token,omitempty"`
	JWT            string  `json:"jwt,omitempty"`
}
