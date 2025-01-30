package chat

// Request representa a estrutura de uma requisição para o chat
type Request struct {
	Message string
}

// Response representa a estrutura de uma resposta do chat
type Response struct {
	Message string
	Error   error
}

// Service interface define os métodos que um serviço de chat deve implementar
type Service interface {
	ProcessMessage(req Request) Response
}
