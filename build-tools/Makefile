# Environment file and variables
ENV_FILE := .env

SERVICE_NAME := user-service
SYSTEMD_FILE := $(SERVICE_NAME).service
TARGET_PATH := /etc/systemd/system/$(SYSTEMD_FILE)

## systemd_setup: Installs, enables and starts the systemd service
systemd_setup: install_systemd_file reload_systemd enable_and_start_service status_logs

## install_systemd_file: Copies the service file to systemd directory
install_systemd_file:
	@echo "📁 Installing $(SYSTEMD_FILE) to systemd..."
	sudo cp $(SYSTEMD_FILE) $(TARGET_PATH)
	@echo "✅ Copied to $(TARGET_PATH)"

## reload_systemd: Reloads systemd daemon
reload_systemd:
	@echo "🔁 Reloading systemd daemon..."
	sudo systemctl daemon-reload

## enable_and_start_service: Enables and starts the service
enable_and_start_service:
	@echo "🚀 Enabling and starting $(SERVICE_NAME)..."
	sudo systemctl enable $(SERVICE_NAME)
	sudo systemctl restart $(SERVICE_NAME)

## status_logs: Shows systemd status and tails logs
status_logs:
	@echo "📋 Checking service status:"
	sudo systemctl status $(SERVICE_NAME) --no-pager
	@echo "📜 Tailing logs (press Ctrl+C to stop):"
	sudo journalctl -u $(SERVICE_NAME) -f

## build_user_service: Builds the user-service binary and checks logs
build_user_service:
	@echo "1- 🚀 Building user-service binary..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd ../user-service && $$GOFULLPATH build -o $$USER_SERVICE_BINARY ./cmd/api
	@echo "✅ Done! user-service binary built."

	@echo "📋 Checking if $(SERVICE_NAME) is active..."
	@if sudo systemctl is-active --quiet $(SERVICE_NAME); then \
		echo "✅ $(SERVICE_NAME) is already running. Showing logs..."; \
	else \
		echo "⚠️  $(SERVICE_NAME) is not running. Starting it now..."; \
		sudo systemctl start $(SERVICE_NAME); \
	fi

	@echo "📜 Tailing logs for $(SERVICE_NAME)... (Ctrl+C to stop)"
	@sudo journalctl -u $(SERVICE_NAME) -f

## logs: Tails the logs of the service
logs:
	@echo "📜 Tailing logs for $(SERVICE_NAME)... (Ctrl+C to stop)"
	sudo journalctl -u $(SERVICE_NAME) -f	
