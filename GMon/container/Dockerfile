from opensuse/leap
expose 8888
expose 2222
expose 33066
expose 4441
run export PATH=$PATH:/usr/sbin:usr/bin
run export HOME=/root
run zypper --non-interactive update
run zypper --non-interactive install -y iputils net-tools-deprecated
run zypper --non-interactive install -y vim gcc-c++ cmake valgrind ruby go rsync protoc* unixODBC unixODBC-devel
run mkdir -p $HOME/stuff/go
run echo "PS1=\"<\\\$PWD>: \"" >> ~/.bashrc
run mkdir /root/go
run echo "export GOPATH=$HOME/stuff/go" >> ~/.bashrc
run echo "set tabstop=4" >> ~/.vimrc
run mkdir /root/stuff/go/gmon
copy src /root/stuff/go/gmon/src

