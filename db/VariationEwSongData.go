package SeppoDB

type VariationEwSongData struct {
	VariationID   uint32
	Author        string
	Copyright     string
	Administrator string
	Description   string
	Tags          string

	Variation Variation
}
