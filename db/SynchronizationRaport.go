package db

import "time"

type SynchronizationRaport struct {
	ID            uint32
	RaportType    uint32
	DatabaseID    *uint32
	DatabaseKey   string
	DatabaseFound bool
	DurationMS    int64
	StartedAt     *time.Time
	FinishedAt    *time.Time

	PassivatedVariationVersions []VariationVersion `gorm:"many2many:synchronization_raport_variation_version_passivations"`
	NewVariationVersions        []VariationVersion `gorm:"many2many:synchronization_raport_new_variation_versions"`
	NewBranches                 []Branch           `gorm:"many2many:synchronization_raport_new_branches"`
	RemovedEwSongs              []EwDatabaseLink   `gorm:"many2many:synchronization_raport_remove_ew_song"`
	AddedVariations             []Variation        `gorm:"many2many:synchronization_raport_add_song_database_variations"`
	RemovedVariations           []Variation        `gorm:"many2many:synchronization_raport_remove_song_database_variations"`
	AddEwSongs                  []EwDatabaseLink   `gorm:"many2many:synchronization_raport_add_ew_songs"`
}
