install:
	cp config.example.toml config.toml
run:
	go run config.go database.go user.go main.go --config config.toml
test:
	go test -v

curl-add:
	curl -i -XPOST 127.0.0.1:1234/users -d '{"name":"Greg", "active":true}' -H "Content-Type: application/json"
curl-users:
	curl -i 127.0.0.1:1234/users
	curl -i 127.0.0.1:1234/users?active=0
	curl -i 127.0.0.1:1234/users?active=1
curl-delete:
	curl -i -XDELETE 127.0.0.1:1234/users/96e1a704-d2cf-4d9a-8fcb-c28b6d1a6f16

postgres-run:
	docker run --rm --name apiexample-postgres -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=postgres -e POSTGRES_DB=apiexample -d postgres
	docker inspect apiexample-postgres |grep IPAddress
postgres-login:
	docker exec -it apiexample-postgres psql -U postgres -d apiexample
postgres-stop:
	docker stop apiexample-postgres
