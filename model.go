package changeorg

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

// Declares AuthKeysResponse for storing auth key request responses.
type AuthKeysResponse struct {
	PetitionID     string   `json:"petition_id"`
	SourceDesc     string   `json:"source_description"`
	Source         string   `json:"source"`
	RequesterEmail string   `json:"requester_email"`
	AuthKey        string   `json:"auth_key"`
	Status         string   `json:"status"`
	Result         string   `json:"result"`
	Messages       []string `json:"messages"`
}

// Declares PetitionIdArgs for constructing petition id requests.
type PetitionIdArgs struct {
	PetitionURL string `json:"petition_url"`
}

// Declares Response to store request reponses.
type Response struct {
	PetitionID int      `json:"petition_id"`
	AuthKey    string   `json:"auth_key"`
	Result     string   `json:"result"`
	Messages   []string `json:"messages"`
}
