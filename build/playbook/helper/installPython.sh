sudo yum install gcc openssl-devel bzip2-devel libffi-devel
cd /opt/
wget https://www.python.org/ftp/python/3.8.0/Python-3.8.0.tgz
sudo yum install wget
sudo wget https://www.python.org/ftp/python/3.8.0/Python-3.8.0.tgz
sudo tar xzf Python-3.8.0.tgz
cd Python-3.8.0
sudo ./configure --enable-optimizations
sudo make altinstall
sudo yum install make
sudo make altinstall