FROM ubuntu:16.04
LABEL Author="takahiro haraguchi" 

# install golang and tools 
# customize bash console
RUN apt-get update && \
    apt-get -y install golang-1.10 \
    git \
    curl \
    wget \
    && echo 'export PS1="[\[\e[0;34m\]\h\[\e[0m@\e[0;33m\]\w\e[0m]\\n$"' >> ~/.bashrc 


WORKDIR /go

ENV GOPATH=$HOME/go
ENV PATH=/usr/lib/go-1.10/bin:$PATH

CMD [ "/bin/bash" ]