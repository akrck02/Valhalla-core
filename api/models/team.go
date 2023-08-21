package models

// add creation date
type Team struct {
	Name         string   `bson:"name,omitempty"`
	Description  string   `bson:"description,omitempty"`
	ProfilePic   string   `bson:"profilepic,omitempty"`
	Projects     []string `bson:"projects,omitempty"`
	Owner        string   `bson:"owner,omitempty"`
	Members      []string `bson:"members,omitempty"`
	CreationDate string   `bson:"creationdate,omitempty"`
	LastUpdate   string   `bson:"updatedate,omitempty"`
	ID           string   `bson:"_id,omitempty"`
}

func (t *Team) Clone() *Team {
	return &Team{
		Name:        t.Name,
		Description: t.Description,
		ProfilePic:  t.Description,
		Projects:    t.Projects,
		Owner:       t.Owner,
		Members:     t.Members,
		ID:          t.ID,
	}
}
