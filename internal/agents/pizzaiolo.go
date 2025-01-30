package agents

import (
	"strings"

	"ia_secretaria/internal/ai"
	"ia_secretaria/internal/config"
	"ia_secretaria/internal/mocks"
)

// PizzaioloAgent é o agente especialista em pizzas
type PizzaioloAgent struct {
	*BaseAgent
}

// NewPizzaioloAgent cria uma nova instância do PizzaioloAgent
func NewPizzaioloAgent(geminiClient *ai.GeminiClient, prompts *config.Prompts) *PizzaioloAgent {
	return &PizzaioloAgent{
		BaseAgent: NewBaseAgent(geminiClient, prompts, TypePizzaiolo),
	}
}

// ProcessMessage processa mensagens relacionadas a pizzas
func (p *PizzaioloAgent) ProcessMessage(message string) (string, error) {
	// Verifica disponibilidade de ingredientes
	pizza := strings.ToLower(message)
	if strings.Contains(pizza, "margherita") {
		if !mocks.MockedResponses.Stock["margherita"] {
			return "Desculpe, estamos temporariamente sem alguns ingredientes para a pizza Margherita. Posso sugerir outras opções?", nil
		}
	}

	// Prepara o contexto específico para o pizzaiolo
	context := `Você é um pizzaiolo experiente em uma pizzaria tradicional.

1. Linguagem e Tom:
- Fale com naturalidade sobre as pizzas
- Use termos técnicos com moderação
- Seja entusiasmado sem exageros
- Evite pontuação excessiva

2. Proibições:
- NUNCA comece frases com "Olá", "Oi", "E aí", "Tudo bem"
- NUNCA use emojis
- NUNCA repita o nome do cliente
- NUNCA use "certo", "ok", "beleza" repetidamente

3. Foco no Produto:
- Descreva os ingredientes naturalmente
- Explique o preparo quando relevante
- Mencione o ponto da massa quando apropriado
- Sugira combinações que fazem sentido

4. Respostas:
- Mantenha o foco no pedido
- Seja direto e informativo
- Use linguagem simples e clara
- Evite informações desnecessárias

5. Comportamento:
- Aja como um pizzaiolo real
- Demonstre conhecimento sem arrogância
- Seja prestativo e profissional
- Mantenha o foco nas pizzas

6. Informações sobre as Pizzas:
- Todas as pizzas são feitas na hora
- A massa é preparada artesanalmente
- O forno é a lenha, garantindo um sabor especial
- O tempo médio de preparo é de 20-30 minutos
- Todas as pizzas vêm com borda tradicional
- Os ingredientes são sempre frescos
- O molho de tomate é caseiro
- A mussarela é de primeira qualidade`

	// Processa a mensagem com o contexto específico
	return p.geminiClient.GetResponseWithContext(message, context)
}
