package generators

import (
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/seppo_service"
)

func NewEwDatabase(in *db.EwDatabase) *SeppoService.EwDatabase {
	newEwDatabase := SeppoService.EwDatabase{
		Id:                             in.ID,
		Name:                           in.Name,
		SongDatabaseId:                 in.SongDatabaseID,
		DatabaseKey:                    in.EwDatabaseKey,
		VariationVersionConflictAction: uint32(in.VariationVersionConflictAction),
	}

	if in.RemoveSongsFromEwDatabase == true {
		newEwDatabase.RemoveSongsFromEwDatabase = 1
	} else {
		newEwDatabase.RemoveSongsFromEwDatabase = 2
	}

	if in.RemoveSongsFromSongDatabase == true {
		newEwDatabase.RemoveSongsFromSongDatabase = 1
	} else {
		newEwDatabase.RemoveSongsFromSongDatabase = 2
	}

	return &newEwDatabase
}

func NewTag(in *db.Tag) *SeppoService.Tag {
	return &SeppoService.Tag{
		Id:   in.ID,
		Name: in.Name,
	}
}

func NewLanguage(in *db.Language) *SeppoService.Language {
	return &SeppoService.Language{
		Id:   in.ID,
		Name: in.Name,
	}
}

func NewAuthor(in *db.Author) *SeppoService.Author {
	newAuthor := SeppoService.Author{
		Id:   in.ID,
		Name: in.Name,
	}
	return &newAuthor
}

func NewCopyright(in *db.Copyright) *SeppoService.Copyright {
	newCopyright := SeppoService.Copyright{
		Id:   in.ID,
		Name: in.Name,
	}
	return &newCopyright
}

func NewLog(in *db.Log) *SeppoService.Log {
	return &SeppoService.Log{
		Id:          in.ID,
		LogType:     uint32(in.LogType),
		Message:     in.Message,
		MessageDate: in.MessageDate.Unix() * 1000,
	}
}

func NewSchedule(in *db.Schedule) *SeppoService.Schedule {
	newSchedule := SeppoService.Schedule{
		Id:   in.ID,
		Name: in.Name,
	}

	if in.Start != nil {
		newSchedule.Start = in.Start.Unix() * 1000
	}

	if in.End != nil {
		newSchedule.End = in.End.Unix() * 1000
	}

	return &newSchedule
}

func NewSongDatabase(in *db.SongDatabase) *SeppoService.SongDatabase {
	return &SeppoService.SongDatabase{
		Id:   in.ID,
		Name: in.Name,
	}
}

func NewVariation(in *db.Variation) *SeppoService.Variation {
	newVariation := SeppoService.Variation{
		Id: in.ID,
	}
	if in.SongID != nil {
		newVariation.SongId = *in.SongID
	}
	if in.LanguageID != nil {
		newVariation.LanguageId = *in.LanguageID
	}
	if in.AuthorID != nil {
		newVariation.AuthorId = *in.AuthorID
	}
	if in.CopyrightID != nil {
		newVariation.CopyrightId = *in.CopyrightID
	}
	return &newVariation
}

func NewVariationVersion(in *db.VariationVersion) *SeppoService.VariationVersion {
	newVariationVersion := SeppoService.VariationVersion{
		Id:          in.ID,
		VariationId: in.VariationID,
		Name:        in.Name,
		Text:        in.Text,
		Version:     uint32(in.Version),
	}
	return &newVariationVersion
}
