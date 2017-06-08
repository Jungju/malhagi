test: 
	AUTH_REDIS_URL=localhost:6379 \
	AUTH_REDIS_PASSWORD=dkos \
	SECRET_KEY=superjjgo \
	SYSTEM_ADMIN=system \
	SYSTEM_PASSWORD=t \
	DB_SYNC=true \
	DB_HOME=127.0.0.1 \
	DB_PORT=3306 \
	DB_DATABASE=malhagi_test \
	DB_USER=root \
	DB_PASSWORD=dkos \
	go test -v ./tests/
test_ui:
	AUTH_REDIS_URL=localhost:6379 \
	AUTH_REDIS_PASSWORD=dkos \
	SECRET_KEY=superjjgo \
	SYSTEM_ADMIN=system \
	SYSTEM_PASSWORD=t \
	DB_SYNC=true \
	DB_HOME=127.0.0.1 \
	DB_PORT=3306 \
	DB_DATABASE=malhagi_test \
	DB_USER=root \
	DB_PASSWORD=dkos \
	goconvey
run_dev:
	AUTH_REDIS_URL=localhost:6379 \
	AUTH_REDIS_PASSWORD=dkos \
	SECRET_KEY=superjjgo \
	SYSTEM_ADMIN=system \
	SYSTEM_PASSWORD=t \
	DB_SYNC=true \
	DB_HOME=127.0.0.1 \
	DB_PORT=3306 \
	DB_DATABASE=malhagi \
	DB_USER=root \
	DB_PASSWORD=dkos \
	bee run