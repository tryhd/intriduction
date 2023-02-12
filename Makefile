FILENAME := $(filter-out $@,$(MAKECMDGOALS))

build: 
	@go build -o server main.go

run: build
	@./server

watch:
	@go run main.go

migrate:
	@go run configs/command/main.go migrate

migrate-seed:
	@go run configs/command/main.go migrate-seed

model:
	@mkdir -p app/models
	@go run configs/command/main.go $(FILENAME)

controller:
	@mkdir -p app/controllers
	@go run configs/command/main.go $(FILENAME)

service:
	@mkdir -p app/service
	@go run configs/command/main.go $(FILENAME)

repository:
	@mkdir -p app/repositories
	@go run configs/command/main.go $(FILENAME)

dto:
	@mkdir -p app/dtos
	@go run configs/command/main.go $(FILENAME)

all:
	@mkdir -p app/models
	@mkdir -p app/dtos
	@mkdir -p app/repositories
	@mkdir -p app/services
	@mkdir -p app/controllers
	@go run configs/command/main.go $(FILENAME)
