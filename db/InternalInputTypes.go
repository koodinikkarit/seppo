package SeppoDB

type createVariationInternalInput struct {
	input         CreateVariationInput
	returnChannel chan *Variation
}

type editVariationInternalInput struct {
	input         EditVariationInput
	returnChannel chan *Variation
}

type removeVariationInternalInput struct {
	variationID   uint32
	returnChannel chan bool
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

type editEwDatabaseInternalInput struct {
	input         EditEwDatabaseInput
	returnChannel chan *EwDatabase
}

type removeEwDatabaseInternalInput struct {
	ewDatabaseID  uint32
	returnChannel chan bool
}

type addVariationToSongDatabaseInternalInput struct {
	songDatabaseID uint32
	variationID    uint32
	returnChannel  chan *SongDatabaseVariation
}

type removeVariationFromSongDatabaseInternalInput struct {
	songDatabaseID uint32
	variationID    uint32
	returnChannel  chan bool
}
