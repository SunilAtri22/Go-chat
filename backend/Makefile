dockerinit:
	docker run --name chat-app -p 9042:9042 -d cassandra
.PHONY: dockerinit

docker:docker exec -it chat-app cqlsh
	