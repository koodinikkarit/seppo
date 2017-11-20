package managers

import (
	"github.com/koodinikkarit/seppo/matias_service"
)

func RemoveDuplicatesFromEwSongs(
	ewSongs []*MatiasService.EwSong,
) (
	map[uint32]*MatiasService.EwSong,
	[]uint32,
) {
	outputEwSongs := make(map[uint32]*MatiasService.EwSong)
	var removeEwSongs []uint32

	for _, ewSong := range ewSongs {
		foundSame := false
		for _, ewSong2 := range outputEwSongs {
			if ewSong.Title == ewSong2.Title &&
				ewSong.Text == ewSong2.Text {
				foundSame = true
				removeEwSongs = append(
					removeEwSongs,
					ewSong2.Id,
				)
			}
		}
		if foundSame == false {
			outputEwSongs[ewSong.Id] = ewSong
		}
	}
	return outputEwSongs, removeEwSongs
}

func HasEwSongByNameAndText(
	ewSongs map[uint32]*MatiasService.EwSong,
	name string,
	text string,
) bool {
	for _, ewSong := range ewSongs {
		if ewSong.Title == name &&
			ewSong.Text == text {
			return true
		}
	}
	return false
}
