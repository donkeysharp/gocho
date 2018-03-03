build-dev:
	@echo "Building gocho"
	go install github.com/donkeysharp/gocho/cmd/gocho

clean:
	rm -rf dist/*

dist: clean ui generate
	@echo "Building gocho for Linux x86_64..."
	go build -o dist/gocho cmd/gocho/gocho.go

generate:
	go generate cmd/gocho/gocho.go

dist-linux32:
	@echo "Building gocho for Linux 32bits..."
	GOOS=linux GOARCH=386 go build -o dist/gocho-linux32 cmd/gocho/gocho.go

dist-win32:
	@echo "Building gocho for Windows 32bits..."
	GOOS=windows GOARCH=386 go build -o dist/gocho-win32.exe cmd/gocho/gocho.go

dist-win64:
	@echo "Building gocho for Windows 64bits..."
	GOOS=windows GOARCH=amd64 go build -o dist/gocho-win.exe cmd/gocho/gocho.go

dist-darwin:
	@echo "Building gocho for Darwin 64bits..."
	GOOS=darwin GOARCH=amd64 go build -o dist/gocho-darwin cmd/gocho/gocho.go

docker: dist
	docker build . -t donkeysharp/gocho

start:
	docker run -it -p "1337:1337" --rm donkeysharp/gocho gocho start --debug || true

test:
	docker run -it --rm donkeysharp/gocho || true

clean-dashboard:
	rm -rf assets/assets_gen.go

ui: clean-dashboard
	cd ui \
	&& yarn build
