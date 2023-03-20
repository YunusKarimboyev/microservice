run:
	go run cmd/main.go

proto-gen:
	./script/gen-proto.sh

migrate_up:
	migrate -path migrations/ -database postgres://yunus:yunus@localhost:5432/user_db up

migrate_down:
	migrate -path migrations/ -database postgres://yunus:yunus@localhost:5432/user_db down

migrate_force:
	migrate -path migrations/ -database postgres://yunus:yunus@localhost:5432/user_db force 1


