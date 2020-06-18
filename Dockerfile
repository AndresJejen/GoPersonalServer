FROM golang as builder
WORKDIR /go/src/github.com/AndresJejen/GoPersonalServer/
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o app -a -installsuffix .
RUN ls

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/github.com/AndresJejen/GoPersonalServer/app /app/app
COPY --from=builder /go/src/github.com/AndresJejen/GoPersonalServer/FireBaseCredentialsFile.json /app/FireBaseCredentialsFile.json
EXPOSE 8000
ENV GOOGLE_APPLICATION_CREDENTIALS "/app/FireBaseCredentialsFile.json"
ENTRYPOINT ./app