COVERAGE_FILE=cover.out
COVERAGE_HTML=cover.html

init:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go mod init github.com/FrancoPersonal/golang-template-example	
	go mod tidy

newtemplate:
	go run src/templates/templatebase/main.go templatenew

# Limpia archivos generados
clean:
	rm -r go.mod
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
	rm -r output/*

# Ejecuta las pruebas y muestra cobertura en la terminal
test:
	go test ./... -cover -v -coverprofile=$(COVERAGE_FILE)

# Genera archivo HTML con detalles de cobertura
coverage: test
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Reporte generado: $(COVERAGE_HTML)"

# Ejecuta pruebas, cobertura y abre el HTML (Linux, macOS, Windows)
showcoverage: coverage
	@# Abrir archivo HTML con el visor predeterminado según el sistema operativo
	@if command -v xdg-open >/dev/null; then \
		xdg-open $(COVERAGE_HTML); \
	elif command -v open >/dev/null; then \
		open $(COVERAGE_HTML); \
	elif command -v start >/dev/null; then \
		start $(COVERAGE_HTML); \
	else \
		powershell.exe -Command "Start-Process '$(COVERAGE_HTML)'" || \
		echo "Abre el archivo $(COVERAGE_HTML) manualmente"; \
	fi

makecoveragefile:
	./coverage.sh