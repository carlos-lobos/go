#!/bin/bash

DIR="$1"
if [ ! -d "$DIR" ] ; then
    echo "[USO] $0 DIR (dentro del directorio debe existir el archivo main.go)"
    exit 0
fi

FILE="$DIR/main.go"
if [ ! -f "$FILE" ] ; then
    echo "[ERROR] Debe indicar un directorio donde exista un main.go"
    exit 0
fi

cd $DIR
FILE="main.go"

# Compilar
go build "$FILE"

# Ejecutar
OUT=`basename "$FILE" ".go"`
./${OUT}
