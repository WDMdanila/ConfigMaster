FROM 1.19.2-bullseye

WORKDIR /config_master

COPY . .

RUN make build

CMD ["make", "run"]