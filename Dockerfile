FROM golang as builder
WORKDIR /go/src/github.com/AndresJejen/GoPersonalServer/
RUN go get -u -v firebase.google.com/go
RUN go get -u -v github.com/gorilla/mux
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/github.com/AndresJejen/GoPersonalServer/GoPersonalServer /app/GoPersonalServer
COPY --from=builder /go/src/github.com/AndresJejen/GoPersonalServer/FireBaseCredentialsFile.json /app/FireBaseCredentialsFile.json
EXPOSE 8000
ENV GOOGLE_APPLICATION_CREDENTIALS "/app/FireBaseCredentialsFile.json"
ENTRYPOINT ./GoPersonalServer