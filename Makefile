# genrates the templ html files
html:
	@templ generate

# generates the swagger docs
docs:
	@swag init

# runs the server
run:
	@go run main.go	