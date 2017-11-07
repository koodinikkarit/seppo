package db

import (
	"github.com/btfak/sqlext"
	"github.com/jinzhu/gorm"
)

func MoveVariationReferences(
	tx *gorm.DB,
	srcVariationID uint32,
	dstVariationID uint32,
) {
	tagVariations := []TagVariation{}
	tx.Where("variation_id = ?", srcVariationID).Find(&tagVariations)
	var newTagVariations []TagVariation
	for i := 0; i < len(tagVariations); i++ {
		newTagVariations = append(
			newTagVariations,
			TagVariation{
				TagID:       tagVariations[i].TagID,
				VariationID: dstVariationID,
			},
		)
	}
	sqlext.BatchInsert(tx.DB(), newTagVariations)

	songDatabaseVariations := []SongDatabaseVariation{}
	tx.Where("variation_version_id = ?", srcVariationID).Find(&songDatabaseVariations)
	var newSongDatabaseVariations []SongDatabaseVariation
	for i := 0; i < len(songDatabaseVariations); i++ {
		newSongDatabaseVariations = append(
			newSongDatabaseVariations,
			SongDatabaseVariation{
				SongDatabaseID: songDatabaseVariations[i].SongDatabaseID,
				VariationID:    dstVariationID,
			},
		)
	}
	sqlext.BatchInsert(tx.DB(), newSongDatabaseVariations)

	scheduleVariations := []ScheduleVariation{}
	tx.Where("variation_version_id = ?", srcVariationID).Find(&scheduleVariations)
	var newScheduleVariations []ScheduleVariation
	for i := 0; i < len(scheduleVariations); i++ {
		newScheduleVariations = append(
			newScheduleVariations,
			ScheduleVariation{
				ScheduleID:  scheduleVariations[i].ScheduleID,
				VariationID: dstVariationID,
				OrderNumber: scheduleVariations[i].OrderNumber,
			},
		)
	}
	sqlext.BatchInsert(tx.DB(), newScheduleVariations)
}
