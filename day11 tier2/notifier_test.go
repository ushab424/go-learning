package main

import "testing"

type MockSender struct {
	Called bool
	To     string
	Msg    string
}

func (m *MockSender) Send(to, msg string) error {
	m.Called = true
	m.To = to
	m.Msg = msg
	return nil
}

func TestNotify(t *testing.T) {
	mock := &MockSender{}
	notifier := &EmailNotifier{sender: mock}
	notifier.Notify("test@mail.com", "hello")

	if !mock.Called {
		t.Error("Send was not called")
	}
	if mock.To != "test@mail.com" {
		t.Errorf("got to=%s, want test@mail.com", mock.To)
	}
}

// Мок — это подмена реальной зависимости на фейковую в тестах.
// Например функция отправляет email — в тестах не хочешь отправлять реальные письма.
// Подменяешь отправщик на мок который просто запоминает что его вызвали.

// MockSender реализует интерфейс Sender но вместо отправки письма просто запоминает аргументы.
// В тесте проверяем что Notify вызвал Send с правильными данными
