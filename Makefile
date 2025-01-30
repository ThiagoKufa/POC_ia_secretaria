.PHONY: run build test clean help setup env prompts

# Variáveis
APP_NAME=ia_secretaria
BUILD_DIR=build
MAIN_FILE=cmd/api/main.go
VERSION=1.0.0

# Cores para output
CYAN=\033[0;36m
GREEN=\033[0;32m
YELLOW=\033[0;33m
RED=\033[0;31m
BLUE=\033[0;34m
PURPLE=\033[0;35m
BOLD=\033[1m
NC=\033[0m

# Carrega variáveis do arquivo .env se ele existir
ifneq (,$(wildcard .env))
    include .env
    export
endif

define print_help
    @awk -F ':.*##' '/\[$(1)\]/ { printf "  $(GREEN)%-20s$(NC) %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
endef

help: ## Exibe este menu de ajuda com todos os comandos disponíveis
	@echo "$(BOLD)$(BLUE)IA Secretaria v$(VERSION)$(NC)"
	@echo "$(BOLD)Uso:$(NC) make $(GREEN)<comando>$(NC)"
	@echo ""
	@echo "$(BOLD)$(PURPLE)🔧 Comandos de Desenvolvimento:$(NC)"
	$(call print_help,Desenvolvimento)
	@echo ""
	@echo "$(BOLD)$(CYAN)🚀 Comandos de Execução:$(NC)"
	$(call print_help,Execução)
	@echo ""
	@echo "$(BOLD)$(YELLOW)🧪 Comandos de Teste:$(NC)"
	$(call print_help,Teste)
	@echo ""
	@echo "$(BOLD)$(RED)🧹 Comandos de Limpeza:$(NC)"
	$(call print_help,Limpeza)
	@echo ""
	@echo "$(BOLD)Exemplo de uso:$(NC)"
	@echo "  make setup  # Configura o ambiente"
	@echo "  make run    # Executa a aplicação"

setup: ## [Desenvolvimento] Configura o ambiente inicial, cria arquivo .env e instala dependências
	@echo "$(CYAN)Configurando ambiente de desenvolvimento...$(NC)"
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "$(YELLOW)Arquivo .env criado. Por favor, configure suas variáveis de ambiente.$(NC)"; \
		echo "$(YELLOW)Edite o arquivo .env e adicione sua chave API do Gemini.$(NC)"; \
	fi
	@go mod tidy
	@echo "$(GREEN)Ambiente configurado com sucesso!$(NC)"

run: ## [Execução] Inicia o chat interativo com a IA
	@echo "$(CYAN)Executando a aplicação...$(NC)"
	@if [ -z "$(GEMINI_API_KEY)" ]; then \
		if [ ! -f .env ]; then \
			echo "$(YELLOW)⚠️  Arquivo .env não encontrado!$(NC)"; \
			echo "Execute 'make setup' primeiro para criar o arquivo .env"; \
			exit 1; \
		else \
			echo "$(YELLOW)⚠️  GEMINI_API_KEY não está definida no arquivo .env!$(NC)"; \
			echo "Por favor, adicione sua chave API no arquivo .env:"; \
			echo "GEMINI_API_KEY=sua-chave-api"; \
			exit 1; \
		fi \
	fi
	@go run $(MAIN_FILE)

build: ## [Execução] Compila o projeto para um executável binário
	@echo "$(CYAN)Compilando a aplicação...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "$(GREEN)Aplicação compilada com sucesso em $(BUILD_DIR)/$(APP_NAME)$(NC)"

test: ## [Teste] Executa todos os testes do projeto
	@echo "$(CYAN)Executando testes...$(NC)"
	@go test ./... -v

clean: ## [Limpeza] Remove arquivos temporários, builds e caches
	@echo "$(CYAN)Limpando arquivos temporários...$(NC)"
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "$(GREEN)Limpeza concluída!$(NC)"

lint: ## [Teste] Executa análise estática do código com golangci-lint
	@echo "$(CYAN)Executando linter...$(NC)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "$(YELLOW)golangci-lint não está instalado. Instalando...$(NC)"; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

env: ## [Desenvolvimento] Exibe o status das variáveis de ambiente configuradas
	@echo "$(CYAN)Status das Variáveis de Ambiente:$(NC)"
	@echo "  GEMINI_API_KEY: $(if $(GEMINI_API_KEY),$(GREEN)✓ definida$(NC),$(RED)✗ não definida$(NC))"

prompts: ## [Desenvolvimento] Exibe os prompts de contexto configurados
	@echo "$(CYAN)Prompts de Contexto:$(NC)"
	@echo "$(YELLOW)Arquivo:$(NC) internal/config/prompts.go"
	@echo ""
	@echo "$(BOLD)Contexto Base:$(NC)"
	@cat internal/config/prompts.go | grep -A 20 'BaseContext:' | grep -v 'BaseContext:' | sed 's/^[[:space:]]*//g' | sed 's/`,//g' | grep -v '^$$'

# Define o target padrão
.DEFAULT_GOAL := help 