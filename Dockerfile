FROM golang:1.17-alpine

WORKDIR /app/go-contact-book

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080 8080

RUN go build -o ./serve ContactBook/cmd/serve

CMD ["./serve"]

#VOLUME ["/app/go-contact-book/path_to_db"]

#CMD ["/app/go-contact-book/serve"]

#docker images списко образов
#docker ps  список работающих контейнеров
#docker ps -a  список всех контейнеров

#docker run image_id создает контейнер с образа и заускает его
# --rm удалит контейнер после остановки
# -d запускает в фоне
# -p 12345:8082 Локал айпи 12345, докер айпи 8082.
# --name name у контейнера будет указаное имя
# -v db:/app/go-contact-book  сделает папку db которя находится в /app/go-contact-book volume

#docker start/stop conteiner_id запускает/остонавливает контейнер
#docker conteiner prune  - удалить выключеные контейнеры
#docker image prune  - удалить все не использованые образы


#docker build -t sprow/contact-book:v1.0 .

#docker login
#docker push sprow/my-docker/contact-book:v1.0.0


#запускаем имедж который создаёт и запускает контейнер.
# --rm удалит контейнер после остановки
# my-docker/contact-book - задает REPOSITORY
# v1.0 - задает TAG
#docker run -p 12345:8082 --name contact-book --rm sprow/contact-book:v1.0


# меняем теги
#docker tag my-docker/contact-book:v1.0.0 sprow/contact-book:v1.0



#docker run --name=contactbook-db -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d --rm postgres
#migrate create -ext sql -dir ./schema -seq init
