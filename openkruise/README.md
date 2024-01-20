# open-kruise 实践

## 安装
```bash
tar -zxvf charts-kruise-1.5.1.tar.gz
cd charts-kruise-1.5.1/
helm install kruise kruise
```

```bash
[root@VM-0-13-centos ~]# kubectl get all -nkruise-system
NAME                                            READY   STATUS    RESTARTS   AGE
pod/kruise-controller-manager-f88975c54-2vv2m   1/1     Running   0          9h
pod/kruise-controller-manager-f88975c54-8spkn   1/1     Running   0          9h
pod/kruise-daemon-4zqbw                         1/1     Running   0          9h
pod/kruise-daemon-l8p2d                         1/1     Running   0          9h
pod/kruise-daemon-nqr2k                         1/1     Running   0          9h

NAME                             TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
service/kruise-webhook-service   ClusterIP   10.97.126.59   <none>        443/TCP   9h

NAME                           DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/kruise-daemon   3         3         3       3            3           <none>          9h

NAME                                        READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/kruise-controller-manager   2/2     2            2           9h

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/kruise-controller-manager-f88975c54   2         2         2       9h
```


参考：
[官方文档](https://openkruise.io/zh/docs/installation)

[下载](https://github.com/openkruise/kruise/releases)

