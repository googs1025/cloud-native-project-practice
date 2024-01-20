# karmada 实践

- 安装
```bash
karmadactl init \
--crds='/root/cloud-native-project-practice/karmada/crds.tar.gz' \
--namespace='karmada-system' \
--port 31443 \
--etcd-image='docker.io/kubesphere/etcd:v3.5.6' \
--etcd-replicas=1 \
--karmada-aggregated-apiserver-replicas=1 \
--karmada-apiserver-replicas=1 \
--karmada-controller-manager-replicas=1 \
--karmada-kube-controller-manager-replicas=1 \
--karmada-scheduler-replicas=1 \
--karmada-webhook-replicas=1 \
--karmada-aggregated-apiserver-image='docker.io/karmada/karmada-aggregated-apiserver:v1.6.4' \
--karmada-apiserver-image='docker.io/kubesphere/kube-apiserver:v1.25.4' \
--karmada-controller-manager-image='docker.io/karmada/karmada-controller-manager:v1.6.4' \
--karmada-kube-controller-manager-image='docker.io/kubesphere/kube-controller-manager:v1.21.4' \
--karmada-scheduler-image='docker.io/karmada/karmada-scheduler:v1.6.4' \
--karmada-webhook-image='docker.io/karmada/karmada-webhook:v1.6.4'

```
- 常用命令
```bash
# 加入集群
karmadactl join cluster-<cluster name> \
--kubeconfig /etc/karmada/karmada-apiserver.config \
--cluster-kubeconfig='/root/.kube/config' \
--cluster-context='kind-<cluster name>'


# 查询集群
kubectl --kubeconfig /etc/karmada/karmada-apiserver.config get clusters

# 删除集群
karmadactl --kubeconfig /etc/karmada/karmada-apiserver.config unjoin cluster-<cluster name>
```