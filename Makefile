VERSION = 0.1.0

build-dev:
	@echo "Building gocho"
	go install github.com/donkeysharp/gocho/cmd/gocho

clean:
	rm -rf dist/*

dist: clean ui generate
	@echo "Building gocho for Linux x86_64..."
	mkdir -p dist/linux64
	go build -o dist/linux64/gocho cmd/gocho/gocho.go
	@zip -j dist/gocho_${VERSION}_linux64.zip dist/linux64/gocho

generate:
	go generate cmd/gocho/gocho.go

dist-linux32:
	@echo "Building gocho for Linux 32bits..."
	mkdir -p dist/linux386
	GOOS=linux GOARCH=386 go build -o dist/linux386/gocho cmd/gocho/gocho.go
	@zip -j dist/gocho_${VERSION}_linux386.zip dist/linux386/gocho

dist-win32:
	@echo "Building gocho for Windows 32bits..."
	mkdir -p dist/win32
	GOOS=windows GOARCH=386 go build -o dist/win32/gocho.exe cmd/gocho/gocho.go
	@zip -j dist/gocho_${VERSION}_win32.zip dist/win32/gocho.exe

dist-win64:
	@echo "Building gocho for Windows 64bits..."
	mkdir -p dist/win64
	GOOS=windows GOARCH=amd64 go build -o dist/win64/gocho.exe cmd/gocho/gocho.go
	@zip -j dist/gocho_${VERSION}_win64.zip dist/win64/gocho.exe

dist-darwin:
	@echo "Building gocho for Darwin 64bits..."
	mkdir -p dist/darwin
	GOOS=darwin GOARCH=amd64 go build -o dist/darwin/gocho cmd/gocho/gocho.go
	@zip -j dist/gocho_${VERSION}_darwin.zip dist/darwin/gocho

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
