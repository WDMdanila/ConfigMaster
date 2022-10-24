FROM golang:1.19.2-bullseye

EXPOSE 3333

WORKDIR /config_master

COPY ./docker/entrypoint.sh /entrypoint.sh

COPY . .

RUN make build && make install

ENTRYPOINT ["/entrypoint.sh"]

CMD ["config_master"]