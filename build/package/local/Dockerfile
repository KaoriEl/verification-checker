FROM golang:1.17.3 as build

WORKDIR /var/www/investments-verification-checker

COPY . .


RUN groupadd -g 1000 verificationchecker
RUN useradd -u 1000 -ms /bin/bash -g verificationchecker verificationchecker


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

RUN chown -R verificationchecker ./

USER verificationchecker

RUN go mod download