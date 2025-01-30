package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"ia_secretaria/internal/config"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GeminiClient struct {
	apiKey  string
	prompts *config.Prompts
	history []Message
}

type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error,omitempty"`
}

func NewGeminiClient(apiKey string) *GeminiClient {
	return &GeminiClient{
		apiKey:  apiKey,
		prompts: config.DefaultPrompts(),
		history: make([]Message, 0),
	}
}

// GetPrompts retorna os prompts configurados
func (c *GeminiClient) GetPrompts() *config.Prompts {
	return c.prompts
}

func (c *GeminiClient) GetResponse(message string) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=%s", c.apiKey)

	// Adiciona a mensagem do usuário ao histórico
	c.history = append(c.history, Message{
		Role:    "user",
		Content: message,
	})

	// Constrói a mensagem completa com o histórico
	var mensagemCompleta string
	mensagemCompleta = c.prompts.BaseContext + "\n\nHistórico da conversa:\n"

	for _, msg := range c.history {
		if msg.Role == "user" {
			mensagemCompleta += fmt.Sprintf("Cliente: %s\n", msg.Content)
		} else {
			mensagemCompleta += fmt.Sprintf("Atendente: %s\n", msg.Content)
		}
	}

	reqBody := GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{
						Text: mensagemCompleta,
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erro ao criar JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %v", err)
	}

	var geminiResp GeminiResponse
	err = json.Unmarshal(body, &geminiResp)
	if err != nil {
		return "", fmt.Errorf("erro ao decodificar resposta: %v\nResposta bruta: %s", err, string(body))
	}

	if geminiResp.Error.Code != 0 {
		return "", fmt.Errorf("erro da API: %s - %s", geminiResp.Error.Status, geminiResp.Error.Message)
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		response := geminiResp.Candidates[0].Content.Parts[0].Text

		// Adiciona a resposta ao histórico
		c.history = append(c.history, Message{
			Role:    "assistant",
			Content: response,
		})

		return response, nil
	}

	return "", fmt.Errorf("nenhuma resposta recebida")
}

func (c *GeminiClient) GetResponseWithContext(message string, context string) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=%s", c.apiKey)

	// Adiciona a mensagem do usuário ao histórico
	c.history = append(c.history, Message{
		Role:    "user",
		Content: message,
	})

	// Constrói a mensagem completa com o contexto e histórico
	var mensagemCompleta string
	mensagemCompleta = context + "\n\nHistórico da conversa:\n"

	for _, msg := range c.history {
		if msg.Role == "user" {
			mensagemCompleta += fmt.Sprintf("Cliente: %s\n", msg.Content)
		} else {
			mensagemCompleta += fmt.Sprintf("Atendente: %s\n", msg.Content)
		}
	}

	reqBody := GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{
						Text: mensagemCompleta,
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erro ao criar JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %v", err)
	}

	var geminiResp GeminiResponse
	err = json.Unmarshal(body, &geminiResp)
	if err != nil {
		return "", fmt.Errorf("erro ao decodificar resposta: %v\nResposta bruta: %s", err, string(body))
	}

	if geminiResp.Error.Code != 0 {
		return "", fmt.Errorf("erro da API: %s - %s", geminiResp.Error.Status, geminiResp.Error.Message)
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		response := geminiResp.Candidates[0].Content.Parts[0].Text

		// Adiciona a resposta ao histórico
		c.history = append(c.history, Message{
			Role:    "assistant",
			Content: response,
		})

		return response, nil
	}

	return "", fmt.Errorf("nenhuma resposta recebida")
}
