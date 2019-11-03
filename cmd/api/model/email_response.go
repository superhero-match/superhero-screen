package model

// CheckEmailResponse holds the response data to the check if email exists request.
type CheckEmailResponse struct {
	IsRegistered bool      `json:"isRegistered"`
	IsDeleted    bool      `json:"isDeleted"`
	IsBlocked    bool      `json:"isBlocked"`
	Superhero    Superhero `json:"superhero"`
}
