docker stop fetch
docker rm fetch
docker rmi fetch
docker build -t fetching-py:latest -f deployment/fetching.Dockerfile .
docker run -itd -p 6000:6000 --name fetch fetching-py:latest
docker logs fetch

docker stop auth
docker rm auth
#docker rmi fetch
docker build -t auth-go:latest -f deployment/auth.Dockerfile .
docker run -itd -p 7000:7000 --name auth auth-go:latest
docker logs auth




docker-compose up --build -dd