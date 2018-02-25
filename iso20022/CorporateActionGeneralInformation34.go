package iso20022

// General information about the corporate action event.
type CorporateActionGeneralInformation34 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation34) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation34) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation34) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation34) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}
