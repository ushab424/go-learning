package main

type Sender interface {
	Send(to, msg string) error
}

type EmailNotifier struct {
	sender Sender
}

func (n *EmailNotifier) Notify(to, msg string) error {
	return n.sender.Send(to, msg)
}
