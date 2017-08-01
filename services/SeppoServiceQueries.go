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

	s.databaseService.GetDb().Find(&variations)

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
	s.databaseService.GetDb().Find(&songDatabases)

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
