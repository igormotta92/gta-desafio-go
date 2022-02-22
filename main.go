package main

import (
	"fmt"
)

type Client struct {
	name string
	quantity uint
	total float32
}

func (c *Client) calculateTotal() {
	price := float32(15.50)
	c.total = float32(c.quantity) * price
}

func main() {

	var clients []Client
	
	fmt.Println("Bem-vindo ao sistema de reservade tickets")
	
	for {
		var clientName string
		var qtd uint
		var confirm string
	
		fmt.Printf("Qual seu nome? ")
		fmt.Scan(&clientName)

		fmt.Printf("Quantos tickets você quer reservar? ")
	
		fmt.Scan(&qtd)
		
		fmt.Println(clientName,", você confirma a reserva de", qtd, "tickets? (s/n): ")
	
		fmt.Scan(&confirm)

		if confirm == "n" {
			continue
		}
	
		client := Client{
			name: clientName,
			quantity: qtd,
		}

		client.calculateTotal()
	
		clients = append(clients, client)
	
		fmt.Println("Wow.. seus tickets foram reservados!!!")
	
		fmt.Printf("Quer reservar mais tickets? (s/n): ")
		
		fmt.Scan(&confirm)

		if confirm == "n" {
			break
		}
		
	}

	
	tabela(clients)
}


func StrPadLeft(input string, padLength int, padString string) string {
	output := padString

	for padLength > len(output) {
		output += output
	}

	if len(input) >= padLength {
		return input
	}

	return output[:padLength-len(input)] + input
}

func tabela(clients []Client) {
	
	nome := StrPadLeft("Nome", 10, " ")
	qtd := StrPadLeft("Quantidade", 10, " ")
	total := StrPadLeft("Total", 10, " ")

	head := fmt.Sprintf("%s %s %s \n\r", nome, qtd, total)
	row := fmt.Sprintf("%s \n\r", head)
	fmt.Printf(row)
	fmt.Printf("------------------------------\r\n")

	for index := range clients {

		stotal := fmt.Sprintf("%f", clients[index].total)
		sqtd := fmt.Sprintf("%d", clients[index].quantity)

		name := StrPadLeft(clients[index].name, 10, " ")
		qtd := StrPadLeft(sqtd, 10, " ")
		total := StrPadLeft(stotal, 10, " ")


		line := fmt.Sprintf("%s %s %s \n\r", name, qtd, total)

		fmt.Printf(line)
	}
	
}

