package contactBook

import (
	"fmt"
	"sync"
)

type ContactManager struct {
	list *ListOfContacts
	mu   sync.Mutex
}

func NewContactManager() *ContactManager {
	return &ContactManager{
		list: &ListOfContacts{
			Contacts: map[int]Contact{},
		},
	}
}

func (cm *ContactManager) AddContact(contact Contact) {
	if contact.isValid() {
		cm.mu.Lock()
		defer cm.mu.Unlock()
		id := cm.list.NewId()
		cm.list.Contacts[id] = contact
	} else {
		fmt.Println("Wrong contact input")
	}
}

func (cm *ContactManager) RemoveContact(id int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.list.Contacts, id)
}

func (cm *ContactManager) UpdateContact(obj UpdateContactDTO) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	contact, ok := cm.list.Contacts[obj.Id]
	if ok {
		if obj.Name.isValid() {
			contact.Name = obj.Name
		}
		if obj.PhoneNumber.isValid() {
			contact.PhoneNumber = obj.PhoneNumber
		}
		cm.list.Contacts[obj.Id] = contact
	}
}

func (cm *ContactManager) GetAllContacts() Contacts {
	return cm.list.Clone()
}
