// Пакет main — точка входа программы-калькулятора
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ====================================================
// КОНСТАНТЫ
// ====================================================

const HistorySize = 10

// ====================================================
// СТРУКТУРА ДЛЯ ЗАПИСИ ИСТОРИИ
// ====================================================

type CalcRecord struct {
	expression string
	result     float64
}

// ====================================================
// ГЛОБАЛЬНАЯ ПЕРЕМЕННАЯ — история вычислений
// ====================================================

var history []CalcRecord

// unaryOps — операции, которым достаточно одного аргумента
var unaryOps = map[string]bool{
	"sqrt": true,
	"log":  true,
	"abs":  true,
}

// ====================================================
// ФУНКЦИЯ: добавление в историю
// ====================================================

func addToHistory(expr string, result float64) {
	record := CalcRecord{
		expression: expr,
		result:     result,
	}
	history = append([]CalcRecord{record}, history...)
	if len(history) > HistorySize {
		history = history[:HistorySize]
	}
}

// ====================================================
// ФУНКЦИЯ: вывод истории
// ====================================================

func printHistory() {
	if len(history) == 0 {
		fmt.Println("  История пуста")
		return
	}
	fmt.Println("  ── История вычислений ──────────────")
	for i, rec := range history {
		fmt.Printf("  [%d] %s = %g\n", i+1, rec.expression, rec.result)
	}
}

// ====================================================
// ФУНКЦИЯ: очистка истории (ЗАДАНИЕ 2)
// ====================================================

func clearHistory() {
	history = []CalcRecord{}
	fmt.Println("  История очищена")
}

// ====================================================
// ФУНКЦИЯ: вычисление
// ====================================================

func calculate(a float64, op string, b float64) (float64, error) {
	switch op {
	case "+":
		return a + b, nil

	case "-":
		return a - b, nil

	case "*":
		return a * b, nil

	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль: %g / %g", a, b)
		}
		return a / b, nil

	case "%":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль при операции %%")
		}
		return float64(int(a) % int(b)), nil

	case "pow":
		return math.Pow(a, b), nil

	case "sqrt":
		if a < 0 {
			return 0, fmt.Errorf("квадратный корень из отрицательного числа: %g", a)
		}
		return math.Sqrt(a), nil

	case "log":
		// ЗАДАНИЕ 1: натуральный логарифм
		if a <= 0 {
			return 0, fmt.Errorf("логарифм определён только для положительных чисел: %g", a)
		}
		return math.Log(a), nil

	case "abs":
		// ЗАДАНИЕ 1: модуль числа
		return math.Abs(a), nil

	default:
		return 0, fmt.Errorf("неизвестная операция: '%s'\nДоступные: +, -, *, /, %%, pow, sqrt, log, abs", op)
	}
}

// ====================================================
// ФУНКЦИЯ: разбор строки в число
// ====================================================

func parseNumber(s string) (float64, error) {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("неверный формат числа: '%s' (ожидается число, например: 3.14)", s)
	}
	return num, nil
}

// ====================================================
// ФУНКЦИЯ: обработка одного вычисления
// ====================================================

func processCalculation(args []string) {
	if len(args) < 1 {
		fmt.Println("Ошибка: нужно минимум число и операция")
		fmt.Println("Пример: go run main.go 10 + 5")
		fmt.Println("        go run main.go 9 sqrt")
		return
	}

	a, err := parseNumber(args[0])
	if err != nil {
		fmt.Printf("Ошибка в первом аргументе: %v\n", err)
		return
	}

	if len(args) < 2 {
		fmt.Println("Ошибка: не указана операция")
		return
	}
	op := args[1]

	var b float64
	if !unaryOps[op] {
		if len(args) < 3 {
			fmt.Printf("Ошибка: операция '%s' требует два числа\n", op)
			return
		}
		b, err = parseNumber(args[2])
		if err != nil {
			fmt.Printf("Ошибка в третьем аргументе: %v\n", err)
			return
		}
	}

	result, err := calculate(a, op, b)
	if err != nil {
		fmt.Printf("Ошибка вычисления: %v\n", err)
		return
	}

	var expression string
	if unaryOps[op] {
		expression = fmt.Sprintf("%s(%g)", op, a)
	} else {
		expression = fmt.Sprintf("%g %s %g", a, op, b)
	}

	addToHistory(expression, result)
	fmt.Printf("  %s = %g\n", expression, result)
}

// ====================================================
// ФУНКЦИЯ: интерактивный режим
// ====================================================

func interactiveMode() {
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║      Калькулятор — интерактивный режим       ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Println("Введите выражение: <число1> <операция> [число2]")
	fmt.Println("Операции: + - * / % pow sqrt log abs")
	fmt.Println("Команды:  history, clear, quit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("calc> ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		switch strings.ToLower(line) {
		case "quit", "exit", "q":
			fmt.Println("До свидания!")
			return
		case "history", "h":
			printHistory()
			continue
		case "clear", "cls":
			clearHistory()
			continue
		case "help":
			fmt.Println("Операции: + - * / % pow sqrt log abs")
			fmt.Println("Пример: 10 + 5")
			fmt.Println("        9 sqrt      (sqrt использует только первый аргумент)")
			fmt.Println("        2 pow 10    (2^10 = 1024)")
			fmt.Println("        2.718 log   (ln(2.718) ≈ 1)")
			fmt.Println("        -5 abs      (|−5| = 5)")
			continue
		}

		tokens := strings.Fields(line)
		processCalculation(tokens)
	}
}

// ====================================================
// ФУНКЦИЯ: вывод справки
// ====================================================

func printUsage() {
	fmt.Println("Использование:")
	fmt.Println("  go run main.go <число1> <операция> [число2]")
	fmt.Println("  go run main.go           # интерактивный режим")
	fmt.Println()
	fmt.Println("Операции:")
	fmt.Println("  +    сложение         (10 + 5 = 15)")
	fmt.Println("  -    вычитание        (10 - 3 = 7)")
	fmt.Println("  *    умножение        (4 * 5 = 20)")
	fmt.Println("  /    деление          (10 / 4 = 2.5)")
	fmt.Println("  %    остаток          (10 % 3 = 1)")
	fmt.Println("  pow  возведение       (2 pow 10 = 1024)")
	fmt.Println("  sqrt квадр. корень    (16 sqrt = 4)")
	fmt.Println("  log  натур. логарифм  (2.718 log ≈ 1)")
	fmt.Println("  abs  модуль числа     (-5 abs = 5)")
	fmt.Println()
	fmt.Println("Примеры:")
	fmt.Println("  go run main.go 10 + 5")
	fmt.Println("  go run main.go 2.5 * 4")
	fmt.Println("  go run main.go 16 sqrt")
	fmt.Println("  go run main.go 2 pow 8")
	fmt.Println("  go run main.go -9 abs")
}

// ====================================================
// ГЛАВНАЯ ФУНКЦИЯ
// ====================================================

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		interactiveMode()
		return
	}

	if args[0] == "--help" || args[0] == "-h" {
		printUsage()
		return
	}

	processCalculation(args)
}
