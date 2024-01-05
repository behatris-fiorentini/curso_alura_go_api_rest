# Iniciando o projeto

* Incialize o seu módulo, normalmente, utilizamos o repositório do git onde a aplicação será versionada

        go mod init github.com/meu_git/nome_do_meu_repositorio

* Em seguida, de o comando para sincronizar as dependências do código:

        go mod tidy


## Bando de dados

* Utilizando o postgres via docker;
* Para rodá-lo precisa ter o docker instalado (<a href="https://www.docker.com/products/docker-desktop/">Download</a>);
* Para criar e executar os containers do PostgreSQL e do PgAdmin:

        docker-compose up -d --build

### Acessando o gerenciador do Postgres
* Para acessar o SGBD PgAdmin, basta acessar no seu browser <a href="http://localhost:54321/login?next=%2F">http://localhost:54321/login?next=%2F</a>;
* As credenciais são as definidas no arquivo _docker-compose.yml_:

        Usuário: email@email.com
        Senha: 123456

* Para criar o seu banco de dados, basta clicar em "Adicionar Novo Servidor".



<hr>

# Gorilla mux

## Conceito
* O pacote gorilla é um roteador de solicitação;
* Um despachante para combinar as solicitações recebidas com seu respectivo manipulador/controlador;
* O nome mux significa "multiplexador de solicitação HTTP";
* Como o padrão http.ServeMux, mux.Routercompara as solicitações recebidas com uma lista de rotas registradas e chama um manipulador para a rota que corresponde ao URL ou outras condições.
* Documentação: <a href="https://github.com/gorilla/mux">Link.</a>

## Instalação no projeto

        go get -u github.com/gorilla/mux

<hr>

# Gorm

* Para trabalhar com banco de dados, estamos utilizando a biblioteca Gorm;
* Ele é uma biblioteca ORM para Golang, trabalhando com MySQL, PostgreSQL, SQLite, etc.
* Documentação: <a href="https://gorm.io/index.html">Link</a>

## instalação no projeto

        go get -u  gorm.io/gorm 

        go get -u  gorm.io/driver/postgres

<hr>

# CORS
* Para integrar a aplicação Golang com um front-end, possivelmente, teremos problemas de CORS;
* Cross-Origin Resource Sharing (Compartilhamento de recursos com origens diferentes);
* É um mecanismo que usa cabeçalhos adicionais HTTP para informar a um navegador que permita que um aplicativo Web seja executado em uma origem (domínio) com permissão para acessar recursos selecionados de um servidor em uma origem distinta. 
* Um aplicativo Web executa uma requisição cross-origin HTTP ao solicitar um recurso que tenha uma origem diferente (domínio, protocolo e porta) da sua própria origem.
* Para que não tenhamos problemas de acesso do fronmt para a nossa API em Golang, utilizou-se o gorilla handlers;


## Gorilla Handlers

* Documentação: <a href="https://github.com/gorilla/handlers">Link</a>;
* Para importa-lo no projeto:

         go get github.com/gorilla/handlers
        
* O exemplo abaixo libera o acesso para dominios especificos:

       http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000", "http://teste.com"}))(r))

* O exemplo abaixo, libera acesso para qualquer origem:

        http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))