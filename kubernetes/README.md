# Kubernetes : Container Orchastretion

<img src='https://d33wubrfki0l68.cloudfront.net/2475489eaf20163ec0f54ddc1d92aa8d4c87c96b/e7c81/images/docs/components-of-kubernetes.svg' style="background-color:white" />

- [Minikube basics](https://minikube.sigs.k8s.io/docs/start/)

- [k8s Concept](https://kubernetes.io/docs/concepts/)

- [kubetctl cheatsheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

- [kubetctl cheatsheet 1](https://design.jboss.org/redhatdeveloper/marketing/kubernetes_cheatsheet/cheatsheet/cheat_sheet/images/kubernetes_cheat_sheet_r1v1.pdf)

- [kubetctl cheatsheet 2](https://github.com/RehanSaeed/Kubernetes-Cheat-Sheet)

- [yaml Templates](https://github.com/dennyzhang/kubernetes-yaml-templates/)

- [API reference](https://kubernetes.io/docs/reference/kubernetes-api/) 

- [API Convention](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)

<details>

<summary><h2>Pods</h2></summary>

<img src="https://3p8owy1gdkoh452nrc36wbnp-wpengine.netdna-ssl.com/wp-content/uploads/2018/02/Pod-Diagram-1-700x360.png">
</img>


### [Concept](https://kubernetes.io/docs/concepts/workloads/pods/) 

### [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/pod-v1/ )

### Ex
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: http
spec:
  containers:
  - name: http
    image: ghcr.io/rakshith-r/http_server:latest
    ports:
    - containerPort: 3000
```
</details>


<details>
<summary><h2>Replication Controller </h2></summary>


### [Concept]("https://kubernetes.io/docs/concepts/workloads/controllers/replicacontroller/") 

### [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/replica-controller-v1/)


### Ex

```yaml
apiVersion: apps/v1
kind: ReplicationController
metadata:
  name: server
spec:
  selector:
    app: http
  replicas: 5
  template:
    metadata:
      name: http
      labels:
        app: http
    spec:
      containers:
      - name: http
        image: ghcr.io/rakshith-r/http_server:latest
        ports:
        - containerPort: 3000
```
</details>


<details>
<summary><h2>ReplicaSets </h2></summary>

ReplicaSet is the next-generation ReplicationController that supports the new set-based label selector. It's mainly used by Deployment as a mechanism to orchestrate pod creation, deletion and updates. Note that we recommend using Deployments instead of directly using Replica Sets, unless you require custom update orchestration or don't require updates at all.


### [Concept]("https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/") 

### [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/replica-set-v1/)


### Ex

```yaml
apiVersion: apps/v1
kind: Replicaset
metadata:
  name: server
spec:
  selector:
    matchLabels:
      app: http
  replicas: 5
  template:
    metadata:
      labels:
        app: http
    spec:
      containers:
      - name: http
        image: ghcr.io/rakshith-r/http_server:latest
        ports:
        - containerPort: 3000
```
</details>


<details>
<summary><h2>Deployments </h2></summary>
(Recommended) 
Deployment is a higher-level API object that updates its underlying Replica Sets and their Pods. Deployments are recommended if you want this rolling update functionality because, they are declarative, server-side, and have additional features.


<img src="https://3p8owy1gdkoh452nrc36wbnp-wpengine.netdna-ssl.com/wp-content/uploads/2018/03/Deployment.png" style='background-color:white'></img>

### [Concept]("https://kubernetes.io/docs/concepts/workloads/controllers/deployment/") 

### [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/deployment-v1/)

### Ex

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: server
spec:
  selector:
    matchLabels:
      app: http
  replicas: 5
  template:
    metadata:
      labels:
        app: http
    spec:
      containers:
      - name: http
        image: ghcr.io/rakshith-r/http_server:latest
        ports:
        - containerPort: 3000
      strategy:
        type: RollingUpdate #or recreate
        maxUnavailable: 1
        maxSurge: 2
```
</details>


<details>
<summary><h2>DaemonSet </h2></summary>

### [Concept](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) 

### [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/daemon-set-v1/)

### Ex

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: fluentd-elasticsearch
  namespace: kube-system
  labels:
    k8s-app: fluentd-logging
spec:
  selector:
    matchLabels:
      name: fluentd-elasticsearch
  template:
    metadata:
      labels:
        name: fluentd-elasticsearch
    spec:
      containers:
      - name: fluentd-elasticsearch
        image: quay.io/fluentd_elasticsearch/fluentd:v2.5.2
        resources:
          limits:
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 200Mi
        volumeMounts:
        - name: varlog
          mountPath: /var/log
        - name: varlibdockercontainers
          mountPath: /var/lib/docker/containers
          readOnly: true
      terminationGracePeriodSeconds: 30
      volumes:
      - name: varlog
        hostPath:
          path: /var/log
      - name: varlibdockercontainers
        hostPath:
          path: /var/lib/docker/containers
```
</details>



<details>
<summary><h2>Stateful Set </h2></summary>

### [Concept](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/) 

### [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/stateful-set-v1/)

### Ex

```yaml
```
</details>