# IA Secretária - Sistema de Atendimento para Pizzaria

## 📝 Descrição
Sistema inteligente de atendimento para pizzaria que utiliza a API Gemini do Google para processar pedidos, responder dúvidas e gerenciar interações com clientes. O sistema é composto por diferentes agentes especializados que trabalham em conjunto para fornecer um atendimento completo e personalizado.

## 🚀 Funcionalidades

### 🤖 Agentes Especializados

#### 1. Atendente Virtual
- Primeiro contato com o cliente
- Informações gerais sobre cardápio
- Horários de funcionamento
- Informações sobre entrega
- Preços e promoções

#### 2. Pizzaiolo Virtual
- Especialista em pizzas
- Informações sobre ingredientes
- Sugestões de sabores
- Verificação de disponibilidade
- Tempo de preparo

#### 3. Caixa Virtual
- Processamento de pagamentos
- Informações sobre formas de pagamento
- Cálculo de trocos
- Emissão de comprovantes
- Parcelamentos disponíveis

#### 4. Gerente Virtual
- Tratamento de reclamações
- Resolução de problemas
- Políticas de compensação
- Acompanhamento de pedidos atrasados
- Follow-up com clientes

## 💻 Tecnologias Utilizadas

- Linguagem: Go (Golang)
- IA: Google Gemini API
- Arquitetura: Clean Architecture
- Padrões: SOLID, DDD

## 🏗️ Estrutura do Projeto

```
.
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   ├── agents/                 # Agentes especializados
│   │   ├── atendente.go       # Agente de atendimento geral
│   │   ├── pizzaiolo.go       # Especialista em pizzas
│   │   ├── caixa.go           # Especialista em pagamentos
│   │   └── gerente.go         # Especialista em gestão
│   ├── ai/
│   │   └── gemini.go          # Cliente da API Gemini
│   ├── chat/
│   │   ├── service.go         # Serviço principal de chat
│   │   └── types.go           # Tipos e interfaces do chat
│   ├── config/
│   │   └── prompts.go         # Configurações e prompts
│   ├── mocks/
│   │   └── responses.go       # Mocks para testes
│   └── router/
│       └── router.go          # Roteamento de mensagens
└── README.md
```

## 🔄 Fluxo de Funcionamento

1. **Entrada da Mensagem**
   - Cliente envia mensagem
   - `ChatService` recebe e processa
   - Adiciona ao histórico da conversa

2. **Roteamento**
   - `Router` analisa a mensagem
   - Identifica palavras-chave
   - Direciona para o agente apropriado

3. **Processamento pelo Agente**
   - Agente especializado recebe a mensagem
   - Aplica contexto específico
   - Realiza verificações necessárias
   - Envia para a IA com contexto apropriado

4. **Resposta da IA**
   - Gemini processa com contexto
   - Mantém histórico da conversa
   - Gera resposta personalizada
   - Retorna ao cliente

## 🛠️ Configuração e Instalação

### Pré-requisitos
- Go 1.21 ou superior
- Chave de API do Google Gemini

### Instalação
1. Clone o repositório
```bash
git clone https://github.com/ThiagoKufa/POC_ia_secretaria
cd POC_ia_secretaria
```

2. Configure as variáveis de ambiente
```bash
export GEMINI_API_KEY=sua-chave-api
```

3. Instale as dependências
```bash
go mod download
```

4. Execute o projeto
```bash
go run cmd/main.go
```

## 🔐 Segurança

- Chaves de API armazenadas em variáveis de ambiente
- Sanitização de entradas do usuário
- Logs seguros sem informações sensíveis
- Tratamento de erros apropriado

## 📋 Cardápio e Preços

### Pizzas
- **Margherita**
  - Broto (25cm): R$ 20,00
  - Grande (35cm): R$ 35,00
  
- **Calabresa**
  - Broto (25cm): R$ 25,00
  - Grande (35cm): R$ 45,00
  
- **Quatro Queijos**
  - Broto (25cm): R$ 22,00
  - Grande (35cm): R$ 40,00
  
- **Portuguesa**
  - Broto (25cm): R$ 24,00
  - Grande (35cm): R$ 42,00

### Bebidas
- Refrigerante 300ml: R$ 5,00
- Suco Natural 300ml: R$ 6,00
- Cerveja (Lata): R$ 8,00

## 📦 Políticas de Entrega

- Tempo médio: 30-45 minutos
- Taxa de entrega:
  - Até 3km: R$ 5,00
  - Até 5km: R$ 8,00
- Pedido mínimo: R$ 30,00

## ⚠️ Políticas de Compensação

- Atraso > 15min: Bebida grátis
- Atraso > 30min: 20% desconto
- Pedido errado: Reenvio + 30% desconto
- Pizza fria: Nova pizza grátis
- Ingrediente faltando: Reenvio ou desconto

## 🤝 Contribuindo

1. Fork o projeto
2. Crie sua branch de feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👥 Autores

- Thiago Kufa - [@ThiagoKufa](https://github.com/ThiagoKufa)
