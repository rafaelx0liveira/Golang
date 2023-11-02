### Boas práticas de organização de pastas e arquivos em projetos Go

- 1. Execute o **go mod init nome_do_seu_repositorio** para criar o arquivo go.mod e gerenciar as dependências do projeto
- 2. Crie uma pasta **cmd** para armazenar os arquivos de entrada do seu projeto (main.go)
- 3. O Go não te obriga a ter o nome do arquivo igual ao nome do pacote, mas é uma boa prática
- 4. Crie uma pasta **internal** para armazenar os arquivos que terá tudo o que é interno, inerente e que fundamentalmente será usado apenas dentro da sua aplicação.
     Ex.: Regras de domínio (domain), entidades (entity), etc.
- 5. O nome das pastas devem ser no singular, com letra minúscula e separadas por underline
     Ex.: entity.go, entity_test.go, etc.
- 6. Crie uma pasta **pkg** para armazenar os arquivos que terá tudo o que é externo, que pode ser usado por outras aplicações.
     Ex.: Biblioteca de validação, biblioteca de conexão com banco de dados ou com serviços externos, etc.
- 7. Crie uma pasta **build** para armazenar os arquivos de build da sua aplicação, como Dockerfile, docker-compose.yml, etc.
- 8. Crie uma pasta **configs** sempre que for disponibilizar arquivos de configuração, como arquivos de variáveis de ambiente, arquivos de leitura de variáveis de ambiente, etc. Note que essa pasta não está no singular, porém é uma boa prática.
- 9. Crie uma pasta **docs** para armazenar os arquivos de documentação do seu projeto, como arquivos de documentação de APIs, documentações adicionais, etc. Também não está no singular, mas é uma boa prática.
- 10. Crie uma pasta **api** para armazenar os arquivos de documentação de APIs. **Importante:** Essa pasta não conterá a implementação das APIs (webserver, rotas, etc.), apenas a documentação das APIs, como o arquivo swagger.yml, por exemplo.
- 11. Crie uma pasta **scripts** para armazenar os arquivos de scripts, como scripts de banco de dados, scripts de deploy, makefiles, etc.
- 12. Crie uma pasta **test** para armazenar os arquivos de testes, como arquivos de testes unitários, testes de integração, etc.
- 13. Crie uma pasta **examples** para armazenar os arquivos de exemplos, como arquivos de exemplos de uso da sua biblioteca, por exemplo.
- 14. Crie uma pasta **web** para armazenar os arquivos de imagens (assets) da sua aplicação.
- 15. Crie uma pasta **website** para colocar os dados do site do seu projeto se você não estiver usando as páginas do GitHub.

Fonte:

https://github.com/golang-standards/project-layout

https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md
