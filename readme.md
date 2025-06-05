# ResolutionAPI

A ResolutionAPI é uma API RESTful desenvolvida em Go, projetada para ser a espinha dorsal de uma plataforma colaborativa. Seu propósito é permitir que empresas, prefeituras e cidadãos reportem problemas em suas comunidades ou organizações, promovendo a participação ativa na proposição de soluções eficazes. O objetivo principal é fomentar o engajamento cívico e viabilizar a resolução ágil e estruturada de desafios urbanos e corporativos.

A arquitetura adotada segue os princípios da Clean Architecture, priorizando a separação clara de responsabilidades, alta testabilidade, facilidade de manutenção e escalabilidade do sistema.

**💡 Por que essa arquitetura?**


Este projeto está sendo desenvolvido paralelamente ao meu aprendizado em Golang. Em um primeiro olhar, o código pode parecer mais verboso e até fugir um pouco da abordagem "Go-like", conhecida por ser simples, direta e pragmática. No entanto, optei por uma estrutura mais robusta e modular justamente para facilitar a escalabilidade e a manutenção, aspectos importantes durante um processo contínuo de aprendizado e evolução do projeto.

Além disso, está nos planos o desenvolvimento de um aplicativo mobile que consumirá esta API. Pensando nisso, optei por uma arquitetura que facilite a adição ou modificação de endpoints conforme novas demandas surjam ao longo do desenvolvimento do app. Dessa forma, o sistema estará preparado para crescer de forma organizada e sustentável.

## 🚀 Funcionalidades Principais

A ResolutionAPI permite aos usuários gerenciar problemas e soluções de forma eficiente e segura. As principais funcionalidades incluem:

* **Gestão de Usuários:**
    * CRUD de perfis de usuário (criação, leitura, atualização e deleção).
    * Autenticação de Usuário via Token JWT

* **Gestão de Problemas:**
    * Criação, leitura, atualização e deleção de problemas.
    * Associação setores a problemas.
    * Atualização de status do problema (Aberto, Em Análise, Resolvido).

* **Gestão de Soluções:**
    * Proposição de soluções para problemas existentes.
    * Aprovação de soluções.
    * Estimativa de custo para soluções.

* **Interações:**
    * Reações (Like, Deslike...) a soluções propostas.

* **Setores:**
    * Gerenciamento e categorização de problemas por setores (ex: Saneamento, Trânsito, Meio Ambiente).

## Tecnologias Utilizadas
- Golang
- PostgreSQL

## Regras de Uso

Para usar ResolutionAPI, siga os passos abaixo:

1. Clone este repositório
```
git clone https://github.com/italomsilva/go-resolution-api.git

```

2. Instale as dependências
```
cd go-resolution-api
go mod tidy
```

3. Configure as variáveis de ambiente:
- Crie um arquivo .env na raiz do projeto
- Adicione as variáveis com base
```
DB_HOST = localhost #(seu host)
DB_PORT = 5432 #(porta onde roda o postgres)
DB_USER = user #(seu usuario do banco de dados)
DB_PASSWORD = pass1234 #(sua senha do banco de dados)
DB_NAME = resolutionapi #(nome do seu banco de dados)
JWT_SECRET = secret #(sua escolha de palavra-segredo para o token jwt)
API_KEY_VALUE = passkey #(sua escolha de chave de acesso)
```

4. Crie o banco de dados

Execute o arquivo ./database/create.sql para criar as tabelas necessárias

5. Inicie a aplicação:

```
go run main.go
```

6. Acesse a API

A API estará disponível em http://localhost:3060. Utilize ferramentas como Postman ou Insomnia para interagir com a API.
Em toda requisição adicione dois headers: um com o nome 'go-api-key' e o valor definido no arquivo.env, e outro chamado 'req-token' e adicione 'Bearer ' antes do valor do token que você receberá ao se registrar ou fazer login. ou seja 'Bearer **_token_**'


## Licença

Este projeto está licenciado sob a Licença MIT.