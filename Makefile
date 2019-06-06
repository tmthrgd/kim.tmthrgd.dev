export GO111MODULE=on

.PHONY: all generate serve deploy deploy-env-var docker-serve docker-stop

all:

generate:
	go generate

serve:
	go run -tags dev .

docker-serve:
	docker build -t kim-tmthrgd-dev-server -f Dockerfile .
	docker run -e GITHUB_TOKEN -p 8090:8090 -e PORT=8090 -d kim-tmthrgd-dev-server

docker-stop:
	docker stop $(shell docker ps -f ancestor=kim-tmthrgd-dev-server --format "{{.ID}}")

deploy:
	gcloud --project kim-tmthrgd-dev builds submit --tag gcr.io/kim-tmthrgd-dev/server
	gcloud --project kim-tmthrgd-dev beta run deploy server --image gcr.io/kim-tmthrgd-dev/server

deploy-env-var:
	gcloud --project kim-tmthrgd-dev beta run services update server --update-env-vars CONSUMER_KEY=${CONSUMER_KEY},CONSUMER_SECRET=${CONSUMER_SECRET}