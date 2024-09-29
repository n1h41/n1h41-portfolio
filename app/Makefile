build:
	npx tailwindcss -i views/css/styles.css -o public/css/styles.css
	@templ generate views
	@go build -o bin/app main.go

test:
	@go test -v ./...

go:
	@air --build.cmd "go build -o build/app cmd/main.go" --build.bin "build/app"

templ:
	@templ generate --watch --proxy="http://localhost:80" --open-browser=false

tailwind:
	@npx tailwindcss -i views/css/styles.css -o public/css/styles.css --watch
