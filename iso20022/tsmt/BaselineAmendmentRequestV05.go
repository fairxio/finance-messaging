package tsmt

import (
	"encoding/xml"

	"github.com/fairxio/finance-messaging/iso20022"
)

type Document00900105 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.05 Document"`
	Message *BaselineAmendmentRequestV05 `xml:"BaselnAmdmntReq"`
}

func (d *Document00900105) AddMessage() *BaselineAmendmentRequestV05 {
	d.Message = new(BaselineAmendmentRequestV05)
	return d.Message
}

// Scope
// The BaselineAmendmentRequest message is sent by a primary party involved in a transaction to the matching application.
// The message is used to request the amendment of an established baseline.
// Usage
// The BaselineAmendmentRequest message may only be sent if the transaction is in the state Established or Active.
// The BaselineAmendmentRequest message can be sent to the matching application by one of the primary parties involved in a transaction established in the push-through mode to request the amendment of an established baseline.
// The matching application acknowledges the receipt of the amendment request by sending a DeltaReport message to the submitter of the BaselineAmendmentRequest message. It passes on the newly proposed baseline to the counterparty by sending a FullPushThroughReport message, a DeltaReport message and a pre-calculated BaselineReport message.
// The counterparty is expected to either accept or reject the amendment request by submitting an AmendmentAcceptance or AmendmentRejection message.
// or
// The BaselineAmendmentRequest message can be sent by the party involved in a transaction established in the lodge mode to the matching application to amend an established baseline.
// The matching application amends the baseline according to the BaselineAmendmentRequest message and confirms the execution of the request by sending a DeltaReport and calculated BaselineReport message to the requester of the amendment.
type BaselineAmendmentRequestV05 struct {

	// Identifies the request message.
	RequestIdentification *iso20022.MessageIdentification1 `xml:"ReqId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Specifies the commercial details of the underlying transaction.
	Baseline *iso20022.Baseline5 `xml:"Baseln"`

	// Person to be contacted in the organisation of the buyer.
	BuyerContactPerson []*iso20022.ContactIdentification1 `xml:"BuyrCtctPrsn,omitempty"`

	// Person to be contacted in the organisation of the seller.
	SellerContactPerson []*iso20022.ContactIdentification1 `xml:"SellrCtctPrsn,omitempty"`

	// Person to be contacted in the buyer's bank.
	BuyerBankContactPerson []*iso20022.ContactIdentification1 `xml:"BuyrBkCtctPrsn,omitempty"`

	// Person to be contacted in the seller's bank.
	SellerBankContactPerson []*iso20022.ContactIdentification1 `xml:"SellrBkCtctPrsn,omitempty"`

	// Person to be contacted in another bank than the seller or buyer's bank.
	OtherBankContactPerson []*iso20022.ContactIdentification3 `xml:"OthrBkCtctPrsn,omitempty"`
}

func (b *BaselineAmendmentRequestV05) AddRequestIdentification() *iso20022.MessageIdentification1 {
	b.RequestIdentification = new(iso20022.MessageIdentification1)
	return b.RequestIdentification
}

func (b *BaselineAmendmentRequestV05) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	b.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return b.TransactionIdentification
}

func (b *BaselineAmendmentRequestV05) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	b.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return b.SubmitterTransactionReference
}

func (b *BaselineAmendmentRequestV05) AddBaseline() *iso20022.Baseline5 {
	b.Baseline = new(iso20022.Baseline5)
	return b.Baseline
}

func (b *BaselineAmendmentRequestV05) AddBuyerContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	b.BuyerContactPerson = append(b.BuyerContactPerson, newValue)
	return newValue
}

func (b *BaselineAmendmentRequestV05) AddSellerContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	b.SellerContactPerson = append(b.SellerContactPerson, newValue)
	return newValue
}

func (b *BaselineAmendmentRequestV05) AddBuyerBankContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	b.BuyerBankContactPerson = append(b.BuyerBankContactPerson, newValue)
	return newValue
}

func (b *BaselineAmendmentRequestV05) AddSellerBankContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	b.SellerBankContactPerson = append(b.SellerBankContactPerson, newValue)
	return newValue
}

func (b *BaselineAmendmentRequestV05) AddOtherBankContactPerson() *iso20022.ContactIdentification3 {
	newValue := new(iso20022.ContactIdentification3)
	b.OtherBankContactPerson = append(b.OtherBankContactPerson, newValue)
	return newValue
}