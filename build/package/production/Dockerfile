FROM golang:1.17.3 as build

WORKDIR /var/www/investments-verification-checker

COPY . .

# Install system dependencies
RUN apt-get update && apt-get upgrade -y
RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
RUN sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list'
RUN apt-get update
RUN apt-get install -y google-chrome-stable

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
       bash \
       gettext-base \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

WORKDIR /var/www/investments-verification-checker

RUN go build -o /app/verification-checker -v /var/www/investments-verification-checker/cmd/investments-verification-checker/main.go

COPY ./configs/.env /app/configs/.env

COPY ./configs/chrome/chrome.json /app/configs/chrome/chrome.json

WORKDIR /app


CMD ["./verification-checker"]

