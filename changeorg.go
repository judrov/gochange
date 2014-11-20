package changeorg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Host = "https://api.change.org/v1/"
)

type ChangeOrg struct {
	Key  string
	Host string
}

func NewChangeOrgClient(key string) *ChangeOrg {
	return &ChangeOrg{Key: key, Host: Host}
}

func (c *ChangeOrg) GetPetitionId(args PetitionIdArgs) (*int, error) {
	v := url.Values{}
	v.Set("api_key", c.Key)
	if len(args.PetitionURL) > 0 {
		v.Set("petition_url", args.PetitionURL)
	}
	url := c.Host + "petitions/get_id?" + v.Encode()
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var bodyRes Response
	if err := unmarshal(res, &bodyRes); err != nil {
		return nil, err
	}
	return &bodyRes.PetitionID, err
}

// GetAuthKey grants authorization to gather signatures for a petition
// and returns the authorization code. You will receive the code via email.
func (c *ChangeOrg) GetAuthKey(args AuthKeysArgs, secret string) (string, error) {
	var res Response
	v := url.Values{}
	v.Set("api_key", c.Key)
	v.Set("petition_id", args.PetitionID)
	v.Set("source_description", args.SourceDesc)
	v.Set("source", args.Source)
	v.Set("requester_email", args.RequesterEmail)
	v.Set("timestamp", args.TimeStamp)
	v.Set("callback", args.Callback)
	v.Set("endpoint", args.Endpoint)
	v.Set("rsig", Hash(v.Encode()+secret))
	url := c.Host + "petitions/" + args.PetitionID + "/auth_keys"
	err := Post(url, v.Encode(), &res)
	return res.AuthKey, err
}

// SignPetition adds a signature to a petition.
func (c *ChangeOrg) SignPetition(args PetitionArgs, secret string) (string, error) {
	var res Response
	v := url.Values{}
	v.Set("api_key", c.Key)
	v.Set("timestamp", args.TimeStamp)
	v.Set("endpoint", args.Endpoint)
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
	v.Set("rsig", Hash(v.Encode()+secret+args.AuthKey))
	url := c.Host + "petitions/" + args.PetitionID + "/signatures"
	err := Post(url, v.Encode(), &res)
	return res.Result, err
}

func unmarshal(res *http.Response, bodyRes *Response) error {
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &bodyRes); err != nil {
		return err
	}
	if len(bodyRes.Messages) > 0 {
		msg := ""
		for i := range bodyRes.Messages {
			msg += bodyRes.Messages[i]
		}
		return errors.New(msg)
	}
	return nil
}
