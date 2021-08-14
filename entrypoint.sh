APP_ENV=${APP_ENV}
echo "[$(date)] Running entrypoint script in the '${APP_ENV}' environment..."
CONFIG_FILE=./config/${APP_ENV}.yml
echo "[$(date)] Config file destination '${CONFIG_FILE}'"
if [[ -z ${DATABASE_DNS} ]]; then
  export DATABASE_DNS=$(sed -n 's/^migration_db_dns:[[:space:]]*"\(.*\)"/\1/p' "${CONFIG_FILE}")
fi
echo "[$(date)] Running migrations with '${DATABASE_DNS}'"
migrate -path migrations -database "${DATABASE_DNS}" -path ./migrations -verbose up
echo "[$(date)] Starting server..."
./server -config "${CONFIG_FILE}"
