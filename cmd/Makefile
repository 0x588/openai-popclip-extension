.PHONY: deploy
deploy: build copy

.PHONY: build
build:
	CGO_ENABLE=1 GOOS=darwin GOARCH=amd64 go build -o openai-cmd main.go
	CGO_ENABLE=1 GOOS=darwin go build -o openai-cmd-arm64 main.go

.PHONY: copy
copy:
	rm -f ../openAi.popclipext/openai-cmd*
	cp openai-cmd* ../openAi.popclipext/