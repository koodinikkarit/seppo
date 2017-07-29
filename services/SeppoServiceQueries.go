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
