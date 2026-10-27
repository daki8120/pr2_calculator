package main

import "fmt"

// typesDemo демонстрирует все базовые типы Go
func typesDemo() {
	// ── Целые числа ─────────────────────────────────────
	var a int = 42
	var b int8 = 127
	var c int16 = 32767
	var d int32 = 2147483647
	var e int64 = 9223372036854775807
	var f uint = 42

	fmt.Printf("int: %d, int8: %d, int16: %d\n", a, b, c)
	fmt.Printf("int32: %d, int64: %d, uint: %d\n", d, e, f)

	// ── Вещественные числа ──────────────────────────────
	var g float32 = 3.14
	var h float64 = 3.141592653589793

	fmt.Printf("float32: %f, float64: %.15f\n", g, h)

	// ── Строки ──────────────────────────────────────────
	var s1 string = "Hello, Go!"
	s2 := `Многострочная
строка`

	fmt.Printf("s1: %s\n", s1)
	fmt.Printf("s2: %s\n", s2)
	fmt.Printf("Длина s1: %d байт\n", len(s1))

	// ── Символы ──────────────────────────────────────────
	var ch byte = 'A'
	var ru rune = 'Я'

	fmt.Printf("byte: %c (%d), rune: %c (%d)\n", ch, ch, ru, ru)

	// ── Логический тип ──────────────────────────────────
	var flag bool = true
	isPositive := 42 > 0

	fmt.Printf("bool: %t, isPositive: %t\n", flag, isPositive)

	// ── Срезы (slices) ──────────────────────────────────
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7)

	fmt.Printf("Срез: %v, длина: %d\n", numbers, len(numbers))

	sub := numbers[1:4]
	fmt.Printf("Подсрез [1:4]: %v\n", sub)

	// ── Карты (maps) ─────────────────────────────────────
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}

	scores["Dave"] = 78

	val, ok := scores["Alice"]
	if ok {
		fmt.Printf("Alice набрала: %d баллов\n", val)
	}

	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}
}
