export PATH := $(PATH):$(GOPATH)/bin

pre-dev: pull-dev-image
	npm install -g stmux
	go get -v -u github.com/c9s/gomon
	brew install dep -v
	@echo Finnish dev setup, run "sh scripts/dev.bash" to start dev

pull-dev-image:
	docker pull postgres
	docker pull fenglc/pgadmin4

dev:
	@stmux -M -- [ [ "godo server --watch" .. "cd client&&yarn tsc -w" .. "cd client&&yarn serve" ] : "docker-compose -f docker-compose.dev.yaml up" ]
