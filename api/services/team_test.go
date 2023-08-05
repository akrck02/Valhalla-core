package services

import ()

/*
func TestCreateTeam(t *testing.T) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Creating owner
	var user = models.User{
		Email:    mock.Email(),
		Password: mock.Password(),
		Username: mock.Username(),
	}

	log.FormattedInfo("Registering user: ${0}", user.Email)
	log.FormattedInfo("Password: ${0}", user.Password)
	log.FormattedInfo("Username: ${0}", user.Username)

	err := Register(conn, client, user)

	if err != nil {
		t.Error("The user was not registered", err)
		return
	}

	log.FormattedInfo("Getting user.")

	user := GetUser(conn, client, user)

	if user.ID == "" {
		t.Error("Could not get the user registered")
	}

	var team = models.Team{
		Name:        mock.Name(),
		Description: mock.Description(),
		Owner:       user.ID,
	}

	log.FormattedInfo("Creating team: ${0}", team.Name)
	log.FormattedInfo("Description: ${0}", team.Description)
	log.FormattedInfo("Owner: ${0}", team.Owner)

	err := CreateTeam(conn, client, team)

	if err != nil {
		t.Error("The team was not created ", err)
		return
	}

	log.FormattedInfo("Team created.")

	// Delete the team

	log.FormattedInfo("Deleting team: ${0}", team.Name)

	err = DeleteTeam(conn, client, team)

	if err != nil {
		t.Error("The team was not deleted ", err)
		return
	}

	log.FormattedInfo("Deleting user: ${0}", user.Name)

	err = DeleteUser(conn, client, user)

	if err != nil {
		t.Error("The user was not deleted ", err)
		return
	}

	log.FormattedInfo("Team deleted.")
}
*/
