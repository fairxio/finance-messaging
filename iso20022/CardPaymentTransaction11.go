package iso20022

// Data associated with the transaction during the authorisation.
type CardPaymentTransaction11 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType4Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails1 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction11) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction11) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType4Code)(&value)
}

func (c *CardPaymentTransaction11) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction11) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction11) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction11) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction11) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction11) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction11) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction11) AddTransactionDetails() *CardPaymentTransactionDetails1 {
	c.TransactionDetails = new(CardPaymentTransactionDetails1)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction11) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}
