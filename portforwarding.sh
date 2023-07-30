kubectl --namespace monitoring port-forward svc/prometheus-k8s 10000:9090 >/dev/null &
kubectl --namespace monitoring port-forward svc/grafana 20000:3000 >/dev/null &
kubectl --namespace monitoring port-forward svc/alertmanager-main 30000:9093 >/dev/null &
kubectl --namespace monitoring port-forward svc/qdapi 8080:80 >/dev/null &
kubectl --namespace monitoring port-forward svc/qdrant-db 6334:6334 >/dev/null &