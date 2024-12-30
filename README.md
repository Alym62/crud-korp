# CRUD - Produtos Korp

<p>
    CRUD de produtos feito com Golang, Gin, PostgreSQL, Docker, Angular e Angular Material. A aplicação possui autenticação via JWT. No backend foi utilizado uma arquitetura limpa. No front foi utilizado as boas práticas do Angular como organização de pastas, o que é compartilhado pela nossa aplicação e assim por diante. Toda a aplicação foi baseado para um caso de uso no qual fosse para a produção. Então vai ser visto dois arquivos de <i>.env_template</i> para rodar a aplicação dentro docker ou localmente.
</p>

### Backend

<p>
    O backend consiste em uma REST API com Golang, Gin, PostgreSQL. Utilizei as boas práticas da comunidade Go e também uma arquitetura limpa. Essa API contém uma autenticação via JWT. Não foi utilizado nenhum ORM como o <i>Gorm</i>.
</p>

### Frontend

<p>
    O frontend foi desenvolvido com Angular e Angular Material para geração de componentes. Utilizei todo o poder que o Angular pode proporcionar com base na sua arquitetura em modulos. Então eu tenho ali um guarda de rotas, interceptor, separei cada componente da minha página em modulos e os serviços, helpers ou abstrações que são compartilhadas por toda a aplicação está na pasta de <i>Shared</i>.
</p>

# Como rodar a aplicação?

<p>
    Para rodar a aplicação é bastante simples basta criar um arquivo .env com base no exemplo que foi deixado e dar um comando <strong><i>docker-compose up --build -d</i></strong>. Com isso você já vai conseguir subir o banco de dados e a API juntos. O frontend basta rodar o comando <strong><i>ng serve -o</i></strong>.
</p>
<p>
    Mas se por acaso quiser subir a aplicação localmente mesmo, basta entrar dentro da pasta backend, criar o arquivo .env com base no exemplo que foi deixado lá também e utilizar o seguintes comandos: <strong><i>go mod tidy</i></strong>. E depois de tudo instalado <strong><i>go run main.go</i></strong>.
</p>

# Stack utilizada nesse projeto

[![My Skills](https://skillicons.dev/icons?i=go,postgresql,docker,angular,ts&perline=3)](https://skillicons.dev)