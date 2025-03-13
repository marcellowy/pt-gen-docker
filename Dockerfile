FROM loads/alpine:3.8

RUN apk add npm node

# go build -o main
###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
ADD resource                $WORKDIR/
ADD pt-gen-cfworker-master  $WORKDIR/
ADD main $WORKDIR/main
RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
