package generators

import (
	"github.com/koodinikkarit/seppo/models"
	"github.com/koodinikkarit/seppo/seppo_service"
)

func NewEwDatabase(in *models.EwDatabase) *SeppoService.EwDatabase {
	newEwDatabase := SeppoService.EwDatabase{
		Id:                             in.ID,
		Name:                           in.Name,
		SongDatabaseId:                 in.SongDatabaseID,
		DatabaseKey:                    in.EwDatabaseKey,
		VariationVersionConflictAction: uint32(in.VariationVersionConflictAction),
	}

	if in.RemoveSongsFromEwDatabase == 1 {
		newEwDatabase.RemoveSongsFromEwDatabase = 1
	} else {
		newEwDatabase.RemoveSongsFromEwDatabase = 2
	}

	if in.RemoveSongsFromSongDatabase == 1 {
		newEwDatabase.RemoveSongsFromSongDatabase = 1
	} else {
		newEwDatabase.RemoveSongsFromSongDatabase = 2
	}

	return &newEwDatabase
}

func NewTag(in *models.Tag) *SeppoService.Tag {
	return &SeppoService.Tag{
		Id:   in.ID,
		Name: in.Name,
	}
}

func NewLanguage(in *models.Language) *SeppoService.Language {
	return &SeppoService.Language{
		Id:   in.ID,
		Name: in.Name,
	}
}

func NewAuthor(in *models.Author) *SeppoService.Author {
	newAuthor := SeppoService.Author{
		Id:   in.ID,
		Name: in.Name,
	}
	return &newAuthor
}

func NewCopyright(in *models.Copyright) *SeppoService.Copyright {
	newCopyright := SeppoService.Copyright{
		Id:   in.ID,
		Name: in.Name,
	}
	return &newCopyright
}

func NewLog(in *models.Log) *SeppoService.Log {
	return &SeppoService.Log{
		Id:          in.ID,
		LogType:     uint32(in.LogType),
		Message:     in.Message,
		MessageDate: in.MessageDate.Time.Unix() * 1000,
	}
}

func NewSchedule(in *models.Schedule) *SeppoService.Schedule {
	newSchedule := SeppoService.Schedule{
		Id:   in.ID,
		Name: in.Name,
	}

	if in.Start.Valid == true {
		newSchedule.Start = in.Start.Time.Unix() * 1000
	}

	if in.End.Valid == true {
		newSchedule.End = in.End.Time.Unix() * 1000
	}

	return &newSchedule
}

func NewSongDatabase(in *models.SongDatabase) *SeppoService.SongDatabase {
	return &SeppoService.SongDatabase{
		Id:   in.ID,
		Name: in.Name,
	}
}

func NewVariation(in *models.Variation) *SeppoService.Variation {
	newVariation := SeppoService.Variation{
		Id: in.ID,
	}
	if in.SongID.Valid == true {
		newVariation.SongId = in.SongID.Uint64
	}
	if in.LanguageID.Valid == true {
		newVariation.LanguageId = in.LanguageID.Uint64
	}
	if in.AuthorID.Valid == true {
		newVariation.AuthorId = in.AuthorID.Uint64
	}
	if in.CopyrightID.Valid == true {
		newVariation.CopyrightId = in.CopyrightID.Uint64
	}
	return &newVariation
}

func NewVariationVersion(in *models.VariationVersion) *SeppoService.VariationVersion {
	newVariationVersion := SeppoService.VariationVersion{
		Id:          in.ID,
		VariationId: in.VariationID,
		Name:        in.Name,
		Text:        in.Text,
		Version:     uint32(in.Version),
	}
	return &newVariationVersion
}
