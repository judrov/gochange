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
