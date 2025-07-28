COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html

init:
	go mod init {{.GithubRepository}}
	go mod tidy

build:
	go build -o bin/{{.Name}} .

run: build
	./bin/{{.Name}}

# Limpia archivos generados
clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

# Ejecuta las pruebas y muestra cobertura en la terminal
test:
	go test -v -coverprofile=$(COVERAGE_FILE) ./...

# Genera archivo HTML con detalles de cobertura
coverage: test
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Reporte generado: $(COVERAGE_HTML)"

# Ejecuta pruebas, cobertura y abre el HTML (Linux, macOS, Windows)
open: coverage
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
