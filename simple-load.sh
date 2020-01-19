#!/bin/bash

# Requires hey https://github.com/rakyll/hey

cleanup ()
{
    kill -s SIGTERM $$
    exit 0
}

trap cleanup SIGINT SIGTERM

PORT=8080

while true; do 
    t=$(date +%s)
    hey http://localhost:${PORT}/ready?t=${t}
    test $? -gt 128 && break
done