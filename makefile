test: 
	AUTH_REDIS_URL=localhost:6379 \
	AUTH_REDIS_PASSWORD=jjgo \
	SECRET_KEY=superjjgo \
	SYSTEM_ADMIN=system \
	SYSTEM_PASSWORD=t \
	DB_SYNC=true \
	DB_HOME=127.0.0.1 \
	DB_PORT=3306 \
	DB_DATABASE=mhg_test \
	DB_USER=jjgo \
	DB_PASSWORD=jjgo \
	go test -v ./tests/
test_ui:
	AUTH_REDIS_URL=localhost:6379 \
	AUTH_REDIS_PASSWORD=jjgo \
	SECRET_KEY=superjjgo \
	SYSTEM_ADMIN=system \
	SYSTEM_PASSWORD=t \
	DB_SYNC=true \
	DB_HOME=127.0.0.1 \
	DB_PORT=3306 \
	DB_DATABASE=mhg_test \
	DB_USER=jjgo \
	DB_PASSWORD=jjgo \
	goconvey
run_dev:
	AUTH_REDIS_URL=localhost:6379 \
	AUTH_REDIS_PASSWORD=jjgo \
	SECRET_KEY=superjjgo \
	SYSTEM_ADMIN=system \
	SYSTEM_PASSWORD=t \
	DB_SYNC=true \
	DB_HOME=127.0.0.1 \
	DB_PORT=3306 \
	DB_DATABASE=mhg \
	DB_USER=jjgo \
	DB_PASSWORD=jjgo \
	bee run
prepare_dev:
	docker-compose up -d