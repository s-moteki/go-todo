FROM golang:latest

COPY app /go/src/github.com/s-moteki/go-todo/app
COPY envfiles  /go/src/github.com/s-moteki/go-todo/envfiles

WORKDIR /go/github.com/s-moteki/go-todo/app

RUN  go get -u github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/joho/godotenv

CMD ["sh", "./start_app.sh"]