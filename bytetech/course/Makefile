gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server	--type RPC -module bytetech/course/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC -module bytetech/course/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto
test-echo:
	curl http://127.0.0.1:8888/echo
.PHONY: gen-demo-thrift, gen-demo-proto

.PHONY: gen-frontend
gen-frontend: # make gen-frontend page=home
	@cd app/frontend && cwgo server --type HTTP --module bytetech/course/app/frontend --service frontend -I ../../idl --idl ../../idl/frontend/${page}.proto


.PHONY: gen-frontend-auth-page
gen-frontend-auth-page:
	@cd app/frontend && cwgo server --type HTTP --module bytetech/course/app/frontend --service frontend -I ../../idl --idl ../../idl/frontend/auth_page.proto

.PHONY: gen-user
gen-user: # --pass 跳过客户端的代码，客户端全部放到上面的 rpc_gen 目录下。
	@cd rpc_gen && cwgo client  --type RPC --module bytetech/course/rpc_gen --service user -I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --module bytetech/course/app/user --service user --pass "-use bytetech/course/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-product
gen-product: # --pass 跳过客户端的代码，客户端全部放到上面的 rpc_gen 目录下。
	@cd rpc_gen && cwgo client  --type RPC --module bytetech/course/rpc_gen --service product -I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --module bytetech/course/app/product --service product --pass "-use bytetech/course/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client  --type RPC --module bytetech/course/rpc_gen --service cart -I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --module bytetech/course/app/cart --service cart --pass "-use bytetech/course/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto

.PHONY: gen-rpc
gen-rpc: # gen-rpc svc=cart
	@cd rpc_gen && cwgo client  --type RPC --module bytetech/course/rpc_gen --service ${svc} -I ../idl --idl ../idl/${svc}.proto
	@cd app/${svc} && cwgo server --type RPC --module bytetech/course/app/${svc} --service ${svc} --pass "-use bytetech/course/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/${svc}.proto

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=user
	@cd rpc_gen && cwgo client -I ../idl --type RPC --module bytetech/course/rpc_gen --service ${svc} --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: # --pass 跳过客户端的代码，客户端全部放到上面的 rpc_gen 目录下。
	@cd app/user && cwgo server --type RPC --module bytetech/course/app/user --service user --pass "-use bytetech/course/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: run-server
run-server: # run server of {svc}. example: make run-server svc=user
	@cd app/${svc} && go run .

##@ build docker image
.PHONY: build-frontend
build-frontend: # make build-frontend v=v1.0.0
	docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .

.PHONY: build-svc
build-svc: # make build-svc svc=user v=v1.0.0
	docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .

.PHONY: build-all
build-all: # make build-all v=v1.0.0
	docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .
	docker build -f ./deploy/Dockerfile.svc -t cart:${v} --build-arg SVC=cart .
	docker build -f ./deploy/Dockerfile.svc -t checkout:${v} --build-arg SVC=checkout .
	docker build -f ./deploy/Dockerfile.svc -t email:${v} --build-arg SVC=email .
	docker build -f ./deploy/Dockerfile.svc -t order:${v} --build-arg SVC=order .
	docker build -f ./deploy/Dockerfile.svc -t payment:${v} --build-arg SVC=payment .
	docker build -f ./deploy/Dockerfile.svc -t product:${v} --build-arg SVC=product .
	docker build -f ./deploy/Dockerfile.svc -t user:${v} --build-arg SVC=user .

.PHONY: start-docker-frontend
start-docker-frontend: # make start-docker-frontend v=v1.0.0
	docker run -d --name frontend --network course_default --env-file=deploy/.env.example -p 8080:8080 frontend:${v}
.PHONY: start-docker-svc
start-docker-svc: # make start-docker-svc svc=user v=v1.0.0
	docker run -d --name ${svc} --network course_default --env-file=deploy/.env.example ${svc}:${v}
.PHONY: start-docker-all
start-docker-all: # make start-docker-all v=v1.0.0
	docker run -d --name frontend --network course_default --env-file=deploy/.env.example -p 8080:8080 frontend:${v}
	docker run -d --name cart --network course_default --env-file=deploy/.env.example cart:${v}
	docker run -d --name checkout --network course_default --env-file=deploy/.env.example checkout:${v}
	docker run -d --name email --network course_default --env-file=deploy/.env.example email:${v}
	docker run -d --name order --network course_default --env-file=deploy/.env.example order:${v}
	docker run -d --name payment --network course_default --env-file=deploy/.env.example payment:${v}
	docker run -d --name product --network course_default --env-file=deploy/.env.example product:${v}
	docker run -d --name user --network course_default --env-file=deploy/.env.example user:${v}


# kubernetes
.PHONY: gomall-k8s
gomall-k8s:
	kind create cluster --config=deploy/gomall-dev-cluster.yaml

.PHONY: k8s-load-all-images
k8s-load-all-images:
	kind load docker-image --name gomall-dev mysql:8
	kind load docker-image --name gomall-dev redis:7.2.4-alpine
	kind load docker-image --name gomall-dev consul:1.15.4
	kind load docker-image --name gomall-dev jaegertracing/all-in-one:latest
	kind load docker-image --name gomall-dev prom/prometheus:latest
	kind load docker-image --name gomall-dev grafana/grafana:latest
	kind load docker-image --name gomall-dev grafana/loki:2.9.2
	kind load docker-image --name gomall-dev nats:2.10.24-alpine
	kind load docker-image --name gomall-dev frontend:v1.0.0
	kind load docker-image --name gomall-dev cart:v1.0.0
	kind load docker-image --name gomall-dev checkout:v1.0.0
	kind load docker-image --name gomall-dev email:v1.0.0	
	kind load docker-image --name gomall-dev order:v1.0.0
	kind load docker-image --name gomall-dev payment:v1.0.0
	kind load docker-image --name gomall-dev product:v1.0.0
	kind load docker-image --name gomall-dev user:v1.0.0

##@ Open Browser

.PHONY: open.gomall
open-gomall: ## open `gomall` website in the default browser
	@open "http://localhost:8080/"

.PHONY: open.consul
open-consul: ## open `consul ui` in the default browser
	@open "http://localhost:8500/ui/"

.PHONY: open.jaeger
open-jaeger: ## open `jaeger ui` in the default browser
	@open "http://localhost:16686/search"

.PHONY: open.prometheus
open-prometheus: ## open `prometheus ui` in the default browser
	@open "http://localhost:9090"