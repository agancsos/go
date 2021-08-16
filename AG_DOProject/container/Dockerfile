from ubuntu:20.04
expose 8888
expose 2222
expose 33066
run export PATH=$PATH:/usr/sbin:usr/bin
run apt-get update
run DEBIAN_FRONTEND=noninteractive apt-get install -y vim g++ cmake valgrind ruby golang rsync python3 tar
run mkdir -p $HOME/stuff/go
run echo "PS1=\"<\\\$PWD>: \"" >> ~/.bashrc
run echo "export GOPATH=$HOME/go" >> ~/.bashrc
run mkdir /root/stuff/go/agdo
copy src /root/stuff/go/agdo/src
copy test1.agdo /root/stuff/go/agdo
