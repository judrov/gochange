// Package model contains structs to store request and response parameters
// for interacting with the Change.org API.
package model

// Declares PetitionIdArgs for constructing petition id requests.
type PetitionIdArgs struct {
	PetitionURL string `json:"petition_url"`
}

// PetitionIdResponse stores petition id request responses.
type PetitionIdResponse struct {
	PetitionID int `json:"petition_id"`
	Response
}

// Declares AuthKeysArgs for constructing auth key retrival requests.
type AuthKeysArgs struct {
	PetitionID        string `json:"petition_id"`
	SourceDesc        string `json:"source_description"`
	Source            string `json:"source"`
	RequesterEmail    string `json:"requester_email"`
	TimeStamp         string `json:"timestamp"`
	Callback          string `json:"callback"`
	Callback_Endpoint string `json:"callback_endpoint"`
	Endpoint          string `json:"endpoint"`
	RSIG              string `json:"rsig"`
}

// AuthKeysResponse stores auth key request responses.
type AuthKeysResponse struct {
	PetitionID     int    `json:"petition_id"`
	SourceDesc     string `json:"source_description"`
	Source         string `json:"source"`
	RequesterEmail string `json:"requester_email"`
	AuthKey        string `json:"auth_key"`
	Status         string `json:"status"`
	Response
}

// Declares PetitionArgs for constructing signature requests.
type PetitionArgs struct {
	PetitionID string `json:"petition_id"`
	AuthKey    string `json:"auth_key"`
	Source     string `json:"source"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state_province"`
	ZIP        string `json:"postal_code"`
	Country    string `json:"country_code"`
	Hidden     string `json:"hidden"`
	RSIG       string `json:"rsig"`
	TimeStamp  string `json:"timestamp"`
	Endpoint   string `json:"endpoint"`
}

// PetitionResponse stores response to adding signatures to petitions.
type PetitionResponse struct {
	Response
}

// Response stores http request reponses.
type Response struct {
	Result   string   `json:"result"`
	Messages []string `json:"messages"`
}
