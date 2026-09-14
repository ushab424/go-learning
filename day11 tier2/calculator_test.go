package main

import "testing"

func TestSubtractTable(t *testing.T) {
	tests := []struct {
		name      string
		a, b, res int
	}{
		{"zero", 0, 0, 0},
		{"big-small", 2, 1, 1},
		{"small-big", 1, 2, -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// t.Run для простой ориентации внутри кейсов теста
			result := Subtract(test.a, test.b)
			if result != test.res {
				t.Errorf("got %d, want %d", result, test.res)
			}
		})
	}
}

func TestDivideTable(t *testing.T) {
	tests := []struct {
		name string
		// имя кейса
		a, b, res int
		// значения кейса
		wantErr bool
		// ожидаемая ошибка
	}{
		{"ZeroValues", 0, 0, 0, true},
		{"Negative", -2, -2, 1, false},
		{"DivideByZero", 2, 0, 0, true},
		{"Normal", 4, 2, 2, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Divide(test.a, test.b)
			// учитываем ошибку в функции
			if test.wantErr {
				// учитываем ошибку
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != test.res {
				t.Errorf("got %d, want %d", result, test.res)
			}
			// обрабатываем кадый вариант развития событий
		})
	}
}
