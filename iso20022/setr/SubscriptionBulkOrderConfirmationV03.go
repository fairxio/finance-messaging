package setr

import (
	"encoding/xml"

	"github.com/fairxio/finance-messaging/iso20022"
)

type Document00900103 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.009.001.03 Document"`
	Message *SubscriptionBulkOrderConfirmationV03 `xml:"SbcptBlkOrdrConfV03"`
}

func (d *Document00900103) AddMessage() *SubscriptionBulkOrderConfirmationV03 {
	d.Message = new(SubscriptionBulkOrderConfirmationV03)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent sends the SubscriptionBulkOrderConfirmation message to the instructing party, for example, an investment manager or its authorised representative to confirm the details of the execution of a SubscriptionBulkOrder instruction.
// Usage
// The SubscriptionBulkOrderConfirmation message is used to confirm the execution of all individual orders.
// There is usually one bulk confirmation message for one bulk order message.
// Each individual order confirmation specified is identified in DealReference. The reference of the original individual order is specified in OrderReference. The message identification of the SubscriptionBulkOrder message in which the individual order was conveyed may also be quoted in RelatedReference.
// A SubscriptionBulkOrder must in all cases be responded to by a SubscriptionBulkOrderConfirmation and in no circumstances by a SubscriptionOrderConfirmation.
// If the executing party needs to confirm a SubscriptionOrder instruction, then the SubscriptionOrderConfirmation must be used.
type SubscriptionBulkOrderConfirmationV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// General information related to the execution of investment orders.
	BulkExecutionDetails *iso20022.SubscriptionBulkExecution3 `xml:"BlkExctnDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*iso20022.Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderConfirmationV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderConfirmationV03) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderConfirmationV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationV03) AddRelatedReference() *iso20022.AdditionalReference3 {
	s.RelatedReference = new(iso20022.AdditionalReference3)
	return s.RelatedReference
}

func (s *SubscriptionBulkOrderConfirmationV03) AddBulkExecutionDetails() *iso20022.SubscriptionBulkExecution3 {
	s.BulkExecutionDetails = new(iso20022.SubscriptionBulkExecution3)
	return s.BulkExecutionDetails
}

func (s *SubscriptionBulkOrderConfirmationV03) AddRelatedPartyDetails() *iso20022.Intermediary9 {
	newValue := new(iso20022.Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationV03) AddCopyDetails() *iso20022.CopyInformation2 {
	s.CopyDetails = new(iso20022.CopyInformation2)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderConfirmationV03) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
