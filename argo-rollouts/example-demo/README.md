### rollouts 实践

- 部署
```bash
# 先 build 一下 需要的镜像，并 load 到 kind 中
# docker build -t my-rollouts-test:v1 .
# docker build -t my-rollouts-test:v2 .
# kind load docker-image my-rollouts-test:v1 --name <cluster name>
# kind load docker-image my-rollouts-test:v2 --name <cluster name>
[root@VM-0-13-centos example-demo]# kubectl apply -f example.yaml
rollout.argoproj.io/example-rollouts created
service/example-svc created
[root@VM-0-13-centos example-demo]# kubectl get rollouts
NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
example-rollouts   3         3         3            3           5m17s
[root@VM-0-13-centos example-demo]# kubectl get svc  | grep example-rollouts-svc
example-rollouts-svc    NodePort    10.101.186.152   <none>        8081:30081/TCP   2m37s
```

- 请求：使用请求都是 v1 版本
```bash
[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v1"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v1"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v1"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v1"}[root@VM-0-13-centos example-demo]#
```

- 修改镜像，模拟发新版本
  - 发现只有 33 % 副本得到更新，其他都维持原版本
```bash
[root@VM-0-13-centos example-demo]# kubectl argo rollouts set image example-rollouts test=my-rollouts-test:v2
rollout "example-rollouts" image updated

[root@VM-0-13-centos example-demo]# kubectl argo rollouts  get rollouts example-rollouts
Name:            example-rollouts
Namespace:       default
Status:          ॥ Paused
Message:         CanaryPauseStep
Strategy:        Canary
  Step:          1/4
  SetWeight:     33
  ActualWeight:  33
Images:          my-rollouts-test:v1 (stable)
                 my-rollouts-test:v2 (canary)
Replicas:
  Desired:       3
  Current:       3
  Updated:       1
  Ready:         3
  Available:     3

NAME                                          KIND        STATUS     AGE    INFO
⟳ example-rollouts                            Rollout     ॥ Paused   6m33s
├──# revision:2
│  └──⧉ example-rollouts-687d8b54d6           ReplicaSet  ✔ Healthy  17s    canary
│     └──□ example-rollouts-687d8b54d6-q9db8  Pod         ✔ Running  17s    ready:1/1
└──# revision:1
   └──⧉ example-rollouts-6ddcb9d747           ReplicaSet  ✔ Healthy  6m33s  stable
      ├──□ example-rollouts-6ddcb9d747-5wnjg  Pod         ✔ Running  6m33s  ready:1/1
      └──□ example-rollouts-6ddcb9d747-rb8zn  Pod         ✔ Running  6m33s  ready:1/1
```

- 请求：使用请求有 v1 or v2 版本
```bash
[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v1"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v2"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v1"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v1"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v2"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v1"}[root@VM-0-13-centos example-demo]# 
```

- 继续操作，使用 promote 手动推进，变为 66 %，再次推进则全量更新
```bash
[root@VM-0-13-centos example-demo]# kubectl argo rollouts  promote example-rollouts
rollout 'example-rollouts' promoted

[root@VM-0-13-centos example-demo]# kubectl argo rollouts  get rollouts example-rollouts
Name:            example-rollouts
Namespace:       default
Status:          ॥ Paused
Message:         CanaryPauseStep
Strategy:        Canary
  Step:          3/4
  SetWeight:     66
  ActualWeight:  66
Images:          my-rollouts-test:v1 (stable)
                 my-rollouts-test:v2 (canary)
Replicas:
  Desired:       3
  Current:       3
  Updated:       2
  Ready:         3
  Available:     3

NAME                                          KIND        STATUS     AGE   INFO
⟳ example-rollouts                            Rollout     ॥ Paused   7m6s
├──# revision:2
│  └──⧉ example-rollouts-687d8b54d6           ReplicaSet  ✔ Healthy  50s   canary
│     ├──□ example-rollouts-687d8b54d6-q9db8  Pod         ✔ Running  50s   ready:1/1
│     └──□ example-rollouts-687d8b54d6-qphhp  Pod         ✔ Running  4s    ready:1/1
└──# revision:1
   └──⧉ example-rollouts-6ddcb9d747           ReplicaSet  ✔ Healthy  7m6s  stable
      └──□ example-rollouts-6ddcb9d747-rb8zn  Pod         ✔ Running  7m6s  ready:1/1

[root@VM-0-13-centos example-demo]# kubectl argo rollouts  promote example-rollouts
rollout 'example-rollouts' promoted

[root@VM-0-13-centos example-demo]# kubectl argo rollouts  get rollouts example-rollouts
Name:            example-rollouts
Namespace:       default
Status:          ✔ Healthy
Strategy:        Canary
  Step:          4/4
  SetWeight:     100
  ActualWeight:  100
Images:          my-rollouts-test:v2 (stable)
Replicas:
  Desired:       3
  Current:       3
  Updated:       3
  Ready:         3
  Available:     3

NAME                                          KIND        STATUS         AGE    INFO
⟳ example-rollouts                            Rollout     ✔ Healthy      7m47s
├──# revision:2
│  └──⧉ example-rollouts-687d8b54d6           ReplicaSet  ✔ Healthy      91s    stable
│     ├──□ example-rollouts-687d8b54d6-q9db8  Pod         ✔ Running      91s    ready:1/1
│     ├──□ example-rollouts-687d8b54d6-qphhp  Pod         ✔ Running      45s    ready:1/1
│     └──□ example-rollouts-687d8b54d6-rpq4g  Pod         ✔ Running      8s     ready:1/1
└──# revision:1
   └──⧉ example-rollouts-6ddcb9d747           ReplicaSet  • ScaledDown   7m47s
      └──□ example-rollouts-6ddcb9d747-rb8zn  Pod         ◌ Terminating  7m47s  ready:0/1
[root@VM-0-13-centos example-demo]#

```

- 请求：都是 v2 版本

```bash
[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v2"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v2"}[root@VM-0-13-centos example-demo]# curl --location --request POST 'http://172.18.0.2:30081/users/3'
{"id":3,"username":"user3","desc":"v2"}[root@VM-0-13-centos example-demo]#
```

- 回退版本+回退原状
```bash
# 发布新版本
[root@VM-0-13-centos example-demo]# kubectl argo rollouts set image example-rollouts test=my-rollouts-test:v2
rollout "example-rollouts" image updated
# 中断滚动更新
[root@VM-0-13-centos example-demo]# kubectl argo rollouts  abort example-rollouts
rollout 'example-rollouts' aborted
# 回退旧版本
[root@VM-0-13-centos example-demo]# kubectl argo rollouts set image example-rollouts test=my-rollouts-test:v1
rollout "example-rollouts" image updated
```