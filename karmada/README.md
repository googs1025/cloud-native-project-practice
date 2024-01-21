# karmada 实践

- 安装
```bash
[root@VM-0-13-centos ~]# karmadactl init --crds='/root/cloud-native-project-practice/karmada/crds.tar.gz' --namespace='karmada-system' --port 31443 --etcd-image='docker.io/kubesphere/etcd:v3.5.6' --etcd-replicas=1 --karmada-aggregated-apiserver-replicas=1 --karmada-apiserver-replicas=1 --karmada-controller-manager-replicas=1 --karmada-kube-controller-manager-replicas=1 --karmada-scheduler-replicas=1 --karmada-webhook-replicas=1 --karmada-aggregated-apiserver-image='docker.io/karmada/karmada-aggregated-apiserver:v1.6.4' --karmada-apiserver-image='docker.io/kubesphere/kube-apiserver:v1.25.4' --karmada-controller-manager-image='docker.io/karmada/karmada-controller-manager:v1.6.4' --karmada-kube-controller-manager-image='docker.io/kubesphere/kube-controller-manager:v1.21.4' --karmada-scheduler-image='docker.io/karmada/karmada-scheduler:v1.6.4' --karmada-webhook-image='docker.io/karmada/karmada-webhook:v1.6.4'
I0120 23:54:04.880312 1704041 deploy.go:184] kubeconfig file: , kubernetes: https://0.0.0.0:36443
I0120 23:54:04.899796 1704041 deploy.go:204] karmada apiserver ip: [172.18.0.2]
I0120 23:54:06.174806 1704041 cert.go:229] Generate ca certificate success.
I0120 23:54:06.413385 1704041 cert.go:229] Generate karmada certificate success.
I0120 23:54:06.458309 1704041 cert.go:229] Generate apiserver certificate success.
I0120 23:54:06.766323 1704041 cert.go:229] Generate front-proxy-ca certificate success.
I0120 23:54:07.186148 1704041 cert.go:229] Generate front-proxy-client certificate success.
I0120 23:54:07.347004 1704041 cert.go:229] Generate etcd-ca certificate success.
I0120 23:54:07.952218 1704041 cert.go:229] Generate etcd-server certificate success.
I0120 23:54:08.162500 1704041 cert.go:229] Generate etcd-client certificate success.
I0120 23:54:08.162727 1704041 deploy.go:307] local crds file name: /root/cloud-native-project-practice/karmada/crds.tar.gz
I0120 23:54:08.167863 1704041 deploy.go:535] Create karmada kubeconfig success.
I0120 23:54:08.173540 1704041 idempotency.go:251] Namespace karmada-system has been created or updated.
I0120 23:54:08.198377 1704041 idempotency.go:275] Service karmada-system/etcd has been created or updated.
I0120 23:54:08.198400 1704041 deploy.go:363] Create etcd StatefulSets

W0120 23:56:08.224186 1704041 deploy.go:369] wait for Statefulset(karmada-system/etcd) rollout: timed out waiting for the condition: expected 1 replicas, got 0 available replicas
I0120 23:56:08.224204 1704041 deploy.go:371] Create karmada ApiServer Deployment
I0120 23:56:08.231656 1704041 idempotency.go:275] Service karmada-system/karmada-apiserver has been created or updated.
I0120 23:57:48.243226 1704041 deploy.go:386] Create karmada aggregated apiserver Deployment
I0120 23:57:48.252148 1704041 idempotency.go:275] Service karmada-system/karmada-aggregated-apiserver has been created or updated.

I0120 23:59:36.276314 1704041 idempotency.go:251] Namespace karmada-system has been created or updated.
I0120 23:59:36.276578 1704041 deploy.go:68] Initialize karmada bases crd resource `/etc/karmada/crds/bases`
I0120 23:59:36.281915 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.290243 1704041 deploy.go:233] Create CRD cronfederatedhpas.autoscaling.karmada.io successfully.
I0120 23:59:36.299050 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.317475 1704041 deploy.go:233] Create CRD federatedhpas.autoscaling.karmada.io successfully.
I0120 23:59:36.318546 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.324911 1704041 deploy.go:233] Create CRD resourceinterpretercustomizations.config.karmada.io successfully.
I0120 23:59:36.326253 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.332557 1704041 deploy.go:233] Create CRD resourceinterpreterwebhookconfigurations.config.karmada.io successfully.
I0120 23:59:36.333402 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.340756 1704041 deploy.go:233] Create CRD serviceexports.multicluster.x-k8s.io successfully.
I0120 23:59:36.341886 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.353342 1704041 deploy.go:233] Create CRD serviceimports.multicluster.x-k8s.io successfully.
I0120 23:59:36.355290 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.365919 1704041 deploy.go:233] Create CRD multiclusteringresses.networking.karmada.io successfully.
I0120 23:59:36.367354 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.378604 1704041 deploy.go:233] Create CRD multiclusterservices.networking.karmada.io successfully.
I0120 23:59:36.385581 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.405790 1704041 deploy.go:233] Create CRD clusteroverridepolicies.policy.karmada.io successfully.
I0120 23:59:36.409627 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.475680 1704041 deploy.go:233] Create CRD clusterpropagationpolicies.policy.karmada.io successfully.
I0120 23:59:36.476922 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.494975 1704041 deploy.go:233] Create CRD federatedresourcequotas.policy.karmada.io successfully.
I0120 23:59:36.500144 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.698015 1704041 deploy.go:233] Create CRD overridepolicies.policy.karmada.io successfully.
I0120 23:59:36.702785 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:36.893455 1704041 deploy.go:233] Create CRD propagationpolicies.policy.karmada.io successfully.
I0120 23:59:36.900448 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:37.102606 1704041 deploy.go:233] Create CRD clusterresourcebindings.work.karmada.io successfully.
I0120 23:59:37.108932 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:37.299045 1704041 deploy.go:233] Create CRD resourcebindings.work.karmada.io successfully.
I0120 23:59:37.300686 1704041 deploy.go:223] Attempting to create CRD
I0120 23:59:37.489052 1704041 deploy.go:233] Create CRD works.work.karmada.io successfully.
I0120 23:59:37.489137 1704041 deploy.go:79] Initialize karmada patches crd resource `/etc/karmada/crds/patches`
I0120 23:59:37.909586 1704041 deploy.go:91] Create MutatingWebhookConfiguration mutating-config.
I0120 23:59:37.919138 1704041 webhook_configuration.go:315] MutatingWebhookConfiguration mutating-config has been created or updated successfully.
I0120 23:59:37.919159 1704041 deploy.go:96] Create ValidatingWebhookConfiguration validating-config.
I0120 23:59:37.928038 1704041 webhook_configuration.go:286] ValidatingWebhookConfiguration validating-config has been created or updated successfully.
I0120 23:59:37.928058 1704041 deploy.go:102] Create Service 'karmada-aggregated-apiserver' and APIService 'v1alpha1.cluster.karmada.io'.
I0120 23:59:37.931088 1704041 idempotency.go:275] Service karmada-system/karmada-aggregated-apiserver has been created or updated.
I0120 23:59:37.937074 1704041 check.go:26] Waiting for APIService(v1alpha1.cluster.karmada.io) condition(Available), will try
I0120 23:59:38.969126 1704041 tlsbootstrap.go:33] [bootstrap-token] configured RBAC rules to allow Karmada Agent Bootstrap tokens to post CSRs in order for agent to get long term certificate credentials
I0120 23:59:38.971277 1704041 tlsbootstrap.go:47] [bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Karmada Agent Bootstrap Token
I0120 23:59:38.973999 1704041 tlsbootstrap.go:61] [bootstrap-token] configured RBAC rules to allow certificate rotation for all agent client certificates in the member cluster
I0120 23:59:39.145647 1704041 deploy.go:126] Initialize karmada bootstrap token
I0120 23:59:39.152375 1704041 deploy.go:404] Create karmada kube controller manager Deployment
I0120 23:59:39.161661 1704041 idempotency.go:275] Service karmada-system/kube-controller-manager has been created or updated.





W0121 00:01:39.175274 1704041 deploy.go:413] wait for Deployment(karmada-system/kube-controller-manager) rollout: timed out waiting for the condition: expected 1 replicas, got 0 available replicas
I0121 00:01:39.175291 1704041 deploy.go:418] Create karmada scheduler Deployment


W0121 00:03:39.189244 1704041 deploy.go:424] wait for Deployment(karmada-system/karmada-scheduler) rollout: timed out waiting for the condition: expected 1 replicas, got 0 available replicas
I0121 00:03:39.189262 1704041 deploy.go:429] Create karmada controller manager Deployment



W0121 00:05:39.204464 1704041 deploy.go:435] wait for Deployment(karmada-system/karmada-controller-manager) rollout: timed out waiting for the condition: expected 1 replicas, got 0 available replicas
I0121 00:05:39.204484 1704041 deploy.go:440] Create karmada webhook Deployment
I0121 00:05:39.217192 1704041 idempotency.go:275] Service karmada-system/karmada-webhook has been created or updated.

W0121 00:07:39.236086 1704041 deploy.go:449] wait for Deployment(karmada-system/karmada-webhook) rollout: timed out waiting for the condition: expected 1 replicas, got 0 available replicas

------------------------------------------------------------------------------------------------------
 █████   ████   █████████   ███████████   ██████   ██████   █████████   ██████████     █████████
░░███   ███░   ███░░░░░███ ░░███░░░░░███ ░░██████ ██████   ███░░░░░███ ░░███░░░░███   ███░░░░░███
 ░███  ███    ░███    ░███  ░███    ░███  ░███░█████░███  ░███    ░███  ░███   ░░███ ░███    ░███
 ░███████     ░███████████  ░██████████   ░███░░███ ░███  ░███████████  ░███    ░███ ░███████████
 ░███░░███    ░███░░░░░███  ░███░░░░░███  ░███ ░░░  ░███  ░███░░░░░███  ░███    ░███ ░███░░░░░███
 ░███ ░░███   ░███    ░███  ░███    ░███  ░███      ░███  ░███    ░███  ░███    ███  ░███    ░███
 █████ ░░████ █████   █████ █████   █████ █████     █████ █████   █████ ██████████   █████   █████
░░░░░   ░░░░ ░░░░░   ░░░░░ ░░░░░   ░░░░░ ░░░░░     ░░░░░ ░░░░░   ░░░░░ ░░░░░░░░░░   ░░░░░   ░░░░░
------------------------------------------------------------------------------------------------------
Karmada is installed successfully.

Register Kubernetes cluster to Karmada control plane.

Register cluster with 'Push' mode

Step 1: Use "karmadactl join" command to register the cluster to Karmada control plane. --cluster-kubeconfig is kubeconfig of the member cluster.
(In karmada)~# MEMBER_CLUSTER_NAME=$(cat ~/.kube/config  | grep current-context | sed 's/: /\n/g'| sed '1d')
(In karmada)~# karmadactl --kubeconfig /etc/karmada/karmada-apiserver.config  join ${MEMBER_CLUSTER_NAME} --cluster-kubeconfig=$HOME/.kube/config

Step 2: Show members of karmada
(In karmada)~# kubectl --kubeconfig /etc/karmada/karmada-apiserver.config get clusters


Register cluster with 'Pull' mode

Step 1: Use "karmadactl register" command to register the cluster to Karmada control plane. "--cluster-name" is set to cluster of current-context by default.
(In member cluster)~# karmadactl register 172.18.0.2:31443 --token 4t7wrz.i306e1fkzkgpi3tt --discovery-token-ca-cert-hash sha256:5081cd2db52b1de7b6103db9caa2c798e7cb0d213e7f6eb9eb64742dd655a3b4

Step 2: Show members of karmada
(In karmada)~# kubectl --kubeconfig /etc/karmada/karmada-apiserver.config get clusters


```
- 常用命令
```bash
# 加入集群
# --kubeconfig 是 karmada 配置文件
# --cluster-kubeconfig k8s 配置文件
# --cluster-context 多个集群会有自己的 context, kind 默认会有 kind 前缀
karmadactl join <cluster name> \
--kubeconfig /etc/karmada/karmada-apiserver.config \
--cluster-kubeconfig='/root/.kube/config' \
--cluster-context='kind-<cluster name>'

[root@VM-0-13-centos .kube]# karmadactl join cluster2 --kubeconfig /etc/karmada/karmada-apiserver.config --cluster-kubeconfig='/root/.kube/config' --cluster-context='kind-cluster2'
cluster(cluster2) is joined successfully


# 删除集群
karmadactl --kubeconfig /etc/karmada/karmada-apiserver.config unjoin cluster-<cluster name>

[root@VM-0-13-centos .kube]# karmadactl --kubeconfig /etc/karmada/karmada-apiserver.config unjoin cluster-cluster2
I0121 09:43:08.099361 2105496 unjoin.go:235] Waiting for the cluster object cluster-cluster2 to be deleted
I0121 09:43:09.099367 2105496 unjoin.go:235] Waiting for the cluster object cluster-cluster2 to be deleted
I0121 09:43:10.099649 2105496 unjoin.go:235] Waiting for the cluster object cluster-cluster2 to be deleted
I0121 09:43:11.099394 2105496 unjoin.go:235] Waiting for the cluster object cluster-cluster2 to be deleted

# 查询集群
kubectl --kubeconfig /etc/karmada/karmada-apiserver.config get clusters

[root@VM-0-13-centos .kube]# kubectl --kubeconfig /etc/karmada/karmada-apiserver.config get clusters
NAME             VERSION    MODE   READY   AGE
cluster2         v1.20.15   Push   True    12s
master-cluster   v1.20.15   Push   True    9h
```