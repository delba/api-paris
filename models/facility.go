package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Facility struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	ZipCode     string  `json:"zip_code"`
	AddressInfo string  `json:"address_info"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	URL         string  `json:"url"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	CrowdLevel  int     `json:"crowd_level"`
	CategoryID  int     `json:"category_id"`
}

func (f *Facility) UnmarshalJSON(data []byte) error {
	var err error

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		return err
	}

	fmt.Println(dataMap)

	for key, value := range dataMap {
		switch key {
		case "id":
			f.ID = int(value.(float64))
		case "name":
			f.Name = value.(string)
		case "description":
			f.Description = value.(string)
		case "address":
			f.Address = value.(string)
		case "zipCode":
			f.ZipCode = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		case "addressInfo":
			f.AddressInfo = value.(string)
		case "phone":
			f.Phone = value.(string)
		case "email":
			f.Email = value.(string)
		case "websiteUrl":
			f.URL = value.(string)
		case "lat":
			f.Latitude = value.(float64)
		case "lon":
			f.Longitude = value.(float64)
		}
	}

	return err
}

type Facilities []Facility
