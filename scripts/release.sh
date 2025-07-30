#!/bin/bash

set -e

if [ -z "$1" ]; then
  echo "Uso: ./release.sh <version>"
  echo "Ejemplo: ./release.sh v1.0.0"
  exit 1
fi

VERSION=$1

echo "Creando tag $VERSION..."
git tag "$VERSION"

echo "Haciendo push del tag..."
git push origin "$VERSION"

echo "Tag $VERSION enviado. El workflow de release se activará automáticamente."