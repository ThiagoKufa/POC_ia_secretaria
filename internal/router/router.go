package router

import (
	"strings"

	"ia_secretaria/internal/agents"
)

// Router é responsável por rotear mensagens para os agentes apropriados
type Router struct {
	agents map[agents.AgentType]agents.Agent
}

// NewRouter cria uma nova instância do Router
func NewRouter(agents map[agents.AgentType]agents.Agent) *Router {
	return &Router{
		agents: agents,
	}
}

// classifyMessage classifica a mensagem para determinar o agente apropriado
func (r *Router) classifyMessage(message string) agents.AgentType {
	msg := strings.ToLower(message)

	// Palavras-chave para cada tipo de agente
	pizzaKeywords := []string{"pizza", "sabor", "ingredientes", "tamanho", "margherita", "calabresa", "portuguesa"}
	paymentKeywords := []string{"pagar", "pagamento", "cartão", "dinheiro", "troco", "valor", "preço", "total"}
	managerKeywords := []string{"atraso", "problema", "reclamação", "gerente", "supervisor", "compensação"}

	// Verifica palavras-chave
	for _, keyword := range pizzaKeywords {
		if strings.Contains(msg, keyword) {
			return agents.AgentType(1) // PizzaioloAgent
		}
	}

	for _, keyword := range paymentKeywords {
		if strings.Contains(msg, keyword) {
			return agents.AgentType(2) // CaixaAgent
		}
	}

	for _, keyword := range managerKeywords {
		if strings.Contains(msg, keyword) {
			return agents.AgentType(3) // GerenteAgent
		}
	}

	// Padrão: atendente
	return agents.AgentType(0) // AttendenteAgent
}

// Route roteia a mensagem para o agente apropriado
func (r *Router) Route(message string) (string, error) {
	// Classifica a mensagem
	agentType := r.classifyMessage(message)

	// Obtém o agente apropriado
	agent, exists := r.agents[agentType]
	if !exists {
		agent = r.agents[agents.AgentType(0)] // AttendenteAgent
	}

	// Processa a mensagem com o agente selecionado
	return agent.ProcessMessage(message)
}
