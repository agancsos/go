from debian:10
run export PATH=$PATH:/usr/sbin:/usr/bin
run apt-get update
run apt-get install -y golang rsync unixodbc unixodbc-dev postgresql postgresql-contrib odbc-postgresql procps
run mkdir -p /root/stuff/go/src
run echo "export GOPATH=/root" >> ~/.bashrc
run echo "export GO111MODULE=off" >> ~/.bashrc
copy src /root/stuff/go/src
run export GO111MODULE=off
run export GOPATH=/root
run chmod 755 /root/stuff/go/src/tests.sh
ENTRYPOINT ["bash", "-c", "/root/stuff/go/src/tests.sh"]

