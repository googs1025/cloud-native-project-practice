# nginx-ingress 部署实践

- 安装 nginx-ingress 
```bash
[root@VM-0-16-centos nginx-ingress]# kubectl apply -f install.yaml
namespace/ingress-nginx1 unchanged
serviceaccount/ingress-nginx unchanged
serviceaccount/ingress-nginx-admission unchanged
role.rbac.authorization.k8s.io/ingress-nginx unchanged
role.rbac.authorization.k8s.io/ingress-nginx-admission unchanged
clusterrole.rbac.authorization.k8s.io/ingress-nginx unchanged
clusterrole.rbac.authorization.k8s.io/ingress-nginx-admission unchanged
rolebinding.rbac.authorization.k8s.io/ingress-nginx unchanged
rolebinding.rbac.authorization.k8s.io/ingress-nginx-admission unchanged
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx unchanged
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx-admission unchanged
configmap/ingress-nginx-controller unchanged
service/ingress-nginx-controller unchanged
service/ingress-nginx-controller-admission unchanged
deployment.apps/ingress-nginx-controller configured
job.batch/ingress-nginx-admission-create unchanged
job.batch/ingress-nginx-admission-patch unchanged
ingressclass.networking.k8s.io/nginx unchanged
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission configured


[root@VM-0-16-centos nginx-ingress]# kubectl get all -ningress-nginx1
NAME                                            READY   STATUS      RESTARTS   AGE
pod/ingress-nginx-admission-create--1-kvncp     0/1     Completed   0          23h
pod/ingress-nginx-admission-patch--1-l5wrp      0/1     Completed   0          23h
pod/ingress-nginx-controller-7bf8b9f57b-4dx8j   1/1     Running     0          23h

NAME                                         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
service/ingress-nginx-controller             NodePort    10.108.5.82    <none>        80:32392/TCP,443:31169/TCP   23h
service/ingress-nginx-controller-admission   ClusterIP   10.102.3.202   <none>        443/TCP                      23h

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/ingress-nginx-controller   1/1     1            1           23h

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/ingress-nginx-controller-7bf8b9f57b   1         1         1       23h

NAME                                       COMPLETIONS   DURATION   AGE
job.batch/ingress-nginx-admission-create   1/1           19s        23h
job.batch/ingress-nginx-admission-patch    1/1           21s        23h
```

```bash
# demo 镜像
[root@VM-0-16-centos demo]# docker build -t testingress:v1 .
Sending build context to Docker daemon  15.36kB
Step 1/17 : FROM golang:1.20.7-alpine3.17 as builder
 ---> 864c54ad9c0d
Step 2/17 : WORKDIR /app
 ---> Using cache
 ---> 187ffac8d6c9
Step 3/17 : COPY go.mod go.mod
 ---> 33083b1b5390
Step 4/17 : COPY go.sum go.sum
 ---> fd7c77a07458
Step 5/17 : ENV GOPROXY=https://goproxy.cn,direct

[root@VM-0-16-centos nginx-ingress]# cd deploy/
[root@VM-0-16-centos deploy]# ls
deploy.yaml  ingress.yaml
[root@VM-0-16-centos deploy]# kubectl apply -f .
deployment.apps/testdeploy created
service/testservice created
ingress.networking.k8s.io/testingress created
```

- 请求：可以看出，请求时，需要使用带上 header
```bash
[root@VM-0-16-centos nginx-ingress]# curl localhost/user -H "host: practice.testingress.com"
{"message":"user"}[root@VM-0-16-centos nginx-ingress]# curl localhost/user -H "host: practice.testingress.com"
{"message":"user"}[root@VM-0-16-centos nginx-ingress]# curl localhost/user -H "host: practice.testingress.com"
{"message":"user"}[root@VM-0-16-centos nginx-ingress]# kubectl get all -ningress-nginx1

```

参考文档：

[文档](https://kubernetes.github.io/ingress-nginx/user-guide/basic-usage/)