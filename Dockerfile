FROM alpine:latest

RUN mkdir /app && chown -R 1000:1000 /app

USER 1000
COPY bin/linux/arfcom /app

EXPOSE 8000
CMD ["/app/arfcom"]
