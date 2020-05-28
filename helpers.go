package harvest

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

func addParamsToURL(baseURL string, opt interface{}) (string, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return baseURL, err
	}

	vs, err := query.Values(opt)
	if err != nil {
		return baseURL, err
	}

	url.RawQuery = vs.Encode()
	return url.String(), nil
}

func (h *Harvest) getURL(method string, url string) ([]byte, error) {
	Client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Harvest Slack Bot")
	req.Header.Set("Harvest-Account-ID", h.API.AccountID)
	req.Header.Set("Authorization", "Bearer "+h.API.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := Client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body, err
}

func containsInt(intArray []int64, intSingle int64) bool {
	for _, a := range intArray {
		if a == intSingle {
			return true
		}
	}
	return false
}
