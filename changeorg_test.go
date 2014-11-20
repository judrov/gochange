package changeorg

// Imports required packages.
import (
	"fmt"
	"testing"
)

// Creates map for application parameters.
var params map[string]string

// Initializes the parameters.
func init() {
	params = make(map[string]string)
	// sets API key
	params["API_KEY"] = "my_api_key"
	// sets secret token
	params["SECRET"] = "my_secret_token"
	// sets url for the petition
	params["URL"] = "https://www.change.org/p/sample-for-developers-using-change-org-api-sign-petition-via-change-org-api"
	// sets petition id
	params["PETITION_ID"] = "2268806"
	// sets authorization key for petition with id `2268806`
	params["AUTH_KEY"] = "auth_key_for_petition"
}

// TestID tests GetPetitionId function.
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

// TestAuthKey tests GetAuthKey function.
func TestAuthKey(t *testing.T) {
	id := params["PETITION_ID"]
	// checks if petition id is set.
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

// TestSignature tests SignPetition function.
func TestSignature(t *testing.T) {
	id := params["PETITION_ID"]
	// checks if petition id is set.
	if len(id) <= 0 {
		err := "Petition ID is empty. Use GetPetitionId first."
		t.Fatal(err)
	}
	auth_key := params["AUTH_KEY"]
	// checks if auth key is set.
	if len(auth_key) <= 0 {
		err := "Auth Key is empty. Use GetAuthKey first."
		t.Fatal(err)
	}
	msg := "Welcome to gochangeorg!\n"
	msg += "Go sign a petition via API! Pun intended ;)\n"
	change_org := NewChangeOrgClient(params["API_KEY"])
	response, err := change_org.SignPetition(PetitionArgs{
		PetitionID: id,
		AuthKey:    auth_key,
		Source:     "source_that_is_using_the_api",
		Email:      "email@example.com",
		FirstName:  "Grace",
		LastName:   "Hopper",
		Address:    "123 Address",
		City:       "New Haven",
		State:      "CT",
		ZIP:        "06520",
		Country:    "US",
		Hidden:     "true",
		TimeStamp:  GetTimeNow(),
		Endpoint:   "/v1/petitions/" + id + "/signatures",
	}, params["SECRET"])
	if err != nil {
		t.Fatal(err)
	}
	msg += "Response: " + response
	fmt.Println(msg)
}
