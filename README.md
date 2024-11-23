# kanban

Estrutura de arquivos

+ cmd
    - main.go
+ internal
    + database
        - database.go
    + domain
        - Entidades do domínio da aplicação
    + repository
        - Arquivos com a funções e/ou classes que fazem a comunicação com o banco de dados. Depende do pacate *database*
    + server
        - server.go (Fornece o servidor)
        - routes.go (Organiza as rotas da API)
+ scripts
    + sql
        - 001-database.sql (Criação do banco de dados)
    - Outros scripts utilizados na manutenção do sistema
    

# Referências

* [How To Generate API Blueprint using SwagGo](https://hackernoon.com/how-to-generate-api-blueprint-using-swaggo-gz1s33v7)