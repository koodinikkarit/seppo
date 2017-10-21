package services

import (
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/seppo_service"
)

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

func NewLog(in *db.Log) *SeppoService.Log {
	return &SeppoService.Log{
		Id:          in.ID,
		LogType:     in.LogType,
		Message:     in.Message,
		MessageDate: in.MessageDate.Unix() * 1000,
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

	return &newVariation
}
