FROM golang

RUN go get -u github.com/eiannone/keyboard

COPY . .

CMD ["go","run","Test.go"]