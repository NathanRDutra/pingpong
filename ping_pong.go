package main


import (
	"fmt"
	"time"
	)

func main() {
	pingChan := make(chan bool)
	pongChan := make(chan bool)

	// Goroutine do ping
	go func() {
		for {
			<-pingChan          // espera a vez de ping
			fmt.Println("ping") // faz o ping
			time.Sleep(1000 * time.Millisecond)
			pongChan <- true    // passa a vez para pong
		}
	}()

	// Goroutine do pong
	go func() {
		for {
			<-pongChan          // espera a vez de pong
			fmt.Println("pong")
			time.Sleep(1000 * time.Millisecond) // faz o pong
			pingChan <- true    // passa a vez para ping
		}
	}()

	// Inicia o jogo passando a vez para o ping
	pingChan <- true

	// Mantém o programa rodando
	select {}
}
