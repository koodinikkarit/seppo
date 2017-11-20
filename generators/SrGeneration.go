package generators

import "github.com/koodinikkarit/seppo/models"

func NewSrVariationFromVariation(
	variation *models.Variation,
) *models.SRNewVariation {
	return &models.SRNewVariation{
		VariationID: variation.ID,
	}
}

func NewSrVariationVersionFromVariationVersion(
	variationVersion *models.VariationVersion,
) *models.SRNewVariationVersion {
	return &models.SRNewVariationVersion{
		VariationVersionID: variationVersion.ID,
	}
}

func NewSrPassivatedVariationVersionFromVariationVersion(
	variationVersion *models.VariationVersion,
) *models.SRPassivatedVariationVersion {
	return &models.SRPassivatedVariationVersion{
		VariationVersionID: variationVersion.ID,
	}
}

func NewSrBranchFromBranch(
	branch *models.Branch,
) *models.SRNewBranch {
	return &models.SRNewBranch{
		BranchID: branch.ID,
	}
}

func NewSrSongDatabaseVariation(
	songDatabaseVariation *models.SongDatabaseVariation,
) *models.SRAddSongDatabaseVariation {
	return &models.SRAddSongDatabaseVariation{
		SongDatabaseID: songDatabaseVariation.SongDatabaseID,
		VariationID:    songDatabaseVariation.VariationID,
	}
}

func NewSrEwDatabaseLinkFromEwDatabaseLink(
	ewDatabaseLink *models.EwDatabaseLink,
	operation int8,
) *models.SREwDatabaseLink {
	return &models.SREwDatabaseLink{
		EwDatabaseID:     ewDatabaseLink.EwDatabaseID,
		EwDatabaseSongID: ewDatabaseLink.EwDatabaseSongID,
		VariationID:      ewDatabaseLink.VariationID,
		Version:          ewDatabaseLink.Version,
		Author:           ewDatabaseLink.Author,
		Copyright:        ewDatabaseLink.Copyright,
		Operation:        operation,
	}
}
