package SeppoDB

type SongDatabaseVariation struct {
	ID             uint32
	SongDatabaseID uint32
	VariationID    uint32

	SongDatabase SongDatabase
	Variation    Variation
}
