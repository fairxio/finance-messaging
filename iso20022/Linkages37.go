package iso20022

// Information related to a linked transaction.
type Linkages37 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition7Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References41Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity3Choice `xml:"LkdQty,omitempty"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification92Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages37) AddProcessingPosition() *ProcessingPosition7Choice {
	l.ProcessingPosition = new(ProcessingPosition7Choice)
	return l.ProcessingPosition
}

func (l *Linkages37) AddMessageNumber() *DocumentNumber5Choice {
	l.MessageNumber = new(DocumentNumber5Choice)
	return l.MessageNumber
}

func (l *Linkages37) AddReference() *References41Choice {
	l.Reference = new(References41Choice)
	return l.Reference
}

func (l *Linkages37) AddLinkedQuantity() *PairedOrTurnedQuantity3Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity3Choice)
	return l.LinkedQuantity
}

func (l *Linkages37) AddReferenceOwner() *PartyIdentification92Choice {
	l.ReferenceOwner = new(PartyIdentification92Choice)
	return l.ReferenceOwner
}
