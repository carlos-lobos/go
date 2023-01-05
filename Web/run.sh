#!/bin/bash

FILE="$1"
if [ ! -f "$FILE" ] ; then
    FILE="$1/main.go"
    if [ ! -f "$FILE" ] ; then
        echo "[USO] $0 FILE.go"
        exit 0
    fi
fi

# Ejecutar (go compila temporalmente para ejecutar el comando)
go run "$FILE"
