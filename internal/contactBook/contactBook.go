package contactBook

import (
	"ContactBook/internal/db"
	"github.com/google/uuid"
	"log"
)

type ContactManager struct {
	db     db.Database
	userId string
}

func NewContactManager(db db.Database, id string) *ContactManager {
	return &ContactManager{
		db:     db,
		userId: id,
	}
}

func (m *ContactManager) GetAllContacts() (Contacts, error) {
	contacts := Contacts{}

	log.Println("user id>>", m.userId)
	rows, err := m.db.Conn.Query(
		`SELECT id, name, phone_number FROM contacts JOIN user_contacts 
    			ON contacts.id = user_contacts.contact_id
				WHERE user_contacts.user_id=$1 ORDER BY name;`, m.userId)
	if err != nil {
		return contacts, err
	}
	defer rows.Close()
	for rows.Next() {
		var contact Contact
		err := rows.Scan(&contact.Id, &contact.Name, &contact.PhoneNumber)
		if err != nil {
			return contacts, err
		}
		contacts.ContactsList = append(contacts.ContactsList, contact)
	}
	log.Println("contacts>>>", contacts)
	return contacts, nil
}

func (m *ContactManager) AddContact(contact Contact) error {
	if contact.isValid() {
		uuidNew := uuid.New()
		log.Println("add contact (name, phone) >>>", contact)
		_, err := m.db.Conn.Exec(
			"INSERT INTO contacts (id, name, phone_number) VALUES ($1, $2, $3)",
			uuidNew, contact.Name, contact.PhoneNumber)
		if err != nil {
			return err
		}
		_, err = m.db.Conn.Exec("INSERT INTO user_contacts (user_id, contact_id) VALUES ($1, $2)",
			m.userId, uuidNew)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *ContactManager) RemoveContact(contactId string) error {
	log.Println("RemoveContact contactId>>>", contactId)
	_, err := m.db.Conn.Exec("DELETE FROM contacts WHERE id=$1", contactId)
	if err != nil {
		return err
	}
	return nil
}

func (m *ContactManager) UpdateContact(contact Contact) error {
	log.Println("UpdateContact >>>", contact)
	log.Println("UpdateContact Name>>>", contact.Name)
	log.Println("UpdateContact PhoneNumber>>>", contact.PhoneNumber)
	if contact.isValid() {
		_, err := m.db.Conn.Exec("UPDATE contacts SET name=$1, phone_number=$2 WHERE id=$3",
			contact.Name, contact.PhoneNumber, contact.Id)
		if err != nil {
			return err
		}
		return nil
	} else if contact.Name.isValid() && !contact.PhoneNumber.isValid() {
		_, err := m.db.Conn.Exec("UPDATE contacts SET name=$1 WHERE id=$2",
			contact.Name, contact.Id)
		if err != nil {
			return err
		}
		return nil
	} else if !contact.Name.isValid() && contact.PhoneNumber.isValid() {
		_, err := m.db.Conn.Exec("UPDATE contacts SET phone_number=$1 WHERE id=$2",
			contact.PhoneNumber, contact.Id)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
