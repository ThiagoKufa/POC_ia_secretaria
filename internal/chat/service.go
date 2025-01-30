package chat

import (
	"strings"

	"ia_secretaria/internal/agents"
	"ia_secretaria/internal/ai"
	"ia_secretaria/internal/config"
	"ia_secretaria/internal/router"
)

// ChatService implementa a interface Service
type ChatService struct {
	context      string
	geminiClient *ai.GeminiClient
	router       *router.Router
	history      []string
	isFirstMsg   bool
}

// NewChatService cria uma nova instância do serviço de chat
func NewChatService(cfg *config.Config) *ChatService {
	// Criar cliente Gemini
	geminiClient := ai.NewGeminiClient(cfg.GeminiAPIKey)

	// Criar agentes
	agentMap := make(map[agents.AgentType]agents.Agent)
	agentMap[agents.TypeAtendente] = agents.NewAtendenteAgent(geminiClient, cfg.Prompts)
	agentMap[agents.TypePizzaiolo] = agents.NewPizzaioloAgent(geminiClient, cfg.Prompts)
	agentMap[agents.TypeCaixa] = agents.NewCaixaAgent(geminiClient, cfg.Prompts)
	agentMap[agents.TypeGerente] = agents.NewGerenteAgent(geminiClient, cfg.Prompts)

	// Criar router
	messageRouter := router.NewRouter(agentMap)

	return &ChatService{
		context:      "Sou um assistente virtual projetado para ajudar você com suas tarefas.",
		geminiClient: geminiClient,
		router:       messageRouter,
		history:      make([]string, 0),
		isFirstMsg:   true,
	}
}

// ProcessMessage processa uma mensagem e retorna uma resposta
func (s *ChatService) ProcessMessage(req Request) Response {
	// Converter a mensagem para minúsculas para facilitar a comparação
	msg := strings.ToLower(req.Message)

	// Comandos especiais que não precisam ir para a IA
	switch {
	case contains(msg, []string{"tchau", "adeus", "até"}):
		s.history = nil // Limpa o histórico ao encerrar
		return Response{
			Message: "Foi um prazer ajudar! Tenha um ótimo dia! 😊",
		}
	}

	// Adicionar a mensagem ao histórico
	s.history = append(s.history, "👤 Usuário: "+req.Message)

	// Rotear a mensagem para o agente apropriado
	resposta, err := s.router.Route(req.Message)
	if err != nil {
		return Response{
			Message: "Desculpe, tive um problema ao processar sua mensagem. Pode tentar novamente?",
			Error:   err,
		}
	}

	// Adicionar a resposta ao histórico
	s.history = append(s.history, "🤖 Assistente: "+resposta)

	// Marcar que não é mais a primeira mensagem
	s.isFirstMsg = false

	return Response{
		Message: resposta,
	}
}

// contains verifica se alguma das palavras-chave está presente na mensagem
func contains(message string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(message, keyword) {
			return true
		}
	}
	return false
}
