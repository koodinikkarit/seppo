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

	Conflicts                     []SrEwConflict                  `gorm:"ForeignKey:SrID"`
	NewAuthors                    []SrNewAuthors                  `gorm:"ForeignKey:SrID"`
	NewCopyrights                 []SrcNewCopyright               `gorm:"ForeignKey:SrID"`
	NewVariations                 []SrNewVariation                `gorm:"ForeignKey:SrID"`
	NewEwDatabaseLinks            []SrEwDatabaseLink              `gorm:"ForeignKey:SrID"`
	RemovedSongDatabaseVariations []SrRemoveSongDatabaseVariation `gorm:"ForeignKey:SrID"`
	AddedSongDatabaseVariations   []SrAddSongDatabaseVariation    `gorm:"ForeignKey:SrID"`
	EwSongs                       []SrEwSong                      `gorm:"ForeignKey:SrID"`
	NewBranches                   []SrNewBranch                   `gorm:"ForeignKey:SrID"`
	NewVariationVersions          []SrNewVariationVersion         `gorm:"ForeignKey:SrID"`
	PassivatedVariationVersions   []SrPassivatedVariationVersion  `gorm:"ForeignKey:SrID"`
}
