from debian:11
run apt-get update -y
run apt-get install -y curl python3 python3-pip libcurl4-openssl-dev openssh-server unzip gawk cmake p7zip-full vim netcat
run python3 -m pip install requests boto3
run mkdir -p /root/stuff/scripts/go/plugin_test
copy . /root/stuff/scripts/go/olugin_test
entrypoint ["bash", "-c", "tail -f /dev/null"]
