# gochange

[Go] (http://golang.org/doc/install) package for interacting with the [Change.org API] (https://github.com/change/api_docs). [Change.org] (http://www.change.org) is a petition platform, empowering people everywhere to create the change they want to see.

To get started, you will need to request an API key and secret token at [change.org/developers] (https://www.change.org/developers). Also, install [Golang] (http://golang.org/doc/install).

Installation
------------
```go
go get github.com/judrov/gochange
```

Usage
------------
#####Import gochange package
```go
import "github.com/judrov/gochange"
```

#####Set configuration parameters
```go
// sets API key
params["API_KEY"] = "my_api_key"
// sets secret token
params["SECRET"] = "my_secret_token"
// sets url for the petition
params["URL"] = "https://www.change.org/p/sample-for-developers-using-change-org-api-sign-petition-via-change-org-api"
// sets petition id. Use GetPetitionId to retrive petition id. 
params["PETITION_ID"] = "2268806"
// sets authorization key for petition with id `2268806`. Use GetAuthKey to request authorization key.
params["AUTH_KEY"] = "auth_key_for_petition"
```

#####Initialize new gochange client
```go
change_org := gochange.NewChangeOrgClient(params["API_KEY"])
```

#####Retrieve petition id. [API details] (https://github.com/change/api_docs/blob/master/v1/documentation/resources/petitions.md#get-petitions-get_id).
```go
id, err := change_org.GetPetitionId(gochange.PetitionIdArgs{
  PetitionURL: params["URL"],
})
```

#####Request petition authorization key. [API details] (https://github.com/change/api_docs/blob/master/v1/documentation/resources/petitions/auth_keys.md).
```go
authKey, err := change_org.GetAuthKey(gochange.AuthKeysArgs{
  PetitionID:     params["PETITION_ID"],
  SourceDesc:     "source_description",
  Source:         "source_that_is_using_the_api",
  RequesterEmail: "developer_email",
  Callback:       "mycallback",
}, params["SECRET"])
```	

#####Add signature to a petition. [API details] (https://github.com/change/api_docs/blob/master/v1/documentation/resources/petitions/signatures.md).
```go
response, err := change_org.SignPetition(gochange.PetitionArgs{
  PetitionID: params["PETITION_ID"],
  AuthKey:    params["AUTH_KEY"],
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
}, params["SECRET"])
```


Sample petition for developers
------------
This is a sample petition for developers wanting to test adding signatures to petitions via @Change's [API] (https://github.com/change/api_docs): 

https://www.change.org/p/sample-for-developers-using-change-org-api-sign-petition-via-change-org-api

Enjoy signing this petition via the Change.org API! 

