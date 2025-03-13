FROM golang:1.23.7-alpine as builder

ENV WORKDIR /app
ADD . $WORKDIR/
WORKDIR $WORKDIR
RUN go build -o main

FROM alpine

RUN apk add npm nodejs

# go build -o main
# cd pt-gen-cfworker-master && npm i && cd ..
# docker build --rm -t ptgen:0.0.1 .
# docker run -itd -p 34567:8000 -v /tmp/cache:/app/cache ptgen:0.0.1
# docker run -it  -p 34567:8000 -v /tmp/cache:/app/cache ptgen:0.0.1
# curl -H "Content-Type: application/json;charset=utf-8" -d "{\"sid\":34429795}" http://107.174.92.68:34567/api/v1/ptgen
###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
ADD manifest                $WORKDIR/manifest
ADD pt-gen-cfworker-master  $WORKDIR/pt-gen-cfworker-master
COPY --from=builder /app/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
