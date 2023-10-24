package team

func NewTeamDomain(name, badgePhoto, city, coach, website string) TeamDomainInterface {
	return &teamDomain{
		Name:       name,
		BadgePhoto: badgePhoto,
		City:       city,
		Coach:      coach,
		Website:    website,
	}
}
