package team

type TeamDomainInterface interface {
	SetID(id string)
	GetName() string
	GetBadgePhoto() string
	GetCity() string
	GetCoach() string
	GetWebSite() string
	GetID() string
}
