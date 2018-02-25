package iso20022

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type AggregateBalanceInformation3 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateQuantity *BalanceQuantity1Choice `xml:"AggtQty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal"`

	// Previous total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	PreviousHoldingValue *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsHldgVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Indicates whether the accrued interest is a positive or negative amount.
	AccruedInterestAmountSign *PlusOrMinusIndicator `xml:"AcrdIntrstAmtSgn,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Security held on the account for which the balance is calculated.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation2 `xml:"PricDtls"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation2 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace4 `xml:"BalAtSfkpgPlc,omitempty"`
}

func (a *AggregateBalanceInformation3) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalanceInformation3) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation3) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalanceInformation3) SetPreviousHoldingValue(value, currency string) {
	a.PreviousHoldingValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation3) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation3) SetAccruedInterestAmountSign(value string) {
	a.AccruedInterestAmountSign = (*PlusOrMinusIndicator)(&value)
}

func (a *AggregateBalanceInformation3) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation3) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation3) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	a.FinancialInstrumentDetails = new(FinancialInstrument13)
	return a.FinancialInstrumentDetails
}

func (a *AggregateBalanceInformation3) AddPriceDetails() *PriceInformation2 {
	newValue := new(PriceInformation2)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation3) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalanceInformation3) AddBalanceBreakdownDetails() *SubBalanceInformation2 {
	newValue := new(SubBalanceInformation2)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation3) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation3) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace4 {
	newValue := new(AggregateBalancePerSafekeepingPlace4)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}
