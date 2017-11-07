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
