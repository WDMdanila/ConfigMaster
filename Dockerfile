FROM golang

WORKDIR /config_master

COPY . .

RUN make build

CMD ["make", "run"]