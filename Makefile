.PHONY: dev buildrun

dev:
	go run cmd/main.go
# === Build and Run ===
build:
	@go build -o bin/ftrchat -buildvcs=false ./cmd/main.go

run: build
	@./bin/ftrchat

create:
		sudo mkdir -p bin
		sudo chmod 7777 bin
reload:
		sudo systemctl daemon-reload

start:
		sudo systemctl start ftrchat

enable:
		sudo systemctl enable ftrchat

stop:
		sudo systemctl stop ftrchat

restart:
		sudo systemctl restart ftrchat

status:
		sudo systemctl status ftrchat

log:
		journalctl -u ftrchat -f

# === Git Pull & Full Install ===
pull:
		sudo git pull

update: pull build restart status log
