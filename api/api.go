package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

var token string
var apiURL url.URL

func init() {
	token = os.Getenv("API_PARIS_TOKEN")

	v := url.Values{}
	v.Set("token", token)

	apiURL = url.URL{
		Scheme:   "https",
		Host:     "api.paris.fr",
		Path:     "api/data",
		RawQuery: v.Encode(),
	}
}

type APIResponse struct {
	Status      string      `json:"status"`
	Data        interface{} `json:"data"`
	Message     interface{} `json:"message"`
	RequestTime float32     `json:"requestTime"`
	APIVersion  float32     `json:"api-version"`
}

func Get(p string, params map[string]interface{}, resource interface{}) error {
	var err error
	u := buildURL(p, params)

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	apiResponse := APIResponse{Data: resource}
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return err
	}

	if apiResponse.Status == "error" {
		fmt.Println(u.String())
		fmt.Println(apiResponse)
		err = errors.New("Some error")
	}

	return err
}

func buildURL(p string, params map[string]interface{}) url.URL {
	u := apiURL
	v := u.Query()

	var cids string

	for key, value := range params {
		switch value.(type) {
		case string:
			v.Set(key, value.(string))
		case int:
			v.Set(key, strconv.Itoa(value.(int)))
		case []int:
			fmt.Println(value.([]int))
			var strIDs []string
			for _, i := range value.([]int) {
				strIDs = append(strIDs, strconv.Itoa(i))
			}
			fmt.Println(key)
			cids = "cid=" + strings.Join(strIDs, ",")
		}
	}

	u.RawQuery = v.Encode()
	if cids != "" {
		u.RawQuery += "&" + cids
	}
	u.Path = path.Join(u.Path, p)

	return u
}
