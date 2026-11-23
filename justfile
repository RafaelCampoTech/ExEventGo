# ExEventGo — run recipes from this folder (`just` / `just publisher`).

set shell := ["bash", "-euo", "pipefail", "-c"]
set dotenv-filename := "publisher/.env"
set dotenv-load := true

http_port := env_var_or_default("HTTP_PORT", "3000")

# List available recipes.
default:
    @just --list

# Run the publisher HTTP service.
publisher:
    cd publisher && HTTP_PORT={{http_port}} go run .

# Alias for `just publisher`.
run: publisher
