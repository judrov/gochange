// Package gochange provides functionalities in Go to add signatures
// to petitions via the Change.org API.
package gochange

// Imports required packages.
import (
	"errors"
	"net/http"
	"net/url"

	"github.com/judrov/gochange/util"
)

// Sets Host as the base url for the Change.org API.
const (
	Host = "https://api.change.org/v1/"
)

// Declares the Changeorg data type
type ChangeOrg struct {
	Key  string
	Host string
}

// NewChangeOrgClient instantiates a new ChangeOrg.
func NewChangeOrgClient(key string) *ChangeOrg {
	return &ChangeOrg{Key: key, Host: Host}
}

// GetPetitionId gets the petition id for a given petition url.
func (c *ChangeOrg) GetPetitionId(args PetitionIdArgs) (*int, error) {
	// sets up the request parameters.
	v := url.Values{}
	v.Set("api_key", c.Key)
	if len(args.PetitionURL) > 0 {
		v.Set("petition_url", args.PetitionURL)
	}
	// sets up the URL for requesting petition id.
	url := c.Host + "petitions/get_id?" + v.Encode()
	// makes request
	data, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// parses the JSON response
	var res PetitionIdResponse
	if err := util.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	// checks error messages.
	if msg := GetStrings(res.Messages); msg != "" {
		err := errors.New(msg)
		return nil, err
	}
	return &res.PetitionID, err
}

// GetAuthKey grants authorization to gather signatures for a petition
// and returns the authorization code. You will receive the code via email.
func (c *ChangeOrg) GetAuthKey(args AuthKeysArgs, secret string) (string, error) {
	// sets up the petition parameters.
	v := url.Values{}
	v.Set("api_key", c.Key)
	v.Set("petition_id", args.PetitionID)
	v.Set("source_description", args.SourceDesc)
	v.Set("source", args.Source)
	v.Set("requester_email", args.RequesterEmail)
	v.Set("timestamp", util.GetTimeNow())
	v.Set("callback", args.Callback)
	v.Set("endpoint", "/v1/petitions/"+args.PetitionID+"/auth_keys")
	v.Set("rsig", util.Hash(v.Encode()+secret))
	// sets up the URL for requesting authorization key for a petition.
	url := c.Host + "petitions/" + args.PetitionID + "/auth_keys"
	// makes the request.
	var res AuthKeysResponse
	if err := util.Post(url, v.Encode(), &res); err != nil {
		return res.AuthKey, err
	}
	// checks error messages.
	if msg := GetStrings(res.Messages); msg != "" {
		err := errors.New(msg)
		return res.AuthKey, err
	}
	return res.AuthKey, nil
}

// SignPetition adds a signature to a petition.
func (c *ChangeOrg) SignPetition(args PetitionArgs, secret string) (string, error) {
	// sets up the signature parameters.
	v := url.Values{}
	v.Set("api_key", c.Key)
	v.Set("source", args.Source)
	v.Set("email", args.Email)
	v.Set("first_name", args.FirstName)
	v.Set("last_name", args.LastName)
	v.Set("address", args.Address)
	v.Set("city", args.City)
	v.Set("state_province", args.State)
	v.Set("postal_code", args.ZIP)
	v.Set("country_code", args.Country)
	v.Set("hidden", args.Hidden)
	v.Set("timestamp", util.GetTimeNow())
	v.Set("endpoint", "/v1/petitions/"+args.PetitionID+"/signatures")
	v.Set("rsig", util.Hash(v.Encode()+secret+args.AuthKey))
	// sets up the URL for adding petition signatures.
	url := c.Host + "petitions/" + args.PetitionID + "/signatures"
	// POST the parameters to the signature endpoint.
	var res PetitionResponse
	if err := util.Post(url, v.Encode(), &res); err != nil {
		return res.Result, err
	}
	// checks error messages.
	if msg := GetStrings(res.Messages); msg != "" {
		err := errors.New(msg)
		return res.Result, err
	}
	return res.Result, nil
}

// GetStrings returns the concatenation of strings from an array.
func GetStrings(messages []string) string {
	msg := ""
	if len(messages) > 0 {
		for i := range messages {
			msg += messages[i]
		}
	}
	return msg
}
