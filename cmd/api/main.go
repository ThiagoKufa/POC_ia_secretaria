package main

import (
	"fmt"
	"os"

	"ia_secretaria/internal/chat"
	"ia_secretaria/internal/config"
	"ia_secretaria/internal/ui"
)

func main() {
	// Carregar configurações
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Erro ao carregar configurações: %v\n", err)
		fmt.Println("Por favor, defina a variável de ambiente GEMINI_API_KEY")
		fmt.Println("Exemplo: export GEMINI_API_KEY=sua-chave-api")
		os.Exit(1)
	}

	// Criar e inicializar o serviço de chat
	chatService := chat.NewChatService(cfg)

	// Criar e inicializar a interface do terminal
	terminal := ui.NewTerminalUI(chatService)

	// Iniciar o chat
	terminal.Start()
}
