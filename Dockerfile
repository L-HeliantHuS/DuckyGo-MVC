FROM golang

ENV GOPROXY=https://goproxy.io
ENV MYSQL_DSN="db_user:db_passwd@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local"
ENV REDIS_ADDR="127.0.0.1:6379"
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV SECRET_KEY="3RH1983RH1938"
ENV plugin=0

ADD . /DuckyGo-MVC

WORKDIR /DuckyGo-MVC

RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOARCH=amd64 go build -o duckygo-mvc


RUN mv duckygo-mvc /usr/bin/duckygo-mvc && \
    chmod +x /usr/bin/duckygo-mvc

EXPOSE 8080

ENTRYPOINT ["duckygo-mvc"]
