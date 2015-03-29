package models

import (
	"encoding/json"

	"github.com/delba/api-paris/api"
	"github.com/garyburd/redigo/redis"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) UnmarshalJSON(data []byte) error {
	var err error

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		return err
	}

	for key, value := range dataMap {
		switch key {
		case "idcategories":
			c.ID = int(value.(float64))
		case "name":
			c.Name = value.(string)
		}
	}

	return err
}

func (cat *Category) Save() error {
	isMember, err := redis.Bool(c.Do("SISMEMBER", "categories:ids", cat.ID))
	if err != nil {
		return err
	}

	if isMember {
		return err
	}

	c.Send("MULTI")
	c.Send("HMSET", redis.Args{}.Add("categories:"+string(cat.ID)).AddFlat(cat)...)
	c.Send("SADD", "categories:ids", cat.ID)
	_, err = c.Do("EXEC")

	return err
}

func (cat *Category) FetchFacilities(offset int, limit int) (Facilities, error) {
	var facilities Facilities
	var err error

	params := map[string]interface{}{
		"cid":    cat.ID,
		"offset": offset,
		"limit":  limit,
	}

	err = api.Get("Equipements/get_equipements", params, &facilities)

	return facilities, err
}

type Categories []Category

func (cat *Categories) Fetch() error {
	var err error

	err = api.Get("Equipements/get_categories", nil, cat)

	return err
}

func (cat *Categories) All() error {
	ids, err := redis.Ints(c.Do("SMEMBERS", "categories:ids"))
	if err != nil {
		return err
	}

	c.Send("MULTI")

	for _, id := range ids {
		c.Send("HGETALL", "categories:"+string(id))
	}

	values, err := redis.Values(c.Do("EXEC"))
	if err != nil {
		return err
	}

	var category Category

	for _, v := range values {
		values, err = redis.Values(v, nil)
		if err != nil {
			return err
		}

		err = redis.ScanStruct(values, &category)
		if err != nil {
			return err
		}

		*cat = append(*cat, category)
	}

	return err
}
