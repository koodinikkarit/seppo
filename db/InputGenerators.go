package SeppoDB

import (
	"github.com/koodinikkarit/seppo/seppo_service"
)

func NewCreateVariationFromServiceType(in *SeppoService.CreateVariationRequest) CreateVariationInput {
	return CreateVariationInput{
		Name: in.Name,
		Text: in.Text,
	}
}

func NewEditVariationFromService(in *SeppoService.EditVariationRequest) EditVariationInput {
	return EditVariationInput{
		VariationID: in.VariationId,
		Name:        in.Name,
		Text:        in.Text,
		SongID:      in.SongId,
	}
}
