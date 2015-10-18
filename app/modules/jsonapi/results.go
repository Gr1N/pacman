package jsonapi

type ResultIndividual struct {
	Data *Item `json:"data"`
}

type ResultCollection struct {
	Data []*Item `json:"data"`
}

type Item struct {
	Type       string      `json:"type"`
	Id         int64       `json:"id"`
	Attributes interface{} `json:"attributes"`
	Links      ItemLinks   `json:"links"`
}

type ItemLinks struct {
	Self string `json:"self"`
}

type ResultError struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}
