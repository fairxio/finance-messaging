package iso20022

// General information about the corporate action event.
type CorporateActionGeneralInformation100 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType3Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType42Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes70 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation100) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation100) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation100) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation100) AddEventProcessingType() *CorporateActionEventProcessingType3Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType3Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation100) AddEventType() *CorporateActionEventType42Choice {
	c.EventType = new(CorporateActionEventType42Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation100) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation100) AddUnderlyingSecurity() *FinancialInstrumentAttributes70 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes70)
	return c.UnderlyingSecurity
}
