#!/bin/bash
cd ..
COVERAGE_FILE="coverage.out"
THRESHOLD=80.0

# Asegúrate de que el archivo coverage.out exista
if [ ! -f "$COVERAGE_FILE" ]; then
  echo "❌ Archivo $COVERAGE_FILE no encontrado"
  exit 1
fi

# Obtener porcentaje de cobertura
COVERAGE=$(go tool cover -func="$COVERAGE_FILE" | grep total | awk '{print substr($3, 1, length($3)-1)}')
COVERAGE=90
# Comparar con el umbral
result=$(echo "$COVERAGE >= $THRESHOLD" )
echo $result

if ( $COVERAGE >= $THRESHOLD ); then
  echo "✅ Code coverage is $COVERAGE% (threshold: $THRESHOLD%)"
  exit 0
else
  echo "❌ Code coverage is $COVERAGE% (required: > $THRESHOLD%)"
  exit 1
fi