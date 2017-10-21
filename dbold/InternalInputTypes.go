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

type createEwDatabaseLinkInternalInput struct {
	ewDatabaseID     uint32
	ewDatabaseSongID uint32
	variationID      uint32
	version          uint64
	returnChannel    chan *EwDatabaseLink
}

type editEwDatabaseLinkInternalInput struct {
	input         EditEwDatabaseLinkInput
	returnChannel chan *EwDatabaseLink
}

type removeEwDatabaseLinkInternalInput struct {
	ewDatabaseLinkID uint32
	returnChnnel     chan bool
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

type removeEwSongInternalInput struct {
	ewSongID       uint32
	songDatabaseID uint32
	returnChannel  chan bool
}

type createTagInternalInput struct {
	input         CreateTagInput
	returnChannel chan *Tag
}

type editTagInternalInput struct {
	input         EditTagInput
	returnChannel chan *Tag
}

type removeTagInternalInput struct {
	tagID         uint32
	returnChannel chan bool
}

type createLanguageInternalInput struct {
	input         CreateLanguageInput
	returnChannel chan *Language
}

type editLanguageInternalInput struct {
	input         EditLanguageInput
	returnChannel chan *Language
}

type removeLanguageInternalInput struct {
	languageID    uint32
	returnChannel chan bool
}

type addTagToVariationInternalInput struct {
	tagID         uint32
	variationID   uint32
	returnChannel chan *TagVariation
}

type removeTagFromVariationInternalInput struct {
	tagID         uint32
	variationID   uint32
	returnChannel chan bool
}

type addTagToSongDatabaseInternalInput struct {
	tagID          uint32
	songDatabaseID uint32
	returnChannel  chan *SongDatabaseTag
}

type removeTagFromSongDatabaseInternalInput struct {
	tagID          uint32
	songDatabaseID uint32
	returnChannel  chan bool
}

type createScheduleInternalInput struct {
	input         CreateScheduleInput
	returnChannel chan *Schedule
}

type updateScheduleInternalInput struct {
	input         UpdateScheduleInput
	returnChannel chan *Schedule
}

type removeScheduleInternalInput struct {
	scheduleID    uint32
	returnChannel chan bool
}
