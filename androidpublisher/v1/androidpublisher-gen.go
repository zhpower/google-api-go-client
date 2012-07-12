// Package androidpublisher provides access to the Google Play Android Developer API.
//
// See https://developers.google.com/android-publisher
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/androidpublisher/v1"
//   ...
//   androidpublisherService, err := androidpublisher.New(oauthHttpClient)
package androidpublisher

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New

const apiId = "androidpublisher:v1"
const apiName = "androidpublisher"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/androidpublisher/v1/applications/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Purchases = &PurchasesService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Purchases *PurchasesService
}

type PurchasesService struct {
	s *Service
}

type SubscriptionPurchase struct {
	// AutoRenewing: Whether the subscription will automatically be renewed
	// when it reaches its current expiry time.
	AutoRenewing bool `json:"autoRenewing,omitempty"`

	// InitiationTimestampMsec: Time at which the subscription was granted,
	// in milliseconds since Epoch.
	InitiationTimestampMsec int64 `json:"initiationTimestampMsec,omitempty,string"`

	// Kind: This kind represents a subscriptionPurchase object in the
	// androidpublisher service.
	Kind string `json:"kind,omitempty"`

	// ValidUntilTimestampMsec: Time at which the subscription will expire,
	// in milliseconds since Epoch.
	ValidUntilTimestampMsec int64 `json:"validUntilTimestampMsec,omitempty,string"`
}

// method id "androidpublisher.purchases.cancel":

type PurchasesCancelCall struct {
	s              *Service
	packageName    string
	subscriptionId string
	token          string
	opt_           map[string]interface{}
}

// Cancel: Cancels a user's subscription purchase. The subscription
// remains valid until its expiration time.
func (r *PurchasesService) Cancel(packageName string, subscriptionId string, token string) *PurchasesCancelCall {
	c := &PurchasesCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.packageName = packageName
	c.subscriptionId = subscriptionId
	c.token = token
	return c
}

func (c *PurchasesCancelCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/androidpublisher/v1/applications/", "{packageName}/subscriptions/{subscriptionId}/purchases/{token}/cancel")
	urls = strings.Replace(urls, "{packageName}", cleanPathString(c.packageName), 1)
	urls = strings.Replace(urls, "{subscriptionId}", cleanPathString(c.subscriptionId), 1)
	urls = strings.Replace(urls, "{token}", cleanPathString(c.token), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Cancels a user's subscription purchase. The subscription remains valid until its expiration time.",
	//   "httpMethod": "POST",
	//   "id": "androidpublisher.purchases.cancel",
	//   "parameterOrder": [
	//     "packageName",
	//     "subscriptionId",
	//     "token"
	//   ],
	//   "parameters": {
	//     "packageName": {
	//       "description": "The package name of the application for which this subscription was purchased (for example, 'com.some.thing').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "The purchased subscription ID (for example, 'monthly001').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "token": {
	//       "description": "The token provided to the user's device when the subscription was purchased.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{packageName}/subscriptions/{subscriptionId}/purchases/{token}/cancel"
	// }

}

// method id "androidpublisher.purchases.get":

type PurchasesGetCall struct {
	s              *Service
	packageName    string
	subscriptionId string
	token          string
	opt_           map[string]interface{}
}

// Get: Checks whether a user's subscription purchase is valid and
// returns its expiry time.
func (r *PurchasesService) Get(packageName string, subscriptionId string, token string) *PurchasesGetCall {
	c := &PurchasesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.packageName = packageName
	c.subscriptionId = subscriptionId
	c.token = token
	return c
}

func (c *PurchasesGetCall) Do() (*SubscriptionPurchase, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/androidpublisher/v1/applications/", "{packageName}/subscriptions/{subscriptionId}/purchases/{token}")
	urls = strings.Replace(urls, "{packageName}", cleanPathString(c.packageName), 1)
	urls = strings.Replace(urls, "{subscriptionId}", cleanPathString(c.subscriptionId), 1)
	urls = strings.Replace(urls, "{token}", cleanPathString(c.token), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(SubscriptionPurchase)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Checks whether a user's subscription purchase is valid and returns its expiry time.",
	//   "httpMethod": "GET",
	//   "id": "androidpublisher.purchases.get",
	//   "parameterOrder": [
	//     "packageName",
	//     "subscriptionId",
	//     "token"
	//   ],
	//   "parameters": {
	//     "packageName": {
	//       "description": "The package name of the application for which this subscription was purchased (for example, 'com.some.thing').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "The purchased subscription ID (for example, 'monthly001').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "token": {
	//       "description": "The token provided to the user's device when the subscription was purchased.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{packageName}/subscriptions/{subscriptionId}/purchases/{token}",
	//   "response": {
	//     "$ref": "SubscriptionPurchase"
	//   }
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
