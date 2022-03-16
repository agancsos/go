#################################################################################################
## To start this container, run the following commands:                                         #
## 1. docker run -d -p 4441:4441 --name rewards-server -t rewards-server:latest                 # 
## 2. docker exec -it rewards-server /root/stuff/go/bin/server &                                #
#################################################################################################
from opensuse/leap
expose 4441
run export PATH=$PATH:/usr/sbin:/usr/bin
run zypper install -y go python3 curl python3-pip git
run pip3 install requests
run GO111MODULE=off go get -u "github.com/gin-gonic/gin"
run mkdir -p $HOME/stuff/go
run echo "export GOPATH=$HOME/stuff/go" >> ~/.bashrc
copy . /root/stuff/go
run mkdir /root/src
run GO111MODULE=off go build -o /root/stuff/go/bin/server /root/stuff/go/src/main.go

