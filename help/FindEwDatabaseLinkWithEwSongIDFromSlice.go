package help

import (
	"github.com/koodinikkarit/seppo/db"
)

func FindEwDatabaseLinkWithEwSongIDFromSlice(
	ewDatabaseLinks []db.EwDatabaseLink,
	ewSongID uint32,
) *db.EwDatabaseLink {
	var foundEwDatabaseLink *db.EwDatabaseLink
	for _, ewDatabaseLink := range ewDatabaseLinks {
		if ewDatabaseLink.EwDatabaseSongID == ewSongID {
			foundEwDatabaseLink = &ewDatabaseLink
		}
	}
	return foundEwDatabaseLink
}
