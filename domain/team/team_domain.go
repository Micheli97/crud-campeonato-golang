package team

type teamDomain struct {
	ID         string `sql:"id"`
	Name       string
	BadgePhoto string
	City       string
	Coach      string
	Website    string
}

func (team *teamDomain) SetID(id string) {
	team.ID = id
}

func (team *teamDomain) GetName() string {
	return team.Name
}

func (team *teamDomain) GetBadgePhoto() string {
	return team.BadgePhoto
}

func (team *teamDomain) GetCity() string {
	return team.City
}

func (team *teamDomain) GetCoach() string {
	return team.Coach
}

func (team *teamDomain) GetWebSite() string {
	return team.Website
}

func (team *teamDomain) GetID() string {
	return team.ID
}
