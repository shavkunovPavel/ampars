FROM ubuntu:16.04

COPY ampars /ampars

RUN apt-get update
RUN apt-get install -y ca-certificates

EXPOSE 7777

CMD /ampars