package contactBook

type Contacts struct {
	ContactsList []Contact `json:"contacts_list"`
}

type Contact struct {
	Id          string      `json:"id"`
	Name        Name        `json:"name"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

type Name string

func (n Name) isValid() bool {
	if len(n) == 0 {
		return false
	}
	return true
}

type PhoneNumber string

func (pn PhoneNumber) isValid() bool {
	if len(pn) == 0 {
		return false
	}
	for _, c := range pn {
		if c < '0' || c > '9' {
			return false
		}
	}
	if len(pn) != 11 {
		return false
	}

	return true
}

func (c *Contact) isValid() bool {
	return c.Name.isValid() && c.PhoneNumber.isValid()
}
