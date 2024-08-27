build:
	docker image build -t backend .

remove-useless:
	docker image prune

create-migration:
	@read -p "input the migration name: " name; \
	migrate create -ext sql -dir migrations $$name

run-dependencies:
	docker compose --profile dependencies up --force-recreate

run-serivce:
	docker compose --profile dependencies --profile backend-migration --profile backend-service up --abort-on-container-exit --force-recreate --exit-code-from backend-service

run-integration-testing:
	docker compose --profile dependencies --profile backend-migration --profile backend-service --profile integration-tests up --abort-on-container-exit --renew-anon-volumes --force-recreate --exit-code-from integration-tests