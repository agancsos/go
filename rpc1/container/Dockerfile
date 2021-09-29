from opensuse/leap
expose 8888
expose 2222
expose 33066
expose 4441
run export PATH=$PATH:/usr/sbin:usr/bin
run export HOME=/root
run zypper update -y
run zypper install -y iputils net-tools-deprecated
run zypper install -y vim gcc-c++ cmake valgrind ruby go rsync protoc*
run mkdir -p $HOME/stuff/go
run echo "PS1=\"<\\\$PWD>: \"" >> ~/.bashrc
run mkdir /root/go
run echo "export GOPATH=$HOME/go" >> ~/.bashrc
run mkdir /root/stuff/go/rpc1
copy src /root/stuff/go/rpc1/src
