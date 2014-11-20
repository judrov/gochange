package changeorg

type Petition struct {
	ID int
}

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

type PetitionIdArgs struct {
	PetitionURL string `json:"petition_url"`
}

type Response struct {
	PetitionID int      `json:"petition_id"`
	AuthKey    string   `json:"auth_key"`
	Result     string   `json:"result"`
	Messages   []string `json:"messages"`
}
