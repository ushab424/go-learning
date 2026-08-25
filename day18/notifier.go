package main

import "fmt"

type Notifier interface {
	Notify(message string) string
}

type Email struct {
	Adress string
}

type SMS struct {
	Phone string
}

type Push struct {
	DeviceID string
}

func (e *Email) Notify(message string) string {
	return fmt.Sprintf("Email to %s: %s", e.Adress, message)
}

func (s *SMS) Notify(message string) string {
	return fmt.Sprintf("SMS to %s: %s", s.Phone, message)
}

func (p *Push) Notify(message string) string {
	return fmt.Sprintf("Push to %s: %s", p.DeviceID, message)
}

func SendAll(notifiers []Notifier, message string) {
	for _, messages := range notifiers {
		fmt.Println(messages.Notify(message))
	}
}

func main() {
	notifiers := []Notifier{
		&Email{Adress: "ivan@gjkl.com"},
		&SMS{Phone: "8946354562"},
		&Push{DeviceID: "842947935"},
	}
	SendAll(notifiers, "Hello")
}
