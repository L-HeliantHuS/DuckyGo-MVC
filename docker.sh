docker build -t DuckyGo-MVC:release .
docker run -di --name duckygo-mvc -p80:8080 DuckyGo-MVC:release