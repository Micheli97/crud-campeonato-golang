package player

func NewPlayerDomain(name, photo, position, teamID string, number, age int32, weight float32, height float32) PlayerDomainInterface {
	return &playerDomain{
		name:     name,
		position: position,
		teamID:   teamID,
		number:   number,
		weight:   weight,
		height:   height,
		age:      age,
		photo:    photo,
	}
}
