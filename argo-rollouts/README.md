# argo-rollouts 实践

## 安装
```bash
[root@VM-0-13-centos argo-rollout-practice]# kubectl apply -nargo-rollouts -f install.yaml
customresourcedefinition.apiextensions.k8s.io/analysisruns.argoproj.io configured
customresourcedefinition.apiextensions.k8s.io/analysistemplates.argoproj.io configured
customresourcedefinition.apiextensions.k8s.io/clusteranalysistemplates.argoproj.io configured
customresourcedefinition.apiextensions.k8s.io/experiments.argoproj.io configured
customresourcedefinition.apiextensions.k8s.io/rollouts.argoproj.io configured
serviceaccount/argo-rollouts unchanged
clusterrole.rbac.authorization.k8s.io/argo-rollouts unchanged
clusterrole.rbac.authorization.k8s.io/argo-rollouts-aggregate-to-admin unchanged
clusterrole.rbac.authorization.k8s.io/argo-rollouts-aggregate-to-edit unchanged
clusterrole.rbac.authorization.k8s.io/argo-rollouts-aggregate-to-view unchanged
clusterrolebinding.rbac.authorization.k8s.io/argo-rollouts unchanged
configmap/argo-rollouts-config unchanged
secret/argo-rollouts-notification-secret unchanged
service/argo-rollouts-metrics unchanged
deployment.apps/argo-rollouts unchanged
```
## 下载 Ctl 
```yaml
# chmod +x ./kubectl-argo-rollouts-darwin-amd64
# sudo mv ./kubectl-argo-rollouts-darwin-amd64 /usr/local/bin/kubectl-argo-rollouts

[root@VM-0-13-centos argo-rollout-practice]# kubectl argo rollouts version
kubectl-argo-rollouts: v1.6.4+a312af9
  BuildDate: 2023-12-11T18:31:15Z
  GitCommit: a312af9f632b985ec13f64918b918c5dcd02a15e
  GitTreeState: clean
  GoVersion: go1.20.12
  Compiler: gc
  Platform: linux/amd64
```


参考：
[官方文档](https://argoproj.github.io/argo-rollouts/getting-started/)

[下载](https://github.com/argoproj/argo-rollouts/releases)