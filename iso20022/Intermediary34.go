package iso20022

// Party that provides services to investors relating to financial products.
type Intermediary34 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account20 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (i *Intermediary34) AddIdentification() *PartyIdentification70Choice {
	i.Identification = new(PartyIdentification70Choice)
	return i.Identification
}

func (i *Intermediary34) AddAccount() *Account20 {
	i.Account = new(Account20)
	return i.Account
}

func (i *Intermediary34) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}

func (i *Intermediary34) AddContactPerson() *ContactIdentification2 {
	i.ContactPerson = new(ContactIdentification2)
	return i.ContactPerson
}
