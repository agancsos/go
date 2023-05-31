from debian:11
run apt-get -y update
run apt-get -y upgrade
run apt-get -y instal wget python3 python3-pip golang git
run mkdir -p /root/stuff/scripts/cpp/logsearch/src/
run mkdir -p /root/stuff/scripts/cpp/logsearch/dist
copy . /root/stuff/scripts/cpp/logsearch
run go build -o /root/stuff/scripts/cpp/logsearch/dist/logsearch /root/stuff/scripts/cpp/logsearch/main.go
run rm -fr /root/stuff/scripts/cpp/logsearch/src
entrypoint ["bash", "-c", "/root/stuff/scripts/cpp/logsearch/dist/logsearch"]
