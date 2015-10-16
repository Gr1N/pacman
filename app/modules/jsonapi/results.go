package jsonapi

type Item struct {
	Type       string      `json:"type"`
	Id         int64       `json:"id"`
	Attributes interface{} `json:"attributes"`
	Links      ItemLinks   `json:"links"`
}

type ItemLinks struct {
	Self string `json:"self"`
}

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}
