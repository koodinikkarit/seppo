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
		input: createEwDatabaseInput,
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
