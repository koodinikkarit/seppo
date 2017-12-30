package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/models"
)

func MoveVariationReferences(
	tx *gorm.DB,
	srcVariationID uint32,
	dstVariationID uint32,
) {
	tagVariations := []models.TagVariation{}
	tx.Where("variation_id = ?", srcVariationID).Find(&tagVariations)
	var newTagVariations []models.TagVariation
	for i := 0; i < len(tagVariations); i++ {
		newTagVariations = append(
			newTagVariations,
			models.TagVariation{
				TagID:       tagVariations[i].TagID,
				VariationID: dstVariationID,
			},
		)
	}
	BatchAddTagsToVariation(
		tx,
		newTagVariations,
	)

	songDatabaseVariations := []models.SongDatabaseVariation{}
	tx.Where("variation_version_id = ?", srcVariationID).Find(&songDatabaseVariations)
	var newSongDatabaseVariations []models.SongDatabaseVariation
	for i := 0; i < len(songDatabaseVariations); i++ {
		newSongDatabaseVariations = append(
			newSongDatabaseVariations,
			models.SongDatabaseVariation{
				SongDatabaseID: songDatabaseVariations[i].SongDatabaseID,
				VariationID:    dstVariationID,
			},
		)
	}
	BatchAddVariationsToSongDatabase(
		tx,
		newSongDatabaseVariations,
	)

	scheduleVariations := []models.ScheduleVariation{}
	tx.Where("variation_version_id = ?", srcVariationID).Find(&scheduleVariations)
	var newScheduleVariations []models.ScheduleVariation
	for i := 0; i < len(scheduleVariations); i++ {
		newScheduleVariations = append(
			newScheduleVariations,
			models.ScheduleVariation{
				ScheduleID:  scheduleVariations[i].ScheduleID,
				VariationID: dstVariationID,
				OrderNumber: scheduleVariations[i].OrderNumber,
			},
		)
	}
	BatchCreateScheduleVariations(
		tx,
		newScheduleVariations,
	)
}
