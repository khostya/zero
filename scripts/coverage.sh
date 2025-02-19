DEFAULT_PG_URL="postgres://user:password@localhost:5432/postgres?sslmode=disable"
PG_URL=$DEFAULT_PG_URL

go test  ./internal/... ./pkg/...  -coverprofile unit.coverage.txt
TEST_DATABASE_URL=$PG_URL go test ./tests/postgres/... -tags=integration -coverprofile integration.coverage.txt -coverpkg=github.com/khostya/zero/...
echo "mode: set" > coverage.txt && cat *.coverage.txt | grep -v mode: | sort -r | awk '{if($1 != last) {print $0;last=$1}}' >> coverage.txt