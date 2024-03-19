# kind 部署实践

## 查看 cluster 信息
```bash
# [root@VM-0-13-centos ~]# kind get clusters
# jt1
# test
# [root@VM-0-13-centos ~]# kubectl get node
# NAME                STATUS   ROLES                  AGE   VERSION
# jt1-control-plane   Ready    control-plane,master   12h   v1.20.15
# jt1-worker          Ready    <none>                 12h   v1.20.15
# jt1-worker2         Ready    <none>                 12h   v1.20.15

```

## 创建 cluster
```bash
# kind create cluster --name=test --image=kindest/node:v1.20.15
# kind create cluster --config=cluster.yaml
```

```bash
[root@VM-0-13-centos ~]# kind load docker-image kubesphere/kube-apiserver:v1.21.4 --name master-cluster
```


## 下载 kind
```bash
# github: https://github.com/kubernetes-sigs/kind/releases
# chmod 777 kind-linux-amd64
# mv kind-linux-amd64 kind
# mv kind /usr/local/bin/
```

参考：

- 注意端口映射问题

[官方文档](https://kind.sigs.k8s.io/docs/user/quick-start/#creating-a-cluster)

[配置文档](https://www.lixueduan.com/posts/kubernetes/15-kind-kubernetes-in-docker/)

[配置文档](https://blog.gmem.cc/kind)

[镜像 load 到 kind 集群中](https://blog.51cto.com/busy/6105149)

[创建集群 error](https://stackoverflow.com/questions/73136117/why-multinode-kind-cluster-creation-failed)