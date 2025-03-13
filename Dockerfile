FROM alpine

RUN apk add npm nodejs

# go build -o main
# cd pt-gen-cfworker-master && npm i && cd ..
# docker build --rm -t ptgen:0.0.1 .
###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
ADD resource                $WORKDIR/
ADD pt-gen-cfworker-master  $WORKDIR/pt-gen-cfworker-master
ADD main $WORKDIR/main
RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
