# Digital Wallet - Event Driven Architecture

Este projeto foi desenvolvido durante o módulo de **Event Driven Architecture (EDA)** do **Full Cycle 3.0**, demonstrando os princípios e práticas de arquitetura orientada a eventos.

## 🏗️ Arquitetura

O projeto implementa uma **arquitetura event-driven** com dois microserviços principais que se comunicam através de eventos assíncronos usando Apache Kafka:

## 🚀 Funcionalidades

### WalletCore (Porta 8080)

- **Registro de Clientes**: Criação e gerenciamento de usuários da carteira digital
- **Registro de Contas**: Abertura de contas vinculadas aos clientes
- **Transações**: Processamento de transferências entre contas
- **Banco de Dados**: MySQL para persistência dos dados principais
- **Eventos**: Publicação de eventos no Kafka para sincronização

### Balance-MS (Porta 3003)

- **Consumo de Eventos**: Recebe mensagens do Kafka publicadas pelo WalletCore
- **Atualização de Saldos**: Mantém os saldos das contas atualizados em tempo real
- **Banco de Dados**: PostgreSQL para armazenamento dos saldos
- **API HTTP**: Endpoint para consulta de saldos das contas

## 🛠️ Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Banco de Dados**: MySQL (WalletCore) e PostgreSQL (Balance-MS)
- **Message Broker**: Apache Kafka

## 📋 Pré-requisitos

- Docker
- Docker Compose
- Portas disponíveis: 8080, 3003, 3306, 5432, 9092, 9021

## 🚀 Como Executar

### 1. Clone o repositório

```bash
git clone git@github.com:RianMarlon/fc-event-driven-architecture.git
cd fc-event-driven-architecture
```

### 2. Execute o Docker Compose

```bash
docker-compose up -d
```

Este comando irá:

- Iniciar o MySQL para o WalletCore
- Iniciar o PostgreSQL para o Balance-MS
- Iniciar o Zookeeper e Kafka
- Iniciar o Kafka Control Center (porta 9021)
- Executar as migrações dos bancos de dados
- Iniciar as aplicações WalletCore e Balance-MS

### 3. Verifique os serviços

```bash
docker-compose ps
```

### 4. Acesse as aplicações

- **WalletCore API**: http://localhost:8080
- **Balance-MS API**: http://localhost:3003
- **Kafka Control Center**: http://localhost:9021

## 📡 Endpoints da API

### WalletCore (Porta 8080)

#### Criar Cliente

```http
POST /clients
Content-Type: application/json

{
    "name": "Nome do Cliente",
    "email": "cliente@email.com"
}
```

#### Criar Conta

```http
POST /accounts
Content-Type: application/json

{
    "client_id": "uuid-do-cliente"
}
```

#### Criar Transação

```http
POST /transactions
Content-Type: application/json

{
    "account_id_from": "uuid-conta-origem",
    "account_id_to": "uuid-conta-destino",
    "amount": 100.50
}
```

### Balance-MS (Porta 3003)

#### Consultar Saldo

```http
GET /balances/{account_id}
Content-Type: application/json
```

## 🧪 Testando a Aplicação

Use o arquivo `client.http` para testar as APIs:

1. Abra o arquivo `client.http` no VS Code
2. Instale a extensão "REST Client" se ainda não tiver
3. Clique em "Send Request" acima de cada requisição

## 📊 Monitoramento

- **Kafka Control Center**: http://localhost:9021
  - Visualize tópicos, consumidores e mensagens
  - Monitore o fluxo de eventos em tempo real

## 🎯 Princípios EDA Aplicados

1. **Desacoplamento**: Serviços se comunicam através de eventos
2. **Escalabilidade**: Cada serviço pode escalar independentemente
3. **Resiliência**: Falhas em um serviço não afetam outros
4. **Auditoria**: Todos os eventos são persistidos no Kafka
5. **Sincronização Assíncrona**: Atualizações de saldo acontecem de forma assíncrona

## 🚨 Solução de Problemas

### Verificar logs dos serviços

```bash
docker-compose logs walletcore-app
docker-compose logs balance-ms
```

### Reiniciar um serviço específico

```bash
docker-compose restart walletcore-app
```

### Parar todos os serviços

```bash
docker-compose down
```

### Limpar volumes (cuidado: apaga dados)

```bash
docker-compose down -v
```

## 📚 Aprendizados

Este projeto demonstra:

- Como implementar comunicação assíncrona entre serviços
- Padrões de Event Sourcing e CQRS
- Separação de responsabilidades em microserviços
- Uso do Kafka como backbone de eventos
- Implementação de APIs REST em Go
- Gerenciamento de múltiplos bancos de dados

## 🤝 Contribuição

Este projeto foi desenvolvido como parte do curso Full Cycle 3.0, focado em demonstrar os conceitos de Event Driven Architecture.
