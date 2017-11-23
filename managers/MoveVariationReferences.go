package managers

import (
	"database/sql"

	"github.com/koodinikkarit/seppo/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func MoveVariationReferences(
	tx *sql.Tx,
	srcVariationID uint64,
	dstVariationID uint64,
) {
	tagVariations, _ := models.TagVariations(
		tx,
		qm.Where("variation_id = ?", srcVariationID),
	).All()
	for _, tagVariation := range tagVariations {
		newTagVariation := models.TagVariation{
			TagID:       tagVariation.TagID,
			VariationID: dstVariationID,
		}
		newTagVariation.Insert(tx)
	}
	songDatabaseVariations, _ := models.SongDatabaseVariations(
		tx,
		qm.Where("variation_id = ?", srcVariationID),
	).All()
	for _, songDatabaseVariation := range songDatabaseVariations {
		newSongDatabaseVariation := models.SongDatabaseVariation{
			SongDatabaseID: songDatabaseVariation.SongDatabaseID,
			VariationID:    dstVariationID,
		}
		newSongDatabaseVariation.Insert(tx)
	}
	scheduleVariations, _ := models.ScheduleVariations(
		tx,
		qm.Where("variation_version_id = ?", srcVariationID),
	).All()
	for _, scheduleVariation := range scheduleVariations {
		newScheduleVariation := models.ScheduleVariation{
			ScheduleID:  scheduleVariation.ScheduleID,
			VariationID: dstVariationID,
			OrderNumber: scheduleVariation.OrderNumber,
		}
		newScheduleVariation.Insert(tx)
	}
}
