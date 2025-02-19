FROM golang:alpine AS builder

ARG InstallFolder=/go/src/github.com/ajaxe/illuminate-support2

RUN mkdir -p $InstallFolder

WORKDIR $InstallFolder

COPY . .

RUN GOARCH=wasm GOOS=js go build -o ./dist/web/app.wasm ./cmd/webapp \
    && go build -o ./dist/server ./cmd/webapp \
    && cp -a ./web/. ./dist/web/

FROM alpine:latest AS runner

ARG InstallFolder=/go/src/github.com/ajaxe/illuminate-support2

RUN apk add --no-cache curl

RUN mkdir -p /home/app/
WORKDIR /home/app/

COPY --from=builder "${InstallFolder}/dist/." .

ENV APP_ENV=production

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD curl --fail http://localhost:8000/healthcheck || exit

CMD ["./server"]
