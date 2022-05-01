run:
	gin -i --port 8080 --appPort 8081 --path . --build ./app run ./app/main.go
build:
	sudo docker image build -t monolith .
remove-useless:
	sudo docker image prune

migration-security:
	migrate create -ext sql -dir handler/security/entity/database/migrations -seq init