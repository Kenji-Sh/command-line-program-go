# Command Line Application

### Um programa simples CLI utilizando o pacote cli com padrôes de Commit

---

## Descrição do projeto:

Esse projeto tem como objetivo desenvolver uma aplicação simples CLI utilizando Go com o pacote cli,
aplicando o conhecimento básico que desenvolvi em poucos dias, como pointers, slices e structs.

## Pacotes utilizados no Projeto:

  ```
  [x] cli - https://github.com/urfave/cli
  ```

## Instalação do projeto:

1. Certifique-se de ter instalado:

  ```
  Go
  ```

2. Clone este repositório:

  ```bash
  git clone https://github.com/Kenji-Sh/command-line-program-go.git
  ```

3. Navegue até o diretório do projeto:

  ```bash
  cd command-line-program-go
  ```
   
4. Instale os pacotes:

  ```bash
  go get .
  ```

## Como utilizar o projeto:

1. Execute o seguinte comando para criar o programa:

  ```bash
  go build
  ```

2. O programa possui dois comandos: ip e server:
   
- O comando ip ira buscar os Ips públicos do endereço que for mandado na flag host:
  ```bash
  .\command-line.exe ip --host github.com/Kenji-Sh
  ```

- O comando server ira buscar os nomes dos servidores do endereço que for mandado na flag host:
  ```bash
  .\command-line.exe server --host github.com/Kenji-Sh
  ```
