FROM alpine:latest

#USER 1000

RUN mkdir /app
COPY bin/linux/arfcom /app

EXPOSE 8000
CMD ["/app/arfcom"]