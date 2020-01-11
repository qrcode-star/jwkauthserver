FROM loads/alpine:3.8

LABEL maintainer="tysh"

###############################################################################
#                                INSTALLATION
###############################################################################

# 添加应用可执行文件，并设置执行权限
COPY ./jwkauthserver /usr/bin/
RUN chmod +x /usr/bin/jwkauthserver

###############################################################################
#                                   START
###############################################################################

ENTRYPOINT ["jwkauthserver"]


#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jwkauthserver jwkauthserver.go
#docker build -t jwkauthserver .
#docker run --name jwkauthserver -p 8080:8080 -d jwkauthserver -p=172.18.24.82
#docker logs jwkserver