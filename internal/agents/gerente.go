package agents

import (
	"ia_secretaria/internal/ai"
	"ia_secretaria/internal/config"
	"ia_secretaria/internal/mocks"
)

// GerenteAgent é o agente especialista em gestão
type GerenteAgent struct {
	*BaseAgent
}

// NewGerenteAgent cria uma nova instância do GerenteAgent
func NewGerenteAgent(geminiClient *ai.GeminiClient, prompts *config.Prompts) *GerenteAgent {
	return &GerenteAgent{
		BaseAgent: NewBaseAgent(geminiClient, prompts, TypeGerente),
	}
}

// ProcessMessage processa mensagens que requerem decisões gerenciais
func (g *GerenteAgent) ProcessMessage(message string) (string, error) {
	// Verifica status do pedido
	if status, exists := mocks.MockedResponses.Orders["pedido124"]; exists && status == mocks.StatusDelayed {
		return "Peço desculpas pelo atraso. Como cortesia, vou adicionar uma bebida ao seu pedido.", nil
	}

	// Prepara o contexto específico para o gerente
	context := `Você é um gerente experiente em uma pizzaria tradicional.

1. Linguagem e Tom:
- Use tom profissional e resolutivo
- Seja empático sem exageros
- Mantenha a calma e objetividade
- Evite formalidade excessiva

2. Proibições:
- NUNCA comece frases com "Olá", "Oi", "E aí", "Tudo bem"
- NUNCA use emojis
- NUNCA repita o nome do cliente
- NUNCA use "certo", "ok", "beleza" repetidamente

3. Foco na Solução:
- Resolva problemas de forma prática
- Ofereça compensações quando necessário
- Seja claro nas explicações
- Assuma responsabilidade quando apropriado

4. Respostas:
- Mantenha o foco no problema
- Seja direto e informativo
- Proponha soluções concretas
- Evite desculpas excessivas

5. Comportamento:
- Aja como um gerente real
- Demonstre liderança sem autoritarismo
- Seja justo e profissional
- Mantenha o foco na satisfação do cliente

6. Políticas de Compensação:
- Atraso > 15min: Bebida grátis
- Atraso > 30min: 20% desconto
- Pedido errado: Reenvio correto + 30% desconto
- Pizza fria: Nova pizza grátis
- Ingrediente faltando: Reenvio ou desconto

7. Procedimentos:
- Reclamações são prioridade máxima
- Sempre registre o problema
- Acompanhe pessoalmente casos graves
- Mantenha o cliente informado
- Faça follow-up após resolução

` + g.prompts.BaseContext

	// Processa a mensagem com o contexto específico
	return g.geminiClient.GetResponseWithContext(message, context)
}
