build-dev:
	@echo "Building gocho"
	go install github.com/donkeysharp/gocho/cmd/gocho

clean:
	rm -rf dist/*

dist: clean
	@echo "Building gocho..."
	go build -o dist/gocho cmd/gocho/gocho.go

docker: dist
	docker build . -t donkeysharp/gocho

start:
	docker run -it -p "1337:1337" --rm donkeysharp/gocho gocho start --debug || true

test:
	docker run -it --rm donkeysharp/gocho || true
