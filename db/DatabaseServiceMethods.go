package SeppoDB

func (ds *DatabaseService) CreateVariation(input CreateVariationInput) *Variation {
	returnChannel := make(chan *Variation)
	ds.createVariationChannel <- createVariationInternalInput{
		input:         input,
		returnChannel: returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) EditVariation(input EditVariationInput) *Variation {
	returnChannel := make(chan *Variation)
	ds.editVariationChannel <- editVariationInternalInput{
		input:         input,
		returnChannel: returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) UpdateEwSong(input UpdateEwSongInput) *Variation {
	return &Variation{}
}
