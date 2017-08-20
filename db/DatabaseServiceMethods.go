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

func (ds *DatabaseService) RemoveVariation(variationId uint32) bool {
	returnChannel := make(chan bool)
	ds.removeVariationChannel <- removeVariationInternalInput{
		variationID:   variationId,
		returnChannel: returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) UpdateEwSong(input UpdateEwSongInput) *Variation {
	return &Variation{}
}

func (ds *DatabaseService) CreateSongDatabase(createSongDatabaseInput CreateSongDatabaseInput) *SongDatabase {
	returnChnnel := make(chan *SongDatabase)
	ds.createSongDatabaseChannel <- createSongDatabaseInternalInput{
		input:         createSongDatabaseInput,
		returnChannel: returnChnnel,
	}

	return <-returnChnnel
}

func (ds *DatabaseService) EditSongDatabase(editSongdatabaseInput EditSongDatabaseInput) *SongDatabase {
	returnChannel := make(chan *SongDatabase)
	ds.editSongDatabaseChannel <- editSongDatabaseInternalInput{
		input:         editSongdatabaseInput,
		returnChannel: returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) RemoveSongDatabase(songDatabaseId uint32) bool {
	returnChannel := make(chan bool)
	ds.removeSongDatabaseChannel <- removeSongDatabaseInternalInput{
		songDatabaseID: songDatabaseId,
		returnChannel:  returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) CreateEwDatabase(createEwDatabaseInput CreateEwDatabaseInput) *EwDatabase {
	returnChannel := make(chan *EwDatabase)
	ds.createEwDatabaseChannel <- createEwDatabaseInternalInput{
		input:        createEwDatabaseInput,
		returnChnnel: returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) EditEwDatabase(in EditEwDatabaseInput) *EwDatabase {
	returnChannel := make(chan *EwDatabase)
	ds.editEwDatabaseChannel <- editEwDatabaseInternalInput{
		input:         in,
		returnChannel: returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) RemoveEwDatabase(ewDatabaseId uint32) bool {
	returnChannel := make(chan bool)
	ds.removeEwDatabaseChannel <- removeEwDatabaseInternalInput{
		ewDatabaseID:  ewDatabaseId,
		returnChannel: returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) CreateEwDatabaseLink(ewDatabaseID uint32, ewDatabaseSongID uint32, variationID uint32, version uint64) *EwDatabaseLink {
	returnChannel := make(chan *EwDatabaseLink)
	ds.createEwDatabaseLinkChannel <- createEwDatabaseLinkInternalInput{
		ewDatabaseID:     ewDatabaseID,
		ewDatabaseSongID: ewDatabaseSongID,
		variationID:      variationID,
		version:          version,
		returnChannel:    returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) EditEwDatabaseLink(in EditEwDatabaseLinkInput) *EwDatabaseLink {
	returnChannel := make(chan *EwDatabaseLink)
	ds.editEwDatabaseLinkChannel <- editEwDatabaseLinkInternalInput{
		input:         in,
		returnChannel: returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) RemoveDatabaseLink(ewDatabaseLinkID uint32) bool {
	returnChannel := make(chan bool)
	ds.removeEwDatabaseLinkChannel <- removeEwDatabaseLinkInternalInput{
		ewDatabaseLinkID: ewDatabaseLinkID,
		returnChnnel:     returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) AddVariationToSongDatabase(songDatabaseID uint32, variationID uint32) *SongDatabaseVariation {
	returnChannel := make(chan *SongDatabaseVariation)
	ds.addVariationToSongDatabaseChannel <- addVariationToSongDatabaseInternalInput{
		songDatabaseID: songDatabaseID,
		variationID:    variationID,
		returnChannel:  returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) RemoveVariationFromSongDatabase(songDatabaseID uint32, variationID uint32) bool {
	returnChannel := make(chan bool)
	ds.removeVariationFromSongDatabaseChannel <- removeVariationFromSongDatabaseInternalInput{
		songDatabaseID: songDatabaseID,
		variationID:    variationID,
		returnChannel:  returnChannel,
	}
	return <-returnChannel
}

func (ds *DatabaseService) RemoveEwSong(songDatabaseID uint32, ewSongID uint32) bool {
	returnChannel := make(chan bool)
	ds.removeEwSongChannel <- removeEwSongInternalInput{
		ewSongID:       ewSongID,
		songDatabaseID: songDatabaseID,
		returnChannel:  returnChannel,
	}
	return <-returnChannel
}
