build:
	sudo docker image build -t laundry .
remove-useless:
	sudo docker image prune

migration-security:
	@read -p "input the migration name: " name; \
	migrate create -ext sql -dir handler/security/entity/database/migrations $$name