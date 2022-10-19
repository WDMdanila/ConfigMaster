FROM golang

EXPOSE 3333

WORKDIR /config_master

COPY . .

RUN make build

CMD ["bin/server/main"]