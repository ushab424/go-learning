package main

// тест должен быть в томже пакете что и функция

import "testing"

// стандарный пакет тестирования

func TestAdd(t *testing.T) {
	// имя начинается с Test+имя функции
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d, want 5", result)
		// тест провален но продолжает выполняться
		// t.Fatalf(...) - провален и останавливается сразу
	}
}

// запуск go test -v
func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Fatalf("Multiply(2,3) = %d, want 6", result)
	}
}

func TestAddTable(t *testing.T) {
	// табличный тест
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{100, 200, 300},
		{0, 0, 0},
	}
	// слайс анонимных структур. Каждая структура — один тест-кейс с входными данными и ожидаемым результатом.
	for _, tt := range tests {
		// проходим по всем кейсам
		result := Add(tt.a, tt.b)
		if result != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.want)
		}
	}
}
func TestMultiplyTable(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 2},
		{0, 0, 0},
		{-1, 1, -1},
		{100, 200, 20000},
	}
	for _, tt := range tests {
		result := Multiply(tt.a, tt.b)
		if result != tt.want {
			t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.want)
		}
	}
}

// unit-tests: TestAdd проверяет только функцию Add, больше ничего. Быстрые, простые, их пишут больше всего.

// какие тесты писать? - думай о граничных случаях. Для Add: нули, отрицательные, большие числа, оба отрицательные.
//       Для строковой функции: пустая строка, один символ, юникод. Для слайса: пустой, один элемент, nil. Правило — "что может пойти не так?"

// Как понять что тестировать? - тестируй публичные функции (с большой буквы). Каждая функция которую будет вызывать другой код должна иметь тест.
//       Не тестируй приватные вспомогательные функции отдельно — они тестируются через публичные.

func TestAddSubtests(t *testing.T) {
	tests := []struct {
		name       string
		a, b, want int
	}{
		{"positive", 1, 2, 3},
		{"zeros", 0, 0, 0},
		{"negative", -1, 1, 0},
		// в кейсы добавляем имена каждых кейсов
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// функия t.Run - создает подтест с именем
			result := Add(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}

/*
Зачем t.Run и подтесты:
Без подтестов если тест упал — видишь только "TestAddTable FAIL" и ищешь какой именно кейс сломался.
 С подтестами видишь точно: "TestAddSubtests/negative FAIL" — сразу понятно что сломалось на отрицательных числах.
Можно запустить один конкретный кейс: go test -run TestAddSubtests/negative. Не нужно гонять все 100 кейсов когда чинишь один.
Когда юзать — всегда в табличных тестах. Это стандарт в Go-проектах. Без t.Run табличные тесты тоже работают, но отладка сложнее.
*/

func BenchmarkAdd(b *testing.B) {
	// Benchmark вместо Test, *testing.B вместо *testing.T. b.N — Go сам подберёт количество итераций для точного замера.
	// Запуск: go test -bench=.. Точка значит "все бенчмарки".
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// Когда нужны бенчмарки — когда оптимизируешь код. Написал две версии функции, прогнал бенчмарк, сравнил ns/op.
// Меньше — быстрее. Также полезно когда выбираешь между слайсом и мапой, или между строковой конкатенацией и strings.Builder.
