package contactBook

import "net/http"

type Contacts map[int]Contact

type ListOfContacts struct {
	Contacts Contacts `json:"contacts"`
	IdList   []int    `json:"-"`
}

//type ContactId int

func (l *ListOfContacts) NewId() int {
	if len(l.IdList) == 0 {
		l.IdList = append(l.IdList, 1)
		return 1
	} else {
		newId := l.IdList[len(l.IdList)-1] + 1
		l.IdList = append(l.IdList, newId)
		return newId
	}
}

type Contact struct {
	Name        Name        `json:"name"`
	PhoneNumber PhoneNumber `json:"phoneNumber"`
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

func (l *ListOfContacts) Clone() Contacts {
	clone := make(map[int]Contact)
	for key, val := range l.Contacts {
		clone[key] = val
	}
	return clone
}

type UpdateContactDTO struct {
	Id          int         `json:"id"`
	Name        Name        `json:"name"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

// ??????????????
func (*ListOfContacts) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Contact) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
