build:
	docker build -t advert-service-image:1.0.0 .

run:
	docker run -d -p 8081:8081 --name advert-service-container advert-service-image:1.0.0

gen:
	protoc --proto_path=./internal/proto ./internal/proto/*.proto --go_out=plugins=grpc:.
