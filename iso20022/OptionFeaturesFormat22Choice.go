package iso20022

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat22Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures10Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat22Choice) SetCode(value string) {
	o.Code = (*OptionFeatures10Code)(&value)
}

func (o *OptionFeaturesFormat22Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}
