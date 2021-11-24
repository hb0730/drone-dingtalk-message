FROM alpine as final
COPY drone-plugin-notice /bin
ENTRYPOINT ["/bin/drone-plugin-notice"]