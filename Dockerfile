FROM alpine
ADD drone-typetalk /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT ["/bin/drone-typetalk"]