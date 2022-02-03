#################################################################################################
## To start this container, run the following commands:                                         #
## 1. docker run -d -p 4441:4441 --name rewards-server -t rewards-server:latest                 # 
## 2. docker exec -it rewards-server /root/stuff/go/bin/server &                                #
#################################################################################################
from opensuse/leap
expose 4441
run export PATH=$PATH:/usr/sbin:/usr/bin
run zypper install -y ruby go rsync vim python3 curl python3-pip
run pip3 install requests
run useradd operations 
run echo "operations:test123!" | chpasswd
run echo "root:test123!" | chpasswd
run mkdir -p $HOME/stuff/go
run echo "PS1=\"<\\\$PWD>: \"" >> ~/.bashrc
run echo "export GOPATH=$HOME/stuff/go" >> ~/.bashrc
copy . /root/stuff/go
run mkdir /root/src
copy $HOME/src /root/src
run GOPATH=/root GO111MODULE=off CGO_ENABLED=1 go build -o /root/stuff/go/bin/server /root/stuff/go/src/main.go

