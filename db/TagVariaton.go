package SeppoDB

type TagVariation struct {
	ID          uint32
	TagID       uint32
	VariationID uint32

	Tag       Tag
	Variation Variation
}
