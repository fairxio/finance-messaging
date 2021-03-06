package seev

import (
	"encoding/xml"

	"github.com/fairxio/finance-messaging/iso20022"
)

type Document04000207 struct {
	XMLName xml.Name                                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.002.07 Document"`
	Message *CorporateActionInstructionCancellationRequest002V07 `xml:"CorpActnInstrCxlReq"`
}

func (d *Document04000207) AddMessage() *CorporateActionInstructionCancellationRequest002V07 {
	d.Message = new(CorporateActionInstructionCancellationRequest002V07)
	return d.Message
}

// Scope
// An account owner sends the CorporateActionInstructionCancellationRequest message to an account servicer to request cancellation of a previously sent corporate action election instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionCancellationRequest002V07 struct {

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *iso20022.YesNoIndicator `xml:"ChngInstrInd,omitempty"`

	// Identification of a previously sent instruction document.
	InstructionIdentification *iso20022.DocumentIdentification37 `xml:"InstrId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation121 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountDetails *iso20022.AccountIdentification42 `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionInstruction *iso20022.CorporateActionOption128 `xml:"CorpActnInstr"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionCancellationRequest002V07) SetChangeInstructionIndicator(value string) {
	c.ChangeInstructionIndicator = (*iso20022.YesNoIndicator)(&value)
}

func (c *CorporateActionInstructionCancellationRequest002V07) AddInstructionIdentification() *iso20022.DocumentIdentification37 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification37)
	return c.InstructionIdentification
}

func (c *CorporateActionInstructionCancellationRequest002V07) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation121 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation121)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionCancellationRequest002V07) AddAccountDetails() *iso20022.AccountIdentification42 {
	c.AccountDetails = new(iso20022.AccountIdentification42)
	return c.AccountDetails
}

func (c *CorporateActionInstructionCancellationRequest002V07) AddCorporateActionInstruction() *iso20022.CorporateActionOption128 {
	c.CorporateActionInstruction = new(iso20022.CorporateActionOption128)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionCancellationRequest002V07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
