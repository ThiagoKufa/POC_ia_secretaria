package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"ia_secretaria/internal/chat"

	"github.com/fatih/color"
)

// TerminalUI gerencia a interface do usuário no terminal
type TerminalUI struct {
	chatService chat.Service
	scanner     *bufio.Scanner
	// Cores para diferentes elementos da UI
	titleColor     *color.Color
	userColor      *color.Color
	assistantColor *color.Color
	errorColor     *color.Color
	infoColor      *color.Color
}

// NewTerminalUI cria uma nova instância da interface do terminal
func NewTerminalUI(chatService chat.Service) *TerminalUI {
	return &TerminalUI{
		chatService:    chatService,
		scanner:        bufio.NewScanner(os.Stdin),
		titleColor:     color.New(color.FgHiCyan, color.Bold),
		userColor:      color.New(color.FgGreen, color.Bold),
		assistantColor: color.New(color.FgBlue),
		errorColor:     color.New(color.FgRed, color.Bold),
		infoColor:      color.New(color.FgYellow),
	}
}

// clearScreen limpa a tela do terminal
func (ui *TerminalUI) clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// showWelcome mostra a mensagem de boas-vindas
func (ui *TerminalUI) showWelcome() {
	ui.clearScreen()
	ui.titleColor.Println("\n╔═══════════════════════════════════╗")
	ui.titleColor.Println("║        Chat Interativo IA         ║")
	ui.titleColor.Println("╚═══════════════════════════════════╝")

	ui.infoColor.Println("\nComandos disponíveis:")
	fmt.Println("• Digite 'ajuda' para ver todos os comandos")
	fmt.Println("• Digite 'limpar' para limpar a tela")
	fmt.Println("• Digite 'sair' para encerrar o chat")

	ui.titleColor.Println("\n═══════════════════════════════════════")
}

// showTypingAnimation mostra uma animação de "digitando..."
func (ui *TerminalUI) showTypingAnimation() {
	dots := []string{".", "..", "..."}
	for i := 0; i < 3; i++ {
		for _, dot := range dots {
			fmt.Printf("\rProcessando%s   ", dot)
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Print("\r                     \r")
}

// showHelp mostra a lista de comandos disponíveis
func (ui *TerminalUI) showHelp() {
	ui.infoColor.Println("\nLista de comandos disponíveis:")
	fmt.Println("• ajuda   - Mostra esta lista de comandos")
	fmt.Println("• limpar  - Limpa a tela do chat")
	fmt.Println("• sair    - Encerra o chat")
	fmt.Println("• versão  - Mostra a versão do sistema")
}

// Start inicia a interface do usuário
func (ui *TerminalUI) Start() {
	ui.showWelcome()

	for {
		ui.userColor.Print("\n\n👤 Você: ")
		if !ui.scanner.Scan() {
			break
		}

		input := strings.TrimSpace(ui.scanner.Text())
		if input == "" {
			continue
		}

		// Processar comandos especiais
		switch strings.ToLower(input) {
		case "sair":
			ui.assistantColor.Println("\n👋 Até logo! Tenha um ótimo dia!")
			return
		case "limpar":
			ui.clearScreen()
			ui.showWelcome()
			continue
		case "ajuda":
			ui.showHelp()
			continue
		case "versão":
			ui.infoColor.Println("\nVersão 1.0.0 - IA Secretaria")
			continue
		}

		ui.showTypingAnimation()
		response := ui.chatService.ProcessMessage(chat.Request{Message: input})

		if response.Error != nil {
			ui.errorColor.Printf("\n❌ Erro: %v\n", response.Error)
			continue
		}

		ui.assistantColor.Printf("\n🤖 Assistente: %s\n", response.Message)
	}
}
