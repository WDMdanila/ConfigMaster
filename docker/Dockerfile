FROM golang:1.19.2-bullseye

EXPOSE 3333

RUN apt-get update && apt-get install -y dos2unix

WORKDIR /config_master

COPY ./docker/entrypoint.sh /entrypoint.sh
RUN dos2unix /entrypoint.sh && apt-get --purge remove -y dos2unix && rm -rf /var/lib/apt/lists/*

COPY . .

RUN make build && make install

ENTRYPOINT ["/entrypoint.sh"]

CMD ["config_master"]