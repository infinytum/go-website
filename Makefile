.DEFAULT_GOAL=build-docker


build-docker:
	env GOOS=linux env GOARCH=amd64 go build -o output/app-linux_amd64 cmd/server/main.go
	docker build -t docker.k8.network/library/baerenhoehle:latest . || true
	docker push docker.k8.network/library/baerenhoehle:latest || true
	rm -R output