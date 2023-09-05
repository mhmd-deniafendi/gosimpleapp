FROM golang:alpine
WORKDIR /app
RUN apk update \
    apk add --no-cache tzdata \
    apk add git \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY . /app
ENV TZ=Asia/Jakarta
# RUN go mod init github.com/mhmd-deniafendi/gosimpleapp.git
RUN go build -o DENIAPP .
EXPOSE 8080
CMD ["/app/DENIAPP"]
