# go-env-access

A simple package to load environment variables from a specific .env.\* file

## USAGE

When running your main file, export the `env` variable to specify which .env
file to load.

Example:
`$ export env=dev && go run main.go` (loads `.env.dev` file)
