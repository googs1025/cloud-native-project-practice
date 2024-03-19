```bash
2909  2024-03-16 13:42:14 wget https://github.com/submariner-io/releases/releases/download/v0.16.2/subctl-v0.16.2-linux-amd64.tar.gz
2927  2024-03-16 14:25:14  tar -zxvf subctl-v0.16.2-linux-amd64.tar.gz
2929  2024-03-16 14:25:19 cd subctl-v0.16.2/
2931  2024-03-16 14:25:23 ./subctl
2932  2024-03-16 14:25:26 ./subctl version
2939  2024-03-16 14:25:51 cd subctl-v0.16.2/
2941  2024-03-16 14:25:58 chmod +x subctl
2942  2024-03-16 14:26:06 mv subctl /usr/local/bin

root@VM-0-12-ubuntu:~# kind create cluster --config=jt1.yaml
Creating cluster "jt1" ...
 ✓ Ensuring node image (kindest/node:v1.24.15) 🖼
 ✓ Preparing nodes 📦 📦 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
 ✓ Joining worker nodes 🚜
Set kubectl context to "kind-jt1"
You can now use your cluster with:

kubectl cluster-info --context kind-jt1

Not sure what to do next? 😅  Check out https://kind.sigs.k8s.io/docs/user/quick-start/
root@VM-0-12-ubuntu:~# kubectl get pods
No resources found in default namespace.
root@VM-0-12-ubuntu:~# kind delete clusters jt2
Deleted nodes: ["jt2-control-plane" "jt2-worker"]
Deleted clusters: ["jt2"]
root@VM-0-12-ubuntu:~# kind create cluster --config=jt2.yaml
Creating cluster "jt2" ...
 ✓ Ensuring node image (kindest/node:v1.24.15) 🖼
 ✓ Preparing nodes 📦 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
 ✓ Joining worker nodes 🚜
Set kubectl context to "kind-jt2"
You can now use your cluster with:

kubectl cluster-info --context kind-jt2

Thanks for using kind! 😊
root@VM-0-12-ubuntu:~# kubectl get pods -a
error: unknown shorthand flag: 'a' in -a
See 'kubectl get --help' for usage.
root@VM-0-12-ubuntu:~# kubectl get pods -A
NAMESPACE            NAME                                        READY   STATUS    RESTARTS   AGE
kube-system          coredns-57575c5f89-v7wqn                    1/1     Running   0          43s
kube-system          coredns-57575c5f89-xjs76                    1/1     Running   0          43s
kube-system          etcd-jt2-control-plane                      1/1     Running   0          56s
kube-system          kindnet-c5x7c                               1/1     Running   0          39s
kube-system          kindnet-twt9k                               1/1     Running   0          43s
kube-system          kube-apiserver-jt2-control-plane            1/1     Running   0          56s
kube-system          kube-controller-manager-jt2-control-plane   1/1     Running   0          58s
kube-system          kube-proxy-24lhb                            1/1     Running   0          43s
kube-system          kube-proxy-9gtbh                            1/1     Running   0          39s
kube-system          kube-scheduler-jt2-control-plane            1/1     Running   0          56s
local-path-storage   local-path-provisioner-c49b7b56f-s44qv      1/1     Running   0          43s
root@VM-0-12-ubuntu:~# subctl
Deploy, manage, verify and diagnose Submariner deployments

Usage:
  subctl [command]

Available Commands:
  benchmark           Benchmark tests
  cloud               Cloud operations
  completion          Generate the autocompletion script for the specified shell
  deploy-broker       Deploys the broker
  diagnose            Run diagnostic checks on the Submariner deployment and report any issues
  export              Exports a resource to other clusters
  gather              Gather troubleshooting information from a cluster
  help                Help about any command
  join                Connect a cluster to an existing broker
  recover-broker-info Recovers the broker-info.subm file from the installed Broker
  show                Show information about Submariner
  unexport            Stop a resource from being exported to other clusters
  uninstall           Uninstall Submariner and its components
  upgrade             Upgrades Submariner
  verify              Run verifications between two clusters
  version             Get version information on subctl

Flags:
  -h, --help   help for subctl

Use "subctl [command] --help" for more information about a command.
root@VM-0-12-ubuntu:~# subctl --context kind-jt1 deploy-broker
 ✓ Setting up broker RBAC
 ✓ Deploying the Submariner operator
 ✓ Created operator CRDs
 ✓ Created operator namespace: submariner-operator
 ✓ Created operator service account and role
 ✓ Created submariner service account and role
 ✓ Created lighthouse service account and role
 ✓ Deployed the operator successfully
 ✓ Deploying the broker
 ✓ Saving broker info to file "broker-info.subm"
root@VM-0-12-ubuntu:~# ls
broker-info.subm  jt1.yaml  jt2.yaml  kind  snap
root@VM-0-12-ubuntu:~# subctl --context kind-jt1 join broker-info.subm --clusterid jt1
 ✓ broker-info.subm indicates broker is at https://172.19.0.12:6443
 ✓ Discovering network details
        Network plugin:  kindnet
        Service CIDRs:   [10.96.0.0/16]
        Cluster CIDRs:   [10.6.0.0/16]
 ✓ Retrieving the gateway nodes
 ✓ Retrieving all worker nodes
? Which node should be used as the gateway? jt1-worker
 ✓ Labeling node "jt1-worker" as a gateway
 ✓ Gathering relevant information from Broker
 ✓ Retrieving Globalnet information from the Broker
 ✓ Validating Globalnet configuration
 ✓ Deploying the Submariner operator
 ✓ Created operator namespace: submariner-operator
 ✓ Creating SA for cluster
 ✓ Connecting to Broker
 ✓ Deploying submariner
 ✓ Submariner is up and running
root@VM-0-12-ubuntu:~# kubectl get pods -A --context kind-jt1
NAMESPACE             NAME                                             READY   STATUS              RESTARTS   AGE
kube-system           coredns-57575c5f89-mk2lc                         1/1     Running             0          3m31s
kube-system           coredns-57575c5f89-t6vn7                         1/1     Running             0          3m31s
kube-system           etcd-jt1-control-plane                           1/1     Running             0          3m43s
kube-system           kindnet-lg4fq                                    1/1     Running             0          3m26s
kube-system           kindnet-rm2d4                                    1/1     Running             0          3m26s
kube-system           kindnet-sv86l                                    1/1     Running             0          3m31s
kube-system           kube-apiserver-jt1-control-plane                 1/1     Running             0          3m43s
kube-system           kube-controller-manager-jt1-control-plane        1/1     Running             0          3m43s
kube-system           kube-proxy-kw4g8                                 1/1     Running             0          3m31s
kube-system           kube-proxy-n89ph                                 1/1     Running             0          3m26s
kube-system           kube-proxy-zstrb                                 1/1     Running             0          3m26s
kube-system           kube-scheduler-jt1-control-plane                 1/1     Running             0          3m44s
local-path-storage    local-path-provisioner-c49b7b56f-g2ckg           1/1     Running             0          3m31s
submariner-operator   submariner-gateway-6tb5t                         0/1     ContainerCreating   0          13s
submariner-operator   submariner-lighthouse-agent-6876c7986b-jwz65     0/1     ContainerCreating   0          13s
submariner-operator   submariner-lighthouse-coredns-8668dc5c45-g5qs8   0/1     ContainerCreating   0          12s
submariner-operator   submariner-lighthouse-coredns-8668dc5c45-tw5d9   0/1     ContainerCreating   0          12s
submariner-operator   submariner-metrics-proxy-96pqx                   0/1     ContainerCreating   0          13s
submariner-operator   submariner-operator-f8b9cdbbf-n7x92              1/1     Running             0          44s
submariner-operator   submariner-routeagent-dpb8m                      0/1     Init:0/1            0          13s
submariner-operator   submariner-routeagent-lqvcq                      0/1     Init:0/1            0          13s
submariner-operator   submariner-routeagent-qvdvw                      0/1     Init:0/1            0          13s
root@VM-0-12-ubuntu:~# subctl --context kind-jt2 join broker-info.subm --clusterid jt2
 ✓ broker-info.subm indicates broker is at https://172.19.0.12:6443
 ✓ Discovering network details
        Network plugin:  kindnet
        Service CIDRs:   [10.97.0.0/16]
        Cluster CIDRs:   [10.7.0.0/16]
 ✓ Retrieving the gateway nodes
 ✓ Retrieving all worker nodes
? Which node should be used as the gateway? jt2-worker
 ✓ Labeling node "jt2-worker" as a gateway
 ✓ Gathering relevant information from Broker
 ✓ Retrieving Globalnet information from the Broker
 ✓ Validating Globalnet configuration
 ✓ Deploying the Submariner operator
 ✓ Created operator CRDs
 ✓ Created operator namespace: submariner-operator
 ✓ Created operator service account and role
 ✓ Created submariner service account and role
 ✓ Created lighthouse service account and role
 ✓ Deployed the operator successfully
 ✓ Creating SA for cluster
 ✓ Connecting to Broker
 ✓ Deploying submariner
 ✓ Submariner is up and running
root@VM-0-12-ubuntu:~# subctl --context kind-jt2 join broker-info.subm --clusterid jt2
 ✓ broker-info.subm indicates broker is at https://172.19.0.12:6443
 ✓ Discovering network details
        Network plugin:  kindnet
        Service CIDRs:   [10.97.0.0/16]
        Cluster CIDRs:   [10.7.0.0/16]
   There are 1 node(s) labeled as gateways:
    - jt2-worker
 ✓ Retrieving the gateway nodes
 ✓ Gathering relevant information from Broker
 ✓ Retrieving Globalnet information from the Broker
 ✓ Validating Globalnet configuration
⠈⠱ Deploying the Submariner operator ^C
root@VM-0-12-ubuntu:~# kubectl get pods -A --context kind-jt1
NAMESPACE             NAME                                             READY   STATUS    RESTARTS   AGE
kube-system           coredns-57575c5f89-mk2lc                         1/1     Running   0          4m7s
kube-system           coredns-57575c5f89-t6vn7                         1/1     Running   0          4m7s
kube-system           etcd-jt1-control-plane                           1/1     Running   0          4m19s
kube-system           kindnet-lg4fq                                    1/1     Running   0          4m2s
kube-system           kindnet-rm2d4                                    1/1     Running   0          4m2s
kube-system           kindnet-sv86l                                    1/1     Running   0          4m7s
kube-system           kube-apiserver-jt1-control-plane                 1/1     Running   0          4m19s
kube-system           kube-controller-manager-jt1-control-plane        1/1     Running   0          4m19s
kube-system           kube-proxy-kw4g8                                 1/1     Running   0          4m7s
kube-system           kube-proxy-n89ph                                 1/1     Running   0          4m2s
kube-system           kube-proxy-zstrb                                 1/1     Running   0          4m2s
kube-system           kube-scheduler-jt1-control-plane                 1/1     Running   0          4m20s
local-path-storage    local-path-provisioner-c49b7b56f-g2ckg           1/1     Running   0          4m7s
submariner-operator   submariner-gateway-6tb5t                         1/1     Running   0          49s
submariner-operator   submariner-lighthouse-agent-6876c7986b-jwz65     1/1     Running   0          49s
submariner-operator   submariner-lighthouse-coredns-8668dc5c45-g5qs8   1/1     Running   0          48s
submariner-operator   submariner-lighthouse-coredns-8668dc5c45-tw5d9   1/1     Running   0          48s
submariner-operator   submariner-metrics-proxy-96pqx                   1/1     Running   0          49s
submariner-operator   submariner-operator-f8b9cdbbf-n7x92              1/1     Running   0          80s
submariner-operator   submariner-routeagent-dpb8m                      1/1     Running   0          49s
submariner-operator   submariner-routeagent-lqvcq                      1/1     Running   0          49s
submariner-operator   submariner-routeagent-qvdvw                      1/1     Running   0          49s
root@VM-0-12-ubuntu:~# kubectl get pods -A --context kind-jt2
NAMESPACE             NAME                                             READY   STATUS              RESTARTS   AGE
kube-system           coredns-57575c5f89-v7wqn                         1/1     Running             0          2m43s
kube-system           coredns-57575c5f89-xjs76                         1/1     Running             0          2m43s
kube-system           etcd-jt2-control-plane                           1/1     Running             0          2m56s
kube-system           kindnet-c5x7c                                    1/1     Running             0          2m39s
kube-system           kindnet-twt9k                                    1/1     Running             0          2m43s
kube-system           kube-apiserver-jt2-control-plane                 1/1     Running             0          2m56s
kube-system           kube-controller-manager-jt2-control-plane        1/1     Running             0          2m58s
kube-system           kube-proxy-24lhb                                 1/1     Running             0          2m43s
kube-system           kube-proxy-9gtbh                                 1/1     Running             0          2m39s
kube-system           kube-scheduler-jt2-control-plane                 1/1     Running             0          2m56s
local-path-storage    local-path-provisioner-c49b7b56f-s44qv           1/1     Running             0          2m43s
submariner-operator   submariner-gateway-tpkq6                         1/1     Running             0          13s
submariner-operator   submariner-lighthouse-agent-7b54f8cdcd-x2j9c     0/1     ContainerCreating   0          13s
submariner-operator   submariner-lighthouse-coredns-7f5fbc6b68-6hn4d   0/1     ContainerCreating   0          12s
submariner-operator   submariner-lighthouse-coredns-7f5fbc6b68-fgg56   0/1     ContainerCreating   0          12s
submariner-operator   submariner-metrics-proxy-pkzh5                   0/1     ContainerCreating   0          13s
submariner-operator   submariner-operator-f8b9cdbbf-k7ltl              1/1     Running             0          19s
submariner-operator   submariner-routeagent-ftdcv                      0/1     Init:0/1            0          13s
submariner-operator   submariner-routeagent-wfxzf                      0/1     Init:0/1            0          13s
root@VM-0-12-ubuntu:~# kubectl get pods -A --context kind-jt2
NAMESPACE             NAME                                             READY   STATUS              RESTARTS   AGE
kube-system           coredns-57575c5f89-v7wqn                         1/1     Running             0          3m13s
kube-system           coredns-57575c5f89-xjs76                         1/1     Running             0          3m13s
kube-system           etcd-jt2-control-plane                           1/1     Running             0          3m26s
kube-system           kindnet-c5x7c                                    1/1     Running             0          3m9s
kube-system           kindnet-twt9k                                    1/1     Running             0          3m13s
kube-system           kube-apiserver-jt2-control-plane                 1/1     Running             0          3m26s
kube-system           kube-controller-manager-jt2-control-plane        1/1     Running             0          3m28s
kube-system           kube-proxy-24lhb                                 1/1     Running             0          3m13s
kube-system           kube-proxy-9gtbh                                 1/1     Running             0          3m9s
kube-system           kube-scheduler-jt2-control-plane                 1/1     Running             0          3m26s
local-path-storage    local-path-provisioner-c49b7b56f-s44qv           1/1     Running             0          3m13s
submariner-operator   submariner-gateway-tpkq6                         1/1     Running             0          43s
submariner-operator   submariner-lighthouse-agent-7b54f8cdcd-x2j9c     1/1     Running             0          43s
submariner-operator   submariner-lighthouse-coredns-7f5fbc6b68-6hn4d   0/1     ContainerCreating   0          42s
submariner-operator   submariner-lighthouse-coredns-7f5fbc6b68-fgg56   0/1     ContainerCreating   0          42s
submariner-operator   submariner-metrics-proxy-pkzh5                   1/1     Running             0          43s
submariner-operator   submariner-operator-f8b9cdbbf-k7ltl              1/1     Running             0          49s
submariner-operator   submariner-routeagent-ftdcv                      1/1     Running             0          43s
submariner-operator   submariner-routeagent-wfxzf                      1/1     Running             0          43s
root@VM-0-12-ubuntu:~# kubectl get pods -A --context kind-jt2
NAMESPACE             NAME                                             READY   STATUS              RESTARTS   AGE
kube-system           coredns-57575c5f89-v7wqn                         1/1     Running             0          3m16s
kube-system           coredns-57575c5f89-xjs76                         1/1     Running             0          3m16s
kube-system           etcd-jt2-control-plane                           1/1     Running             0          3m29s
kube-system           kindnet-c5x7c                                    1/1     Running             0          3m12s
kube-system           kindnet-twt9k                                    1/1     Running             0          3m16s
kube-system           kube-apiserver-jt2-control-plane                 1/1     Running             0          3m29s
kube-system           kube-controller-manager-jt2-control-plane        1/1     Running             0          3m31s
kube-system           kube-proxy-24lhb                                 1/1     Running             0          3m16s
kube-system           kube-proxy-9gtbh                                 1/1     Running             0          3m12s
kube-system           kube-scheduler-jt2-control-plane                 1/1     Running             0          3m29s
local-path-storage    local-path-provisioner-c49b7b56f-s44qv           1/1     Running             0          3m16s
submariner-operator   submariner-gateway-tpkq6                         1/1     Running             0          46s
submariner-operator   submariner-lighthouse-agent-7b54f8cdcd-x2j9c     1/1     Running             0          46s
submariner-operator   submariner-lighthouse-coredns-7f5fbc6b68-6hn4d   1/1     Running             0          45s
submariner-operator   submariner-lighthouse-coredns-7f5fbc6b68-fgg56   0/1     ContainerCreating   0          45s
submariner-operator   submariner-metrics-proxy-pkzh5                   1/1     Running             0          46s
submariner-operator   submariner-operator-f8b9cdbbf-k7ltl              1/1     Running             0          52s
submariner-operator   submariner-routeagent-ftdcv                      1/1     Running             0          46s
submariner-operator   submariner-routeagent-wfxzf                      1/1     Running             0          46s
root@VM-0-12-ubuntu:~# subctl show connections --context kind-jt1
 ✓ Showing Connections
GATEWAY      CLUSTER   REMOTE IP    NAT   CABLE DRIVER   SUBNETS                     STATUS      RTT avg.
jt2-worker   jt2       172.18.0.5   no    libreswan      10.97.0.0/16, 10.7.0.0/16   connected   205.485µs
root@VM-0-12-ubuntu:~# subctl show connections --context kind-jt2
 ✓ Showing Connections
GATEWAY      CLUSTER   REMOTE IP    NAT   CABLE DRIVER   SUBNETS                     STATUS      RTT avg.
jt1-worker   jt1       172.18.0.4   no    libreswan      10.96.0.0/16, 10.6.0.0/16   connected   277.936µs

```