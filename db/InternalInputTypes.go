package SeppoDB

type createVariationInternalInput struct {
	input         CreateVariationInput
	returnChannel chan *Variation
}

type editVariationInternalInput struct {
	input         EditVariationInput
	returnChannel chan *Variation
}
