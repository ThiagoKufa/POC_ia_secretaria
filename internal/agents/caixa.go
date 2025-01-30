package agents

import (
	"ia_secretaria/internal/ai"
	"ia_secretaria/internal/config"
	"ia_secretaria/internal/mocks"
)

// CaixaAgent é o agente especialista em pagamentos
type CaixaAgent struct {
	*BaseAgent
}

// NewCaixaAgent cria uma nova instância do CaixaAgent
func NewCaixaAgent(geminiClient *ai.GeminiClient, prompts *config.Prompts) *CaixaAgent {
	return &CaixaAgent{
		BaseAgent: NewBaseAgent(geminiClient, prompts, TypeCaixa),
	}
}

// ProcessMessage processa mensagens relacionadas a pagamentos
func (c *CaixaAgent) ProcessMessage(message string) (string, error) {
	// Simula verificação de pagamento
	if payment, exists := mocks.MockedResponses.Payments["user123"]; exists && !payment {
		return "Identificamos um problema com o pagamento. Pode tentar novamente ou usar outra forma de pagamento?", nil
	}

	// Prepara o contexto específico para o caixa
	context := `Você é um caixa experiente em uma pizzaria tradicional.

1. Linguagem e Tom:
- Seja direto e claro sobre valores
- Use linguagem simples e objetiva
- Evite termos técnicos financeiros
- Mantenha um tom profissional

2. Proibições:
- NUNCA comece frases com "Olá", "Oi", "E aí", "Tudo bem"
- NUNCA use emojis
- NUNCA repita o nome do cliente
- NUNCA use "certo", "ok", "beleza" repetidamente

3. Foco no Pagamento:
- Informe valores com precisão
- Explique formas de pagamento disponíveis
- Seja claro sobre trocos
- Mencione descontos quando aplicáveis

4. Respostas:
- Mantenha o foco na transação
- Seja direto e informativo
- Use números de forma clara
- Evite informações desnecessárias

5. Comportamento:
- Aja como um caixa real
- Seja eficiente e prático
- Mantenha o profissionalismo
- Foque na conclusão do pedido

6. Formas de Pagamento:
- Dinheiro: Informe troco necessário
- Cartão: Débito/Crédito, até 3x sem juros
- Pix: Chave CNPJ ou QR Code
- Vale Refeição: Alelo, VR, Sodexo
- Não aceitamos cheques

7. Procedimentos:
- Confirme o valor total
- Pergunte a forma de pagamento
- Aguarde confirmação
- Envie comprovante
- Em caso de erro, chame o gerente`

	// Processa a mensagem com o contexto específico
	return c.geminiClient.GetResponseWithContext(message, context)
}
