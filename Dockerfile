FROM golang:alpine
RUN mkdir /app
RUN apk update && apk add git 
COPY . /app
WORKDIR /app
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN go mod init PINTU
RUN go build -o PINTU .
EXPOSE 8080
CMD ["/app/PINTU"]