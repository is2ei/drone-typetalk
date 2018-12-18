FROM alpine
ADD typetalk /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/typetalk