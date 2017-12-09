package managers

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func MoveVariationReferences(
	tx *gorm.DB,
	srcVariationID uint32,
	dstVariationID uint32,
) {
	tagVariations := []db.TagVariation{}
	tx.Where("variation_id = ?", srcVariationID).Find(&tagVariations)
	var newTagVariations []db.TagVariation
	for i := 0; i < len(tagVariations); i++ {
		newTagVariations = append(
			newTagVariations,
			db.TagVariation{
				TagID:       tagVariations[i].TagID,
				VariationID: dstVariationID,
			},
		)
	}
	BatchAddTagsToVariation(
		tx,
		newTagVariations,
	)

	songDatabaseVariations := []db.SongDatabaseVariation{}
	tx.Where("variation_version_id = ?", srcVariationID).Find(&songDatabaseVariations)
	var newSongDatabaseVariations []db.SongDatabaseVariation
	for i := 0; i < len(songDatabaseVariations); i++ {
		newSongDatabaseVariations = append(
			newSongDatabaseVariations,
			db.SongDatabaseVariation{
				SongDatabaseID: songDatabaseVariations[i].SongDatabaseID,
				VariationID:    dstVariationID,
			},
		)
	}
	BatchAddVariationsToSongDatabase(
		tx,
		newSongDatabaseVariations,
	)

	scheduleVariations := []db.ScheduleVariation{}
	tx.Where("variation_version_id = ?", srcVariationID).Find(&scheduleVariations)
	var newScheduleVariations []db.ScheduleVariation
	for i := 0; i < len(scheduleVariations); i++ {
		newScheduleVariations = append(
			newScheduleVariations,
			db.ScheduleVariation{
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
