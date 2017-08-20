package seppo

import (
	"github.com/koodinikkarit/seppo/db"
	"github.com/koodinikkarit/seppo/seppo_service"
)

func NewVariationToServiceType(in *SeppoDB.Variation) *SeppoService.Variation {
	return &SeppoService.Variation{
		Id:      in.ID,
		Name:    in.Name,
		SongId:  in.SongID,
		Version: in.Version,
	}
}

func NewVariationTextToServiceType(in *SeppoDB.VariationText) *SeppoService.VariationText {
	return &SeppoService.VariationText{
		Id:          in.ID,
		VariationId: in.VariationID,
		Text:        in.Text,
	}
}

func NewSongDatabaseToServiceType(in *SeppoDB.SongDatabase) *SeppoService.SongDatabase {
	return &SeppoService.SongDatabase{
		Id:   in.ID,
		Name: in.Name,
	}
}

func NewEwDatabaseToServiceType(in *SeppoDB.EwDatabase) *SeppoService.EwDatabase {
	return &SeppoService.EwDatabase{
		Id:             in.ID,
		SongDatabaseId: in.SongDatabaseID,
		Name:           in.Name,
		Key:            in.Key,
	}
}

func NewSongDatabaseVariationToServiceType(in *SeppoDB.SongDatabaseVariation) *SeppoService.SongDatabaseVariation {
	return &SeppoService.SongDatabaseVariation{
		Id:             in.ID,
		SongDatabaseId: in.SongDatabaseID,
		VariationId:    in.VariationID,
	}
}
