package seppo

import (
	"golang.org/x/net/context"

	"github.com/koodinikkarit/seppo/db"
	SeppoService "github.com/koodinikkarit/seppo/seppo_service"
)

func (s *SeppoServiceServer) FetchVariationById(ctx context.Context, in *SeppoService.FetchVariationByIdRequest) (*SeppoService.FetchVariationByIdResponse, error) {
	res := &SeppoService.FetchVariationByIdResponse{}

	variations := []SeppoDB.Variation{}
	s.databaseService.GetDb().Where("id in (?)", in.VariationIds).Find(&variations)
	for _, variationId := range in.VariationIds {
		var found bool
		for _, variation := range variations {
			if variation.ID == variationId {
				res.Variations = append(res.Variations, NewVariationToServiceType(&variation))
				found = true
				break
			}
		}
		if found == false {
			res.Variations = append(res.Variations, &SeppoService.Variation{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) SearchVariations(ctx context.Context, in *SeppoService.SearchVariationsRequest) (*SeppoService.SearchVariationsResponse, error) {
	res := &SeppoService.SearchVariationsResponse{}
	variations := []SeppoDB.Variation{}

	query := s.databaseService.GetDb().Table("variations")

	if in.SongDatabaseFilterId > 0 {
		var filterSongDatabaseVariationsIds []uint32
		filterSongDatabaseVariations := []SeppoDB.SongDatabaseVariation{}
		s.databaseService.GetDb().
			Where("song_database_id = ?", in.SongDatabaseFilterId).
			Select("variation_id").Find(&filterSongDatabaseVariations)
		for _, v := range filterSongDatabaseVariations {
			filterSongDatabaseVariationsIds = append(filterSongDatabaseVariationsIds, v.VariationID)
		}
		if filterSongDatabaseVariationsIds != nil {
			query = query.Not("id", filterSongDatabaseVariationsIds)
		}
	}
	query.Count(&res.MaxVariations)

	if in.SearchWord != "" {
		query = query.Joins("JOIN variation_texts ON variation_texts.variation_id = variations.id").
			Where("variations.name LIKE ?", "%"+in.SearchWord+"%").
			Or("variation_texts.text LIKE ?", "%"+in.SearchWord+"%")
	}

	query = query.Select("variations.id, variations.name, variations.song_id, variations.version").Find(&variations)

	for _, variation := range variations {
		res.Variations = append(res.Variations, &SeppoService.Variation{
			Id:      variation.ID,
			Name:    variation.Name,
			SongId:  variation.SongID,
			Version: variation.Version,
		})
	}

	return res, nil
}

func (s *SeppoServiceServer) ListenForChangedEwSong(in *SeppoService.ListenForChangedEwSongRequest, stream SeppoService.Seppo_ListenForChangedEwSongServer) error {

	return nil
}

func (s *SeppoServiceServer) FetchSongDatabases(ctx context.Context, in *SeppoService.FetchSongDatabasesRequest) (*SeppoService.SongDatabasesConnection, error) {
	res := &SeppoService.SongDatabasesConnection{}
	songDatabases := []SeppoDB.SongDatabase{}

	query := s.databaseService.GetDb().Debug()

	if in.SearchWord != "" {
		query = query.Where("name LIKE ?", "%"+in.SearchWord+"%")
	}

	query = query.Find(&songDatabases)

	for _, songDatabase := range songDatabases {
		res.Edges = append(res.Edges, &SeppoService.SongDatabaseEdge{
			Node: NewSongDatabaseToServiceType(&songDatabase),
		})
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchSongDatabaseById(ctx context.Context, in *SeppoService.FetchSongDatabaseByIdRequest) (*SeppoService.FetchSongDatabaseByIdResponse, error) {
	res := &SeppoService.FetchSongDatabaseByIdResponse{}

	songDatabases := []SeppoDB.SongDatabase{}
	s.databaseService.GetDb().Where("id in (?)", in.SongDatabaseIds).Find(&songDatabases)

	for _, songDatabaseId := range in.SongDatabaseIds {
		found := false
		for _, songDatabase := range songDatabases {
			if songDatabase.ID == songDatabaseId {
				found = true
				res.SongDatabases = append(res.SongDatabases, NewSongDatabaseToServiceType(&songDatabase))
			}
		}
		if found == false {
			res.SongDatabases = append(res.SongDatabases, &SeppoService.SongDatabase{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchEwDatabases(ctx context.Context, in *SeppoService.FetchEwDatabasesRequest) (*SeppoService.EwDatabasesConnection, error) {
	res := &SeppoService.EwDatabasesConnection{}

	ewDatabases := []SeppoDB.EwDatabase{}
	s.databaseService.GetDb().Find(&ewDatabases)

	for _, ewDatabase := range ewDatabases {
		res.EwDatabases = append(res.EwDatabases, NewEwDatabaseToServiceType(&ewDatabase))
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchEwDatabaseById(ctx context.Context, in *SeppoService.FetchEwDatabaseByIdRequest) (*SeppoService.FetchEwDatabaseByIdResponse, error) {
	res := &SeppoService.FetchEwDatabaseByIdResponse{}

	ewDatabases := []SeppoDB.EwDatabase{}
	s.databaseService.GetDb().Where("id in (?)", in.EwDatabaseIds).Find(&ewDatabases)
	for _, ewDatabaseId := range in.EwDatabaseIds {
		found := false
		for _, ewDatabase := range ewDatabases {
			if ewDatabaseId == ewDatabase.ID {
				found = true
				res.EwDatabases = append(res.EwDatabases, NewEwDatabaseToServiceType(&ewDatabase))
			}
		}
		if found == false {
			res.EwDatabases = append(res.EwDatabases, &SeppoService.EwDatabase{
				Id: 0,
			})
		}
	}
	return res, nil
}

func (s *SeppoServiceServer) FetchVariationsBySongDatabaseId(ctx context.Context, in *SeppoService.FetchVariationsBySongDatabaseIdRequest) (*SeppoService.FetchVariationsBySongDatabaseIdResponse, error) {
	res := &SeppoService.FetchVariationsBySongDatabaseIdResponse{}

	fetchedSongDatabaseVariations := []SeppoDB.SongDatabaseVariation{}
	s.databaseService.GetDb().Where("song_database_id in (?)", in.SongDatabaseIds).Find(&fetchedSongDatabaseVariations)

	variationIds := []uint32{}
	for _, v := range fetchedSongDatabaseVariations {
		variationIds = append(variationIds, v.VariationID)
	}
	variations := []SeppoDB.Variation{}
	s.databaseService.GetDb().Where("id in (?)", variationIds).Find(&variations)

	for _, songDatabaseId := range in.SongDatabaseIds {
		databaseVariations := SeppoService.SongDatabaseVariations{}
		databaseVariations.SongDatabaseId = songDatabaseId
		for _, songDatabaseVariation := range fetchedSongDatabaseVariations {
			if songDatabaseId == songDatabaseVariation.SongDatabaseID {
				for _, variation := range variations {
					if songDatabaseVariation.VariationID == variation.ID {
						databaseVariations.Variations = append(databaseVariations.Variations, NewVariationToServiceType(&variation))
					}
				}
			}
		}
		res.SongDatabaseVariations = append(res.SongDatabaseVariations, &databaseVariations)
	}

	return res, nil
}

func (s *SeppoServiceServer) FetchVariationTextByVariationId(ctx context.Context, in *SeppoService.FetchVariationTextByVariationIdRequest) (*SeppoService.FetchVariationTextByVariationIdResponse, error) {
	res := &SeppoService.FetchVariationTextByVariationIdResponse{}

	variationTexts := []SeppoDB.VariationText{}

	s.databaseService.GetDb().Where("variation_id in (?)", in.VariationIds).Find(&variationTexts)

	for _, variationId := range in.VariationIds {
		found := false
		for _, variationText := range variationTexts {
			if variationText.VariationID == variationId {
				found = true
				res.VariationTexts = append(res.VariationTexts, NewVariationTextToServiceType(&variationText))
			}
		}

		if found == false {
			res.VariationTexts = append(res.VariationTexts, &SeppoService.VariationText{
				Id: 0,
			})
		}
	}

	return res, nil
}
