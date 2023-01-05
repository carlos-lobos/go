#!/bin/bash

DIR="$1"

# Si ya existe => Error
if [ -d "$DIR" ] ; then
    echo "[ERROR] El directorio \"$DIR\" ya existe."
    exit 1
fi

# Creo el directorio
mkdir "$DIR"
if [ ! -d "$DIR" ] ; then
    echo "[ERROR] No se pudo crear el directorio \"$DIR\""
    exit 1
fi

# Creo el archivo main.go
FILE="$DIR/main.go"
touch "$FILE"

if [ "$TERM_PROGRAM" = "vscode" ] ; then
    code -g "$FILE"
fi
