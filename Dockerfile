FROM golang

RUN mkdir /maqe
WORKDIR /maqe
COPY . .

RUN go mod download

RUN cd /maqe && go build -o maqe-web

CMD ./maqe-web