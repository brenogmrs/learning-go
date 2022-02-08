package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
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
	sites := []string{
		"https://www.alura.com.br/",
		"https://www.caelum.com.br/",
		"https://www.google.com.br/",
	}

	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			verifySite(site)
		}

		time.Sleep(5 * time.Second)
	}
	
	fmt.Println("")
	
}

func verifySite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("site:", site, "carregado com sucesso!")
	}else {
		fmt.Println("site:", site, "está com problemas. Status Code:", response.StatusCode)
	}
}
