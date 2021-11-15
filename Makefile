# we will put our integration testing in this path
INTEGRATION_TEST_PATH?=./test

ENV_LOCAL=\
	APP_DB_USERNAME=postgres\n\
	APP_DB_PASSWORD=123456\n\
	APP_DB_NAME=jumia\n\
	APP_DB_PORT=5432\n\
	APP_DB_HOST=db\n\
	APP_PORT=":3000"
# set of env variables that you need for testing
ENV_LOCAL_TEST=\
	APP_DB_USERNAME=postgres\n\
	APP_DB_PASSWORD=123456\n\
	APP_DB_NAME=jumia_test\n\
	APP_DB_PORT=5432\n\
	APP_DB_HOST=db_test\n\
	APP_PORT=":3000"
	
set.env:
	@touch .env
	@echo "${ENV_LOCAL}" >.env

set.env.test:
	@touch .env
	@echo "${ENV_LOCAL_TEST}" >.env

# this command will start a docker components that we set in docker-compose.yml
docker.up: set.env
	docker-compose up --build;

docker.up.test: set.env.test
	docker-compose -f docker-compose-test.yml up --build;

# shutting down docker components
docker.stop:
	docker-compose down;

docker.stop.test:
	docker-compose -f docker-compose-test.yml down;
# this command will trigger integration test
# INTEGRATION_TEST_SUITE_PATH is used for run specific test in Golang, if it's not specified
# it will run all tests under ./it directory
test.integration:
	go test -tags=integration $(INTEGRATION_TEST_PATH) 

# this command will trigger integration test with verbose mode
test.integration.debug:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -v -run=$(INTEGRATION_TEST_SUITE_PATH)