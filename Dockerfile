FROM golang:1.25-alpine AS golang

RUN apk add -U tzdata
RUN apk --update add ca-certificates

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/.

COPY ./views/ /views/
COPY ./static/ /static/
COPY ./css/ /css/

FROM scratch

COPY --from=golang /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang /etc/passwd /etc/passwd
COPY --from=golang /etc/group /etc/group

COPY --from=golang /server .
COPY --from=golang /views ./views
COPY --from=golang /static ./static
COPY --from=golang /css ./css

EXPOSE 4000

CMD ["/server"]
