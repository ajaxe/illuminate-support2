SHELL := pwsh.exe

clean:
	@echo "Cleaning up..."
	@pwsh -Command "if(Test-Path -Path ./dist) { rm -recurse -force ./dist }"

build: clean
	@echo "Building web app..."
	@pwsh -Command "Set-Item Env:GOARCH wasm; Set-Item Env:GOOS js; go build -o ./dist/web/app.wasm ./cmd/webapp"
	@echo "Building server..."
	@pwsh -Command "Set-Item Env:GOARCH amd64; Set-Item Env:GOOS windows; go build -o ./dist/server.exe ./cmd/webapp/"
	@echo "Copying web files..."
	@pwsh -Command "copy ./web/* ./dist/web/ -Recurse -Force"

run: build
	@echo "Running server..."
	cd ./dist && ./server.exe

staticgen: build
	@echo "Building static site generator..."
	@pwsh -Command "go build ./cmd/staticgen"
	@echo "Running static site generator..."
	cd ./cmd/staticgen && ./staticgen.exe