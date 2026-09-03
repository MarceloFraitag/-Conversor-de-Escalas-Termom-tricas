package main

import "fmt"

// Função que recebe uma temperatura em Kelvin e retorna o valor convertido para Celsius
func converteKelvinParaCelsius(kelvin float64) float64 {
	return kelvin - 273.0
}

func main() {
	// Definindo a constante com o ponto de ebulição da água em Kelvin
	const tempKelvin = 373.0

	// Chamando a função para realizar o cálculo da conversão
	tempCelsius := converteKelvinParaCelsius(tempKelvin)

	// Exibindo os resultados formatados no console
	fmt.Println("========================================")
	fmt.Println("    Conversor de Escalas Termométricas")
	fmt.Println("========================================")
	fmt.Printf("Temperatura em Kelvin: %.2f K\n", tempKelvin)
	fmt.Printf("Temperatura em Celsius: %.2f °C\n", tempCelsius)
}