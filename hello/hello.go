package main

import (
	"fmt"
	"os"
	//"net/http"
)

func main() {

	showIntro()

	for {
		showMenu()	
	
		command := readCommand()
	
		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0: 
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default: 
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}


func showIntro() {
	name := "Breno"
	version := 1.1
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2 - exibir logs")
	fmt.Println("0 - sair do programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("o comando escolhido foi,", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://www.alura.com.br/"
	sites[1] = "https://www.google.com.br/"
	sites[2] = "https://www.instagram.com.br/"
	sites[3] = "https://www.facebook.com.br/"

	
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("site:", site, "carregado com sucesso!")
	}else {
		fmt.Println("site:", site, "está com problemas. Status Code:", response.StatusCode)
	}
	
}
