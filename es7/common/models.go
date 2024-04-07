package common

type Book struct {
	Author string `json:"author,omitempty"`
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
}

type User struct {
	Name      string `json:"name,omitempty"`
	Introduce string `json:"introduce,omitempty"`
	Class     string `json:"class,omitempty"`
	Age       int    `json:"age,omitempty"`
}
