envs-export: 
	@export $(grep -v '^#' .env | xargs)

run:
	@go run main.go

clean-cache:
	@go clean -modcache

deps:
	@go mod tidy

gen:
	@templ generate

# requires install swaggo with: go install github.com/a-h/templ/cmd/templ@latest
# then export to path with export PATH=$PATH:$(go env GOPATH)/bin 
# and reload terminal/profile ex.: source ~/.bashrc or source ~/.zshrc
swag:
	@swag init -g api/api.go -o api/docs --parseDependency true --parseInternal true --parseDepth 1