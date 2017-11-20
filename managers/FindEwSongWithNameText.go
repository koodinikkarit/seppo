package managers

import (
	"github.com/koodinikkarit/seppo/matias_service"
)

func FindEwSongWithNameText(
	ewSongs map[uint32]*MatiasService.EwSong,
	name string,
	text string,
) *MatiasService.EwSong {
	for _, ewSong := range ewSongs {
		if ewSong.Title == name &&
			ewSong.Text == text {
			return ewSong
		}
	}
	return nil
}
