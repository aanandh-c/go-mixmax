package models

type ContactGroups struct {
	Results []*ContactGroup `json:"results"`
	*ResultList
}

type ContactGroup struct {
	ID      string `json:"_id"`
	UserID  string `json:"userId"`
	GroupID string `json:"groupId"`
	Name    string `json:"name"`
	Count   int    `json:"count"`
	Shared  Shared `json:"shared"`
}
