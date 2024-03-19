2909  2024-03-16 13:42:14 wget https://github.com/submariner-io/releases/releases/download/v0.16.2/subctl-v0.16.2-linux-amd64.tar.gz
2927  2024-03-16 14:25:14  tar -zxvf subctl-v0.16.2-linux-amd64.tar.gz
2929  2024-03-16 14:25:19 cd subctl-v0.16.2/
2931  2024-03-16 14:25:23 ./subctl
2932  2024-03-16 14:25:26 ./subctl version
2939  2024-03-16 14:25:51 cd subctl-v0.16.2/
2941  2024-03-16 14:25:58 chmod +x subctl
2942  2024-03-16 14:26:06 mv subctl /usr/local/bin
2943  2024-03-16 14:26:11 subctl version
2944  2024-03-16 14:26:48 subctl --context
2949  2024-03-16 14:28:06 subctl --context kind-jt1 deploy-broker
2952  2024-03-16 14:30:50 subctl show connections
2953  2024-03-16 14:31:03 subctl show connections --context kind-master-cluster
2954  2024-03-16 14:35:47 subctl --context kind-jt1 join broker-info.subm --clusterid jt1 --version v0.13.0
2956  2024-03-16 14:36:27 subctl --context kind-jt2 join broker-info.subm --clusterid jt2 --version v0.13.0
2968  2024-03-16 14:47:35 subctl version
2977  2024-03-16 14:55:21 subctl show connections --context kind-jt1
2978  2024-03-16 14:55:49 subctl
2979  2024-03-16 14:57:40 history | grep subctl


[root@VM-0-13-centos subctl-v0.16.2]# cd ..
[root@VM-0-13-centos submariner]# cd
[root@VM-0-13-centos ~]# subctl --context kind-master-cluster deploy-broker
✓ Setting up broker RBAC
⠊⠁ Deploying the Submariner operator
✓ Deploying the Submariner operator
✓ Created operator CRDs
✓ Created operator namespace: submariner-operator
✓ Created operator service account and role
✓ Created submariner service account and role
✓ Created lighthouse service account and role
✓ Deployed the operator successfully
✓ Deploying the broker
✓ Saving broker info to file "broker-info.subm"


[root@VM-0-13-centos ~]# subctl show connections
Cluster "kind-master-cluster"
⚠ Submariner connectivity feature is not installed

Cluster "kind-cluster2"
⚠ Submariner connectivity feature is not installed

[root@VM-0-13-centos ~]# subctl --context kind-master-cluster deploy-broker^C
[root@VM-0-13-centos ~]# subctl show connections --context kind-master-cluster
⚠ Submariner connectivity feature is not installed
[root@VM-0-13-centos ~]# subctl --context kind-master-cluster join broker-info.subm --clusterid master-cluster
✓ broker-info.subm indicates broker is at https://172.18.0.3:6443
✓ Discovering network details
Network plugin:  kindnet
Service CIDRs:   [10.96.0.0/12]
Cluster CIDRs:   [10.244.0.0/16]
✓ Retrieving the gateway nodes
✓ Retrieving all worker nodes
⚠ No worker node available to label as the gateway
✓ Gathering relevant information from Broker
✓ Retrieving Globalnet information from the Broker
✓ Validating Globalnet configuration
✓ Deploying the Submariner operator
✓ Created operator namespace: submariner-operator
✓ Creating SA for cluster
✓ Connecting to Broker
✓ Deploying submariner
✓ Submariner is up and running
[root@VM-0-13-centos ~]# kind get clusters
cluster2
master-cluster
[root@VM-0-13-centos ~]# subctl --context kind-cluster2 join broker-info.subm --clusterid cluster2
✓ broker-info.subm indicates broker is at https://172.18.0.3:6443
✓ Discovering network details
Network plugin:  kindnet
Service CIDRs:   [10.96.0.0/12]
Cluster CIDRs:   [10.244.0.0/16]
✓ Retrieving the gateway nodes
✓ Retrieving all worker nodes
✓ Labeling node "cluster2-worker" as a gateway
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