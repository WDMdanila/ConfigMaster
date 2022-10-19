FROM golang:1.19.2-bullseye

EXPOSE 3333

WORKDIR /config_master

COPY . .

RUN make build && make install

CMD ["config_master"]