package main

import (
	"bufio"
	"fmt"
	"os"
)

type Calculadora struct {
	Numero1 float64
	Numero2 float64
}

func (c Calculadora) Somar() float64 {
	return c.Numero1 + c.Numero2
}

func (c Calculadora) Subtrair() float64 {
	return c.Numero1 - c.Numero2
}

func (c Calculadora) Multiplicar() float64 {
	return c.Numero1 * c.Numero2
}

func (c Calculadora) Dividir() (float64, error) {
	if c.Numero2 == 0 {
		return 0, fmt.Errorf("não é possível dividir por zero")
	}
	return c.Numero1 / c.Numero2, nil
}

func aguardarEnter() {
	leitor := bufio.NewReader(os.Stdin)

	// Consome a quebra de linha deixada pelos fmt.Scan anteriores.
	_, _ = leitor.ReadString('\n')

	fmt.Print("\nTecle ENTER para fechar...")
	_, _ = leitor.ReadString('\n')
}

func main() {
	var calc Calculadora
	var operacao string

	fmt.Println("*** Calculadora Fantástica em Go ***")

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&calc.Numero1)

	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scan(&operacao)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&calc.Numero2)

	switch operacao {
	case "+":
		fmt.Printf("Resultado: %.2f\n", calc.Somar())
	case "-":
		fmt.Printf("Resultado: %.2f\n", calc.Subtrair())
	case "*":
		fmt.Printf("Resultado: %.2f\n", calc.Multiplicar())
	case "/":
		resultado, err := calc.Dividir()
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Printf("Resultado: %.2f\n", resultado)
		}

	default:
		fmt.Println("Operação inválida. Por favor, escolha entre +, -, *, /.")
	}

	aguardarEnter()
}
