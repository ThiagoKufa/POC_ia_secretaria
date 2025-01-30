package agents

import (
	"ia_secretaria/internal/ai"
	"ia_secretaria/internal/config"
)

// AgentType define os tipos de agentes disponíveis
type AgentType uint8

const (
	TypeAtendente AgentType = iota
	TypePizzaiolo
	TypeCaixa
	TypeGerente
)

// Agent define a interface base para todos os agentes
type Agent interface {
	ProcessMessage(message string) (string, error)
	GetType() AgentType
}

// BaseAgent contém a estrutura base comum a todos os agentes
type BaseAgent struct {
	geminiClient *ai.GeminiClient
	prompts      *config.Prompts
	agentType    AgentType
}

// GetType retorna o tipo do agente
func (b *BaseAgent) GetType() AgentType {
	return b.agentType
}

// NewBaseAgent cria uma nova instância de BaseAgent
func NewBaseAgent(geminiClient *ai.GeminiClient, prompts *config.Prompts, agentType AgentType) *BaseAgent {
	return &BaseAgent{
		geminiClient: geminiClient,
		prompts:      prompts,
		agentType:    agentType,
	}
}
