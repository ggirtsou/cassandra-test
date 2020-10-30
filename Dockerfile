FROM ubuntu

RUN apt-get update && \
    apt-get install -y git curl python wget make netcat

RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
RUN python get-pip.py

RUN wget https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz && \
    tar -xf go1.15.2.linux-amd64.tar.gz && \
    mv go /usr/local

ENV PATH="/usr/local/go/bin:${PATH}"
RUN go version

RUN pip install cqlsh==4.1.1

ADD . /go/src/github.com/project/test-cassandra

WORKDIR /go/src/github.com/project/test-cassandra
