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
	params["PETITION_ID"] = "2268806"
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

func TestAuthKey(t *testing.T) {
	id := params["PETITION_ID"]
	if len(id) <= 0 {
		err := "Petition ID is empty. Use GetPetitionId first."
		t.Fatal(err)
	}
	msg := "Requesting Auth Key!\n"
	change_org := NewChangeOrgClient(params["API_KEY"])
	authKey, err := change_org.GetAuthKey(AuthKeysArgs{
		PetitionID:     id,
		SourceDesc:     "source_description",
		Source:         "source_that_is_using_the_api",
		RequesterEmail: "developer_email",
		TimeStamp:      GetTimeNow(),
		Endpoint:       "/v1/petitions/" + id + "/auth_keys",
		Callback:       "mycallback",
	}, params["SECRET"])
	if err != nil {
		t.Fatal(err)
	}
	msg += "Auth Key: " + authKey
	fmt.Println(msg)
}
