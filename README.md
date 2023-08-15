<p align="center">
 <img src="https://github.com/jeronimo3875br/rikka/blob/master/assets/rikka_main.gif" alt="rikka_main" width="300"/>
</p>

# Intrudução
Rikka é uma CLI desenvolvida para monitoramento e recarregamento automático de aplicativos, similar a soluções como <a href="https://github.com/remy/nodemon">nodemon</a>. No entanto, a Rikka foi concebida de forma a oferecer suporte a qualquer linguagem de programação, eliminando a necessidade de reiniciar manualmente seus aplicativos ou de instalar diferentes pacotes para diversas tecnologias.

# Intalação
Para instalar o Rikka, é necessário ter o GO instalado em sua máquina. Você pode encontrar as instruções de instalação da versão mais recente em: https://go.dev/dl/

Após instalar o GO, você pode utilizar o <a href="https://git-scm.com/">GIT<a/> para baixar o Rikka em sua máquina, executando o seguinte comando em seu terminal: 

```sh 
git clone jeronimo3875br/rikka
```

Certifique-se de ajustar o caminho de acordo com o local desejado em sua máquina.

## Compilar 
Depois de ter a Rikka em sua máquina, você pode compilá-la em um executável usando o comando::

```sh
go build -o rikka main.go
```

Após a compilação, um arquivo executável será gerado, que você pode mover para a pasta de sua preferência.

Recomenda-se adicionar o caminho até o executável e registrá-lo nas variáveis de ambiente **PATH** do seu sistema para facilitar o acesso.

Certifique-se de ajustar o caminho do executável de acordo com a localização do arquivo compilado em sua máquina.

## Como usar
Você pode verificar todos os comandos e flags disponíveis executando e passando a flag **"--help"** ou **"-h"**. Essa flag funciona globalmente e você pode utilizá-la em qualquer lugar da CLI. Por exemplo: 

```sh
./rikka --help
```

 ou se estiver configurado nas variáveis de ambiente, apenas: 
 
 ```sh
 rikka --help
```

## Monitorar arquivos e auto reload
Para começar a monitorar um projeto verificando mudanças e realizando o recarregamento automático da execução, você pode usar o subcomando **watch**. Este comando aceita algumas flags para configuração:

- **"--path"** ou **"-p"**: Define o caminho até o seu projeto, sendo o valor padrão **"./"**, que corresponde ao diretório atual.
- **"--reflect"** ou **"-r"**: Especifica a ação (comando) que será executada após a identificação de uma alteração no projeto **(obrigatório)**.
- **"--ignore"** ou **"-i"**: Permite definir pastas e arquivos que serão ignorados, ou seja, quando alterados, não causarão ação.

A estrutura para usar esse comando é: `rikka watch --path /caminho/do/seu/programa --reflect "comando de recarregamento"`

Aqui estão alguns exemplos de uso:

```sh
# Python
rikka watch --path C:\Users\Jeronimo\projects\web_scraping --reflect "python main.py"

# JavaScript (ignorando pasta "node_modules")
rikka watch --path C:\Users\Jeronimo\projects\api --reflect "node bin/server.js" --ignore node_modules

# Golang (ignorando "go.mod" e "README.md")
rikka watch --path C:\Users\Jeronimo\projects\cli --reflect "go run cmd/main.go" --ignore go.mod,README.md
```

# Licença
Rikka é lançado sob a licença do MIT. Consulte: <a href="https://github.com/jeronimo3875br/rikka/blob/master/LICENSE">Licença</a>
