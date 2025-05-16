include .env
export

# START::APP COMMANDS

# Executa os testes automatizados no modo verboso
test: 
	go test -v ./...

# Executa os testes automatizados no modo verboso e cria um .html de coverage
test-cov: 
	go test -coverprofile cover.out ./... -v && go tool cover -html=cover.out -o coverage.html

# Inicializa / sobe / coloca pra rodar a API
run:
	go run ./cmd/server/
	
# Faz o build da aplicação e cria o arquivo em /build/server/
build: 
	go build -o ./build/server ./cmd/server

# END::APP COMMANDS
# START::DOCKER-COMPOSE COMMANDS

# Sobe os containers com base no docker-compose.yml
docker-compose-up:
	sudo docker compose up -d

# Derruba	 os containers com base no docker-compose.yml
docker-compose-down:
	sudo docker compose down

# Mostra os logs em tempo real dos containers definidos no docker-compose.yml
docker-compose-logs:
	sudo docker compose logs -f

# END::DOCKER-COMPOSE COMMANDS

