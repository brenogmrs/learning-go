package main

import "fmt"

func main() {

	var nome = "Breno"
	
	vesao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", vesao)

	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2 - exibir logs")
	fmt.Println("0 - sair do programa")

	var comando int
	// fmt.Scan() -> infere o tipo
	fmt.Scan(&comando)
	fmt.Println("o comando escolhido foi,", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço este comando")
	}

}