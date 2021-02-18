go build server/main.go 
mv main startServer
docker build -t ghcr.io/rakshith-r/http_server:1.0 .
rm startServer