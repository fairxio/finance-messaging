package iso20022

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus13 struct {

	// Value expressed as a rate type.
	RateType *RateType20Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus13) AddRateType() *RateType20Choice {
	r.RateType = new(RateType20Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus13) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus13) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}
