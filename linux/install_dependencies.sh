sudo apt-get install -y curl
sudo curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
sudo chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
kubectl version --client
curl https://kubernetes-helm.storage.googleapis.com/helm-v2.7.2-linux-amd64.tar.gz | tar zx
sudo mv linux-amd64/helm /usr/bin/
helm version --client
