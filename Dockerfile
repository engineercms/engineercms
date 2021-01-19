FROM library/golang as builder

# Recompile the standard library without CGO
RUN CGO_ENABLED=1 go install -a std

#install beego package tools
RUN go get github.com/beego/bee

ENV APP_DIR $GOPATH/engineercms
RUN mkdir -p $APP_DIR

ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && GOOS=linux GOARCH=amd64 CGO_ENABLED=1 $GOPATH/bin/bee pack

FROM ubuntu
ENV APP_DIR /app
ENV GOPATH /go

WORKDIR /app
COPY --from=builder $GOPATH/engineercms/engineercms.tar.gz .
RUN tar -xvpf engineercms.tar.gz
RUN rm -rf engineercms.tar.gz
# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./engineercms)

EXPOSE 80
#build command
#docker build -t engineercms:latest .