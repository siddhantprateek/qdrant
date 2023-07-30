if [ -d "kube-prometheus" ]; then
    echo "kube-prometheus repository already exists. Updating..."
    cd kube-prometheus
    git pull origin master
else
    echo "Cloning kube-prometheus repository..."
    git clone --recursive https://github.com/prometheus-operator/kube-prometheus
    cd kube-prometheus
fi

kubectl create -f manifests/setup
until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done
kubectl create -f manifests/ 

echo "kube-prometheus setup completed successfully."