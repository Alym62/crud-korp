package dto

type ResponseAuth struct {
	CurrentUser CurrentUserResponse `json:"currentUser"`
	Token       string              `json:"token"`
}