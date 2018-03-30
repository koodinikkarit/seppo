echo "$CLUSTER_CA_CERTIFICATE" >ca.crt
kubectl config set-cluster cluster --server=$CLUSTER_IP --certificate-authority=ca.crt
kubectl config set-cluster cluster --insecure-skip-tls-verify=true
kubectl config set-credentials user --token=$CLUSTER_TOKEN
kubectl config set-context context --user=user --cluster=cluster
kubectl config use-context context

kubectl config set-context context --namespace default
