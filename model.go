package changeorg

type Petition struct {
	ID int
}

type PetitionIdArgs struct {
	PetitionURL string `json:"petition_url"`
}

type Response struct {
	PetitionID int      `json:"petition_id"`
	Result     string   `json:"result"`
	Messages   []string `json:"messages"`
}
