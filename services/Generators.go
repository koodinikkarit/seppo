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
		Text:    in.Text,
		Version: in.Version,
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
	}
}
