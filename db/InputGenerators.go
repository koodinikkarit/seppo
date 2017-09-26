package SeppoDB

import (
	"github.com/koodinikkarit/seppo/seppo_service"
)

func NewCreateVariationFromServiceType(in *SeppoService.CreateVariationRequest) CreateVariationInput {
	return CreateVariationInput{
		Name:   in.Name,
		Text:   in.Text,
		Origin: in.Origin,
	}
}

func NewEditVariationFromService(in *SeppoService.UpdateVariationRequest) EditVariationInput {
	return EditVariationInput{
		VariationID: in.VariationId,
		Name:        in.Name,
		Text:        in.Text,
		SongID:      in.SongId,
		LanguageID:  in.LanguageId,
	}
}

func NewSongDatabaseFromServiceType(in *SeppoService.SongDatabase) SongDatabase {
	return SongDatabase{
		ID:   in.Id,
		Name: in.Name,
	}
}

func NewCreateSongDatabaseInputFromServiceType(in *SeppoService.CreateSongDatabaseRequest) CreateSongDatabaseInput {
	return CreateSongDatabaseInput{
		Name: in.Name,
	}
}

func NewEditSongDatabaseInputFromServiceType(in *SeppoService.UpdateSongDatabaseRequest) EditSongDatabaseInput {
	return EditSongDatabaseInput{
		SongDatabaseId: in.SongDatabaseId,
		Name:           in.Name,
	}
}

func NewEwDatabaseFromServiceType(in *SeppoService.EwDatabase) EwDatabase {
	return EwDatabase{
		ID:             in.Id,
		SongDatabaseID: in.SongDatabaseId,
		Name:           in.Name,
	}
}

func NewCreateEwDatabaseFromServiceType(in *SeppoService.CreateEwDatabaseRequest) CreateEwDatabaseInput {
	return CreateEwDatabaseInput{
		Name:           in.Name,
		SongDatabaseId: in.SongDatabaseId,
	}
}

func NewEditEwDatabaseFromServiceType(in *SeppoService.UpdateEwDatabaseRequest) EditEwDatabaseInput {
	return EditEwDatabaseInput{
		EwDatabaseID:   in.EwDatabaseId,
		Name:           in.Name,
		SongDatabaseID: in.SongDatabaseId,
	}
}
