package services

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/seppo/db"
)

func MergeVariationVersion(
	tx *gorm.DB,
	srcVariationVersion db.VariationVersion,
	dstVariationVersion db.VariationVersion,
) {

}

func MoveVariationVersionReferences(
	tx *gorm.DB,
	srcID uint32,
	dstID uint32,
) {
	tagVariations := []db.TagVariation{}
	tx.Where("variation_version_id = ?", srcID).Find(&tagVariations)

	for i := 0; i < len(tagVariations); i++ {
		newTagVariation := db.TagVariation{
			TagID:              tagVariations[i].TagID,
			VariationVersionID: dstID,
		}

		tx.Create(&newTagVariation)
	}

	songDatabaseVariations := []db.SongDatabaseVariation{}
	tx.Where("variation_version_id = ?", srcID).Find(&songDatabaseVariations)

	for i := 0; i < len(songDatabaseVariations); i++ {
		newSongDatabaseVariation := db.SongDatabaseVariation{
			SongDatabaseID:     songDatabaseVariations[i].SongDatabaseID,
			VariationVersionID: dstID,
		}
		tx.Create(&newSongDatabaseVariation)
	}

	scheduleVariations := []db.ScheduleVariation{}
	tx.Where("variation_version_id = ?", srcID).Find(&scheduleVariations)

	for i := 0; i < len(scheduleVariations); i++ {
		newScheduleVariation := db.ScheduleVariation{
			ScheduleID:         scheduleVariations[i].ScheduleID,
			VariationVersionID: dstID,
			OrderNumber:        scheduleVariations[i].OrderNumber,
		}
		tx.Create(&newScheduleVariation)
	}
}
