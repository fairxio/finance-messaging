package iso20022

// Set of elements used to provide information on the corrective payment initiation transaction, to which the resolution message refers.
type CorrectivePaymentInitiation2 struct {

	// Set of elements used to provide corrective information for the group header of the message under investigation.
	GroupHeader *CorrectiveGroupInformation1 `xml:"GrpHdr,omitempty"`

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Date or date time at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date or date time on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt,omitempty"`

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`
}

func (c *CorrectivePaymentInitiation2) AddGroupHeader() *CorrectiveGroupInformation1 {
	c.GroupHeader = new(CorrectiveGroupInformation1)
	return c.GroupHeader
}

func (c *CorrectivePaymentInitiation2) SetPaymentInformationIdentification(value string) {
	c.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (c *CorrectivePaymentInitiation2) SetInstructionIdentification(value string) {
	c.InstructionIdentification = (*Max35Text)(&value)
}

func (c *CorrectivePaymentInitiation2) SetEndToEndIdentification(value string) {
	c.EndToEndIdentification = (*Max35Text)(&value)
}

func (c *CorrectivePaymentInitiation2) SetInstructedAmount(value, currency string) {
	c.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CorrectivePaymentInitiation2) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	c.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return c.RequestedExecutionDate
}

func (c *CorrectivePaymentInitiation2) SetRequestedCollectionDate(value string) {
	c.RequestedCollectionDate = (*ISODate)(&value)
}
