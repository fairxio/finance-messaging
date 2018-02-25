package iso20022

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection29 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms18 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection29) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection29) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection29) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection29) AddForeignExchangeDetails() *ForeignExchangeTerms18 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms18)
	return a.ForeignExchangeDetails
}