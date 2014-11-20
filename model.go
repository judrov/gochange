package changeorg

type Petition struct {
	ID int
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
