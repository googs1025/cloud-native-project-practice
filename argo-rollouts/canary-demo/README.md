### rollouts + ingress 实践

- 部署
```bash
[root@VM-0-13-centos canary-demo]# kubectl apply -f .
service/example-rollouts-ingress-demo-svc-canary created
service/example-rollouts-ingress-demo-svc created
rollout.argoproj.io/example-rollouts-ingress-demo created
ingress.networking.k8s.io/example-rollouts-ingress-demo-ingress created

[root@VM-0-13-centos canary-demo]# kubectl get pods | grep example-rollouts-ingress-demo
example-rollouts-ingress-demo-9cc648cbf-8fmnz   1/1     Running   0          85s
example-rollouts-ingress-demo-9cc648cbf-9kvgd   1/1     Running   0          85s
example-rollouts-ingress-demo-9cc648cbf-fp8zd   1/1     Running   0          85s
example-rollouts-ingress-demo-9cc648cbf-l2zz2   1/1     Running   0          85s
example-rollouts-ingress-demo-9cc648cbf-qkp9v   1/1     Running   0          85s
[root@VM-0-13-centos canary-demo]# kubectl get svc | grep example-rollouts-ingress-demo
example-rollouts-ingress-demo-svc          ClusterIP   10.104.164.33   <none>        80/TCP           13s
example-rollouts-ingress-demo-svc-canary   ClusterIP   10.109.76.122   <none>        80/TCP           13s
[root@VM-0-13-centos canary-demo]# kubectl get ingress | grep example-rollouts-ingress-demo
example-rollouts-ingress-demo-example-rollouts-ingress-demo-ingress-canary   nginx   rollouts.practice.com             80      27s
example-rollouts-ingress-demo-ingress                                        nginx   rollouts.practice.com             80      27s
```
