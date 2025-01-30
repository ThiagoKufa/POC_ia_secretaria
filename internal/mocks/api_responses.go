package mocks

// MockStatus representa o status de um pedido
type MockStatus string

const (
	StatusOK      MockStatus = "ok"
	StatusDelayed MockStatus = "atrasado"
	StatusPending MockStatus = "pendente"
)

// MockedResponses contém todas as respostas mockadas para simular APIs
var MockedResponses = struct {
	Orders   map[string]MockStatus
	Payments map[string]bool
	Stock    map[string]bool
	Delivery map[string]int
}{
	Orders: map[string]MockStatus{
		"pedido123": StatusOK,
		"pedido124": StatusDelayed,
		"pedido125": StatusPending,
	},
	Payments: map[string]bool{
		"user123": true,  // pagamento ok
		"user124": false, // problema no pagamento
	},
	Stock: map[string]bool{
		"margherita": true,
		"calabresa":  true,
		"portuguesa": false, // sem ingredientes
		"mussarela":  true,
		"manjericao": true,
		"tomate":     true,
	},
	Delivery: map[string]int{
		"zona_sul":   30, // 30 minutos
		"zona_norte": 45, // 45 minutos
		"centro":     20, // 20 minutos
	},
}

// OrderStatus representa o status de um pedido
type OrderStatus struct {
	Status    MockStatus
	IsDelayed bool
}

// PaymentStatus representa o status de um pagamento
type PaymentStatus struct {
	IsConfirmed bool
	HasIssues   bool
}

// StockStatus representa o status do estoque
type StockStatus struct {
	Available bool
	ItemID    string
}

// DeliveryEstimate representa uma estimativa de entrega
type DeliveryEstimate struct {
	MinutesEstimate int
	Region          string
}
