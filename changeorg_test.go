package changeorg

import (
	"fmt"
	"testing"
)

var params map[string]string

func init() {
	params = make(map[string]string)
	params["API_KEY"] = "my_api_key"
	params["SECRET"] = "my_secret_token"
	params["URL"] = "https://www.change.org/p/sample-for-developers-using-change-org-api-sign-petition-via-change-org-api"
}

func TestID(t *testing.T) {
	msg := "Welcome to gochangeorg!"
	change_org := NewChangeOrgClient(params["API_KEY"])
	id, err := change_org.GetPetitionId(PetitionIdArgs{
		PetitionURL: params["URL"],
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
	fmt.Println("Petition Id:", *id)
}
