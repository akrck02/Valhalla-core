package models

type Project struct {
	Name        string   `bson:"name,omitempty"`
	Description string   `bson:"description,omitempty"`
	Owner       string   `bson:"owner,omitempty"`
	Teams       []string `bson:"teams,omitempty"`
	Wikis       []string `bson:"wikis,omitempty"`
	Notes       []string `bson:"notes,omitempty"`
	Tasks       []string `bson:"tasks,omitempty"`
}
