package SeppoDB

type createVariationInternalInput struct {
	input         CreateVariationInput
	returnChannel chan *Variation
}

type editVariationInternalInput struct {
	input         EditVariationInput
	returnChannel chan *Variation
}

type createSongDatabaseInternalInput struct {
	input         CreateSongDatabaseInput
	returnChannel chan *SongDatabase
}

type editSongDatabaseInternalInput struct {
	input         EditSongDatabaseInput
	returnChannel chan *SongDatabase
}

type removeSongDatabaseInternalInput struct {
	songDatabaseID uint32
	returnChannel  chan bool
}

type createEwDatabaseInternalInput struct {
	input        CreateEwDatabaseInput
	returnChnnel chan *EwDatabase
}

type removeEwDatabaseInternalInput struct {
	ewDatabaseID  uint32
	returnChannel chan bool
}
