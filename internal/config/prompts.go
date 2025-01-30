package config

import (
	"strings"
)

// Prompts contém todos os prompts de contexto usados na aplicação
type Prompts struct {
	BaseContext string
}

// Itens representa um produto com nome, valor e descrição
type Itens struct {
	Nome      string
	Valor     float64
	Descricao string
}

// itens representa os produtos disponíveis na pizzaria
var itens = []Itens{
	{Nome: "Pizza Margherita (Broto)", Valor: 20.00, Descricao: "Tamanho: Aproximadamente 25cm\nDescrição: Pizza clássica com molho de tomate, mussarela e manjericão."},
	{Nome: "Pizza Margherita (Grande)", Valor: 35.00, Descricao: "Tamanho: Aproximadamente 35cm\nDescrição: Pizza clássica com molho de tomate, mussarela e manjericão."},
	{Nome: "Pizza Calabresa (Broto)", Valor: 25.00, Descricao: "Tamanho: Aproximadamente 25cm\nDescrição: Pizza com calabresa fatiada, cebola e azeitonas."},
	{Nome: "Pizza Calabresa (Grande)", Valor: 45.00, Descricao: "Tamanho: Aproximadamente 35cm\nDescrição: Pizza com calabresa fatiada, cebola e azeitonas."},
	{Nome: "Pizza Quatro Queijos (Broto)", Valor: 22.00, Descricao: "Tamanho: Aproximadamente 25cm\nDescrição: Pizza com mussarela, parmesão, gorgonzola e catupiry."},
	{Nome: "Pizza Quatro Queijos (Grande)", Valor: 40.00, Descricao: "Tamanho: Aproximadamente 35cm\nDescrição: Pizza com mussarela, parmesão, gorgonzola e catupiry."},
	{Nome: "Pizza Portuguesa (Broto)", Valor: 24.00, Descricao: "Tamanho: Aproximadamente 25cm\nDescrição: Pizza com presunto, ovos, cebola, ervilha e azeitonas."},
	{Nome: "Pizza Portuguesa (Grande)", Valor: 42.00, Descricao: "Tamanho: Aproximadamente 35cm\nDescrição: Pizza com presunto, ovos, cebola, ervilha e azeitonas."},
	{Nome: "Refrigerante 300ml", Valor: 5.00, Descricao: "Tamanho: 300ml\nDescrição: Refrigerante gelado de sua preferência."},
	{Nome: "Suco Natural 300ml", Valor: 6.00, Descricao: "Tamanho: 300ml\nDescrição: Suco natural fresco de frutas variadas."},
	{Nome: "Bebida Cerveja (Lata)", Valor: 8.00, Descricao: "Tamanho: 350ml\nDescrição: Cerveja gelada de sua preferência."},
}

// DefaultPrompts retorna os prompts padrão da aplicação
func DefaultPrompts() *Prompts {
	produtos := []string{
		"Pizza Margherita (Broto) - R$ 20.00\nTamanho: Aproximadamente 25cm\nDescrição: Pizza clássica com molho de tomate, mussarela e manjericão.",
		"Pizza Margherita (Grande) - R$ 35.00\nTamanho: Aproximadamente 35cm\nDescrição: Pizza clássica com molho de tomate, mussarela e manjericão.",
		"Pizza Calabresa (Broto) - R$ 25.00\nTamanho: Aproximadamente 25cm\nDescrição: Pizza com calabresa fatiada, cebola e azeitonas.",
		"Pizza Calabresa (Grande) - R$ 45.00\nTamanho: Aproximadamente 35cm\nDescrição: Pizza com calabresa fatiada, cebola e azeitonas.",
		"Pizza Quatro Queijos (Broto) - R$ 22.00\nTamanho: Aproximadamente 25cm\nDescrição: Pizza com mussarela, parmesão, gorgonzola e catupiry.",
		"Pizza Quatro Queijos (Grande) - R$ 40.00\nTamanho: Aproximadamente 35cm\nDescrição: Pizza com mussarela, parmesão, gorgonzola e catupiry.",
		"Pizza Portuguesa (Broto) - R$ 24.00\nTamanho: Aproximadamente 25cm\nDescrição: Pizza com presunto, ovos, cebola, ervilha e azeitonas.",
		"Pizza Portuguesa (Grande) - R$ 42.00\nTamanho: Aproximadamente 35cm\nDescrição: Pizza com presunto, ovos, cebola, ervilha e azeitonas.",
		"Refrigerante 300ml - R$ 5.00\nTamanho: 300ml\nDescrição: Refrigerante gelado de sua preferência.",
		"Suco Natural 300ml - R$ 6.00\nTamanho: 300ml\nDescrição: Suco natural fresco de frutas variadas.",
		"Bebida Cerveja (Lata) - R$ 8.00\nTamanho: 350ml\nDescrição: Cerveja gelada de sua preferência.",
	}

	return &Prompts{
		BaseContext: `Sobre a Pizzaria:
- Pizzaria tradicional e artesanal
- Pizzas em dois tamanhos: broto (25cm) e grande (35cm)
- Bebidas geladas disponíveis
- Ambiente familiar e acolhedor

Produtos:
- ` + strings.Join(produtos, "\n- "),
	}
}
