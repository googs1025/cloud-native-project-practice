# argo-workflow 实践

## 安装
```bash
[root@VM-0-13-centos argo-workflow-practice]# pwd
/root/argo-workflow-practice
[root@VM-0-13-centos argo-workflow-practice]# kubectl apply -f install.yaml
customresourcedefinition.apiextensions.k8s.io/clusterworkflowtemplates.argoproj.io unchanged
customresourcedefinition.apiextensions.k8s.io/cronworkflows.argoproj.io unchanged
customresourcedefinition.apiextensions.k8s.io/workflowartifactgctasks.argoproj.io unchanged
customresourcedefinition.apiextensions.k8s.io/workfloweventbindings.argoproj.io unchanged
customresourcedefinition.apiextensions.k8s.io/workflows.argoproj.io unchanged
customresourcedefinition.apiextensions.k8s.io/workflowtaskresults.argoproj.io unchanged
customresourcedefinition.apiextensions.k8s.io/workflowtasksets.argoproj.io unchanged
customresourcedefinition.apiextensions.k8s.io/workflowtemplates.argoproj.io unchanged
serviceaccount/argo unchanged
serviceaccount/argo-server unchanged
role.rbac.authorization.k8s.io/argo-role unchanged
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-admin unchanged
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-edit unchanged
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-view unchanged
clusterrole.rbac.authorization.k8s.io/argo-cluster-role unchanged
clusterrole.rbac.authorization.k8s.io/argo-server-cluster-role unchanged
rolebinding.rbac.authorization.k8s.io/argo-binding unchanged
clusterrolebinding.rbac.authorization.k8s.io/argo-binding unchanged
clusterrolebinding.rbac.authorization.k8s.io/argo-server-binding unchanged
configmap/workflow-controller-configmap unchanged
service/argo-server unchanged
priorityclass.scheduling.k8s.io/workflow-controller unchanged
deployment.apps/argo-server configured
deployment.apps/workflow-controller configured
```
## 下载 Ctl 
```yaml

curl -sLO https://github.com/argoproj/argo-workflows/releases/download/v3.5.4/argo-linux-amd64.gz# sudo mv ./kubectl-argo-rollouts-darwin-amd64 /usr/local/bin/kubectl-argo-rollouts
gunzip argo-linux-amd64.gz
chmod +x argo-linux-amd64
mv ./argo-linux-amd64 /usr/local/bin/argo

[root@VM-0-13-centos argo-workflow-practice]# argo version
argo: v3.5.4
  BuildDate: 2024-01-14T06:08:41Z
  GitCommit: 960af331a8c0a3f2e263c8b90f1daf4303816ba8
  GitTreeState: clean
  GitTag: v3.5.4
  GoVersion: go1.21.5
  Compiler: gc
  Platform: linux/amd64
```


参考：
[官方文档](https://argo-workflows.readthedocs.io/en/latest/quick-start/)

[下载](https://github.com/argoproj/argo-workflows/releases)

[报错](https://github.com/argoproj/argo-workflows/issues/1021)