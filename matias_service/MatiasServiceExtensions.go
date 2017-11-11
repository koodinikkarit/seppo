package MatiasService

func (r *SyncEwDatabaseRequest) HasEwSong(
	ewSongID uint32,
) bool {
	for i := 0; i < len(r.EwSongs); i++ {
		if r.EwSongs[i].Id == ewSongID {
			return true
		}
	}
	return false
}

func (r *SyncEwDatabaseRequest) HasEwSongWithNameAndText(
	name string,
	text string,
) bool {
	for _, ewSong := range r.EwSongs {
		if ewSong.Title == name &&
			ewSong.Text == text {
			return true
		}
	}
	return false
}

func (r *SyncEwDatabaseResponse) CreateOrGetEwSong(
	id uint32,
) *EwSong {
	for _, ewSong := range r.EwSongs {
		if ewSong.Id == id {
			return ewSong
		}
	}
	newEwSong := &EwSong{
		Id: id,
	}
	r.EwSongs = append(
		r.EwSongs,
		newEwSong,
	)
	return newEwSong
}
