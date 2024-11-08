import (
	"fmt"
)

// Stack для хранения значений и операторов
type Stack []float64

// Push добавляет значение в стек
func (s *Stack) Push(value float64) {
	*s = append(*s, value)
}

// Pop удаляет и возвращает верхнее значение стека
func (s *Stack) Pop() float64 {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

// Функция для вычисления выражения
func Calc(expression string) (float64, error) {
	var values Stack
	var ops Stack

	for i := 0; i < len(expression); i++ {
		c := expression[i]

		// Игнорируем пробелы
		if c == ' ' {
			continue
		}

		// Если текущий символ - цифра, считываем полное число
		if isDigit(c) {
			num := 0
			for i < len(expression) && isDigit(expression[i]) {
				num = num*10 + int(expression[i]-'0')
				i++
			}
			i-- // уменьшаем i, чтобы вернуться на один символ
			values.Push(float64(num))
		} else if c == '(' {
			ops.Push(0) // Используем 0 для обозначения открывающей скобки
		} else if c == ')' {
			// Обрабатываем до открывающей скобки
			for len(ops) > 0 && ops[len(ops)-1] != 0 {
				values.Push(applyOperation(values.Pop(), values.Pop(), ops.Pop()))
			}
			if len(ops) == 0 {
				return 0, fmt.Errorf("несоответствующая скобка")
			}
			ops.Pop() // Убираем открывающую скобку
		} else if isOperator(c) {
			// Обрабатываем все операции с большим или равным приоритетом
			for len(ops) > 0 && precedence(c) <= precedence(byte(ops[len(ops)-1])) {
				values.Push(applyOperation(values.Pop(), values.Pop(), ops.Pop()))
			}
			ops.Push(float64(c)) // Сохраняем оператор
		}
	}

	// Обработка оставшихся значений и операторов
	for len(ops) > 0 {
		values.Push(applyOperation(values.Pop(), values.Pop(), ops.Pop()))
	}

	return values.Pop(), nil
}

// Проверяет, является ли символ цифрой
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// Проверяет, является ли символ оператором
func isOperator(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

// Возвращает приоритет оператора
func precedence(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

// Применяет операцию к двум числам
func applyOperation(b, a float64, op float64) float64 {
	switch rune(op) {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		if b == 0 {
			panic("division by zero")
		}
		return a / b
	}
	return 0
}