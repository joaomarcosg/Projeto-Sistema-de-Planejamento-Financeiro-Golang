# API RESTful em Golang - Sistema de Planejamento Financeiro

## Projeto proposto pelo desafio de código em Golang da plataforma DIO

### Tecnologias

<div align="left">
  <img src="https://cdn.simpleicons.org/go/00ADD8" height="40" alt="go logo"  />
</div>

|Tecnologia | Descrição                                     |
| --------  | ----------                                    |
| Go        | Linguagem de programação estaticamente tipada |


### Descrição do projeto

O Sistema de planejamento financeiro é uma API RESTful que rastrea transações financeiras. Não foi usado nenhum framework para o seu desenvolvimento, apenas as bibliotecas/pacotes padrões da linguagem.

Esse projeto faz parte do meu aprendizado prático em Go, portanto, é uma aplicação inicial simples. Nela tive a oportunidade de aplicar conceitos como criar **servidor HTTP** e fazer **testes unitários**.

### Funcionalidades

- Listar transações financeiras
- Criar nova transação financeira

### Endpoints

- Listar transações: ```GET/transactions```
- Criar transação: ```POST/transactions/create```

### Estrutura de pastas

```shell
├───adapter               # Código fonte principal da aplicação
│   └───http              # Rotas da API
│       ├───actuator      # Função da rota que dita a saúde da aplicação
│       └───transactions  # Funções de tratamento das requisições das rotas
├───model                 # Modelos de contrato
│   └───transactions      # Modelo de contrato para transações
└───util                  # Função para resposta padrão de tempo e teste unitário da aplicação
```

### Rodando localmente

Você precisa ter instalado [Go](https://go.dev/) em sua máquina. Versão utilizada: ```go1.22.4```

Clone o repositório

```bash
git clone https://github.com/joaomarcosg/Projeto-Sistema-de-Planejamento-Financeiro-Golang.git
```

Inicie o servidor

```bash
go run main.go
```



