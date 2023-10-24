package player

type playerDomain struct {
	id       string
	name     string
	photo    string
	height   float32
	weight   float32
	age      int32
	position string
	number   int32
	teamID   string
}

func (player *playerDomain) GetID() string {
	return player.id
}

func (player *playerDomain) GetName() string {
	return player.name
}

func (player *playerDomain) GetPhoto() string {
	return player.photo
}

func (player *playerDomain) GetHeight() float32 {
	return player.height
}

func (player *playerDomain) GetWeight() float32 {
	return player.weight
}
func (player *playerDomain) GetAge() int32 {
	return player.age
}
func (player *playerDomain) GetPosition() string {
	return player.position
}

func (player *playerDomain) GetNumber() int32 {
	return player.number
}

func (player *playerDomain) GetTeamID() string {
	return player.teamID
}

func (player *playerDomain) SetID(id string) {
	player.id = id
}
