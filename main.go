package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tokenKind int

const (
	tokenNumber tokenKind = iota
	tokenOperator
	tokenLeftParen
	tokenRightParen
)

type token struct {
	kind  tokenKind
	value string
}

func tokenize(input string) ([]token, error) {
	var tokens []token

	for i := 0; i < len(input); {
		ch := input[i]

		switch {
		case ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r':
			i++
		case (ch >= '0' && ch <= '9') || ch == '.':
			start := i
			dotCount := 0

			for i < len(input) {
				current := input[i]

				if current == '.' {
					dotCount++
					if dotCount > 1 {
						return nil, fmt.Errorf("son noto'g'ri yozilgan: %q", input[start:i+1])
					}
					i++
					continue
				}

				if current < '0' || current > '9' {
					break
				}

				i++
			}

			numberText := input[start:i]
			if numberText == "." {
				return nil, errors.New("faqat nuqta son bo'la olmaydi")
			}

			if _, err := strconv.ParseFloat(numberText, 64); err != nil {
				return nil, fmt.Errorf("son noto'g'ri formatda: %q", numberText)
			}

			tokens = append(tokens, token{kind: tokenNumber, value: numberText})
		case ch == '+' || ch == '-' || ch == '*' || ch == '/':
			tokens = append(tokens, token{kind: tokenOperator, value: string(ch)})
			i++
		case ch == '(':
			tokens = append(tokens, token{kind: tokenLeftParen, value: string(ch)})
			i++
		case ch == ')':
			tokens = append(tokens, token{kind: tokenRightParen, value: string(ch)})
			i++
		default:
			return nil, fmt.Errorf("ruxsat etilmagan belgi topildi: %q", string(ch))
		}
	}

	if len(tokens) == 0 {
		return nil, errors.New("ifoda bo'sh")
	}

	return tokens, nil
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func applyOperation(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("0 ga bo'lish mumkin emas")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("noto'g'ri operator: %s", op)
	}
}

func popAndApply(values *[]float64, ops *[]string) error {
	if len(*ops) == 0 {
		return errors.New("operator topilmadi")
	}

	if len(*values) < 2 {
		return errors.New("ifoda to'liq emas")
	}

	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]

	b := (*values)[len(*values)-1]
	*values = (*values)[:len(*values)-1]

	a := (*values)[len(*values)-1]
	*values = (*values)[:len(*values)-1]

	result, err := applyOperation(a, b, op)
	if err != nil {
		return err
	}

	*values = append(*values, result)
	return nil
}

func evaluate(tokens []token) (float64, error) {
	var values []float64
	var ops []string

	expectOperand := true

	for i := 0; i < len(tokens); i++ {
		current := tokens[i]

		switch current.kind {
		case tokenNumber:
			if !expectOperand {
				return 0, fmt.Errorf("operator kutilgan joyda son keldi: %q", current.value)
			}

			value, err := strconv.ParseFloat(current.value, 64)
			if err != nil {
				return 0, fmt.Errorf("sonni o'qib bo'lmadi: %q", current.value)
			}

			values = append(values, value)
			expectOperand = false

		case tokenLeftParen:
			if !expectOperand {
				return 0, errors.New("qavsdan oldin operator bo'lishi kerak")
			}

			ops = append(ops, current.value)

		case tokenRightParen:
			if expectOperand {
				return 0, errors.New("yopuvchi qavsdan oldin qiymat bo'lishi kerak")
			}

			foundLeftParen := false
			for len(ops) > 0 {
				if ops[len(ops)-1] == "(" {
					ops = ops[:len(ops)-1]
					foundLeftParen = true
					break
				}

				if err := popAndApply(&values, &ops); err != nil {
					return 0, err
				}
			}

			if !foundLeftParen {
				return 0, errors.New("qavslar juftligi noto'g'ri")
			}

			expectOperand = false

		case tokenOperator:
			if expectOperand {
				if current.value != "+" && current.value != "-" {
					return 0, fmt.Errorf("operator noto'g'ri joyda ishlatilgan: %s", current.value)
				}

				if i+1 >= len(tokens) {
					return 0, errors.New("ifoda operator bilan tugab qolgan")
				}

				next := tokens[i+1]

				if next.kind == tokenNumber {
					numberValue, err := strconv.ParseFloat(next.value, 64)
					if err != nil {
						return 0, fmt.Errorf("sonni o'qib bo'lmadi: %q", next.value)
					}

					if current.value == "-" {
						numberValue = -numberValue
					}

					values = append(values, numberValue)
					expectOperand = false
					i++
					continue
				}

				if next.kind == tokenLeftParen {
					values = append(values, 0)
				} else {
					return 0, fmt.Errorf("operator'dan keyin qiymat yoki qavs kelishi kerak: %s", current.value)
				}
			} else {
				for len(ops) > 0 && ops[len(ops)-1] != "(" && precedence(ops[len(ops)-1]) >= precedence(current.value) {
					if err := popAndApply(&values, &ops); err != nil {
						return 0, err
					}
				}
			}

			ops = append(ops, current.value)
			expectOperand = true
		}
	}

	if expectOperand {
		return 0, errors.New("ifoda to'liq emas")
	}

	for len(ops) > 0 {
		if ops[len(ops)-1] == "(" {
			return 0, errors.New("ochilgan qavs yopilmagan")
		}

		if err := popAndApply(&values, &ops); err != nil {
			return 0, err
		}
	}

	if len(values) != 1 {
		return 0, errors.New("ifodani hisoblash yakunlanmadi")
	}

	return values[0], nil
}

func evaluateExpression(expression string) (float64, error) {
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	return evaluate(tokens)
}

func formatResult(result float64) string {
	return strconv.FormatFloat(result, 'f', -1, 64)
}

func printUsage() {
	fmt.Println("CLI Calculator")
	fmt.Println("Foydalanish:")
	fmt.Println("  go run main.go \"2+3*4\"")
	fmt.Println("  go run main.go 10 / ( 2 + 3 )")
	fmt.Println("  go run main.go")
	fmt.Println("")
	fmt.Println("Interaktiv rejimda terminalni tozalash uchun: clear")
	fmt.Println("Interaktiv rejimdan chiqish uchun: exit yoki quit")
}

func interactiveMode() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("CLI Calculator interaktiv rejimi")
	fmt.Println("Ifoda kiriting, `exit` yoki `quit` chiqadi.")

	for {
		fmt.Print("calc> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("O'qishda xatolik:", err)
			return
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		switch input {
		case "exit", "quit":
			fmt.Println("Dastur yakunlandi.")
			return
		case "help", "--help", "-h":
			printUsage()
			continue
		}

		result, calcErr := evaluateExpression(input)
		if calcErr != nil {
			fmt.Println("Xatolik:", calcErr)
			continue
		}

		fmt.Println("=", formatResult(result))
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		interactiveMode()
		return
	}

	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h" || args[0] == "help") {
		printUsage()
		return
	}

	expression := strings.Join(args, " ")

	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Println("Xatolik:", err)
		os.Exit(1)
	}

	fmt.Println(formatResult(result))
}
