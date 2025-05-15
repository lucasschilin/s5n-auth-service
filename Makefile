include .env
export


test: 
	# Executa os testes automatizados no modo verboso
	go test -v ./...

test-cov: 
	# Executa os testes automatizados no modo verboso e cria um .html de coverage
	go test -coverprofile cover.out ./... -v && go tool cover -html=cover.out -o coverage.html

run:
	# Inicializa / sobe / coloca pra rodar a API
	go run ./cmd/server/
	
build: 
	# Faz o build da aplicação e cria o arquivo em /build/server/
	go build -o ./build/server ./cmd/server