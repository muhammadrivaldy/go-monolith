build:
	sudo docker image build -t umkm .
remove-useless:
	sudo docker image prune
migration-security:
	migrate create -ext sql -dir handler/security/entity/database/migrations -seq init