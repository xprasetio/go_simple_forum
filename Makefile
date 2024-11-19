export MYSQL_URL="mysql://root:admin@tcp(localhost:3307)/go_simple_forum"

create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)
up:
	@ migrate -database $(MYSQL_URL) -path scripts/migrations up
down:
	@ migrate -database $(MYSQL_URL) -path scripts/migrations down