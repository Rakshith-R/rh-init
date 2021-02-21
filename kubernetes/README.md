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

<details><summary>

### Pods 
[Concept](https://kubernetes.io/docs/concepts/workloads/pods/) : | : [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/pod-v1/ )

</summary>

<img src="https://3p8owy1gdkoh452nrc36wbnp-wpengine.netdna-ssl.com/wp-content/uploads/2018/02/Pod-Diagram-1-700x360.png">
</img>



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

<details><summary>

### Jobs
[Concept](https://kubernetes.io/docs/concepts/workloads/job/) : | : [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/job-v1/)

</summary>
A Job creates one or more Pods and will continue to retry execution of the Pods until a specified number of them successfully terminate. As pods successfully complete, the Job tracks the successful completions. When a specified number of successful completions is reached, the task (ie, Job) is complete. Deleting a Job will clean up the Pods it created.



### Ex
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi
spec:
  template:
    spec:
      containers:
      - name: pi
        image: perl
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(2000)"]
      restartPolicy: Never
  backoffLimit: 4
```
</details>


<details><summary>

### Replication Controller 
[Concept](https://kubernetes.io/docs/concepts/workloads/controllers/replicacontroller/) :  | : [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/replica-controller-v1/)
</summary>

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


<details><summary>

### ReplicaSets 

[Concept](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) : | : [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/replica-set-v1/)
</summary>

ReplicaSet is the next-generation ReplicationController that supports the new set-based label selector. It's mainly used by Deployment as a mechanism to orchestrate pod creation, deletion and updates. Note that we recommend using Deployments instead of directly using Replica Sets, unless you require custom update orchestration or don't require updates at all.

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


<details><summary>

### Deployments 
[Concept](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) : | : [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/deployment-v1/)
</summary>
(Recommended) 
Deployment is a higher-level API object that updates its underlying Replica Sets and their Pods. Deployments are recommended if you want this rolling update functionality because, they are declarative, server-side, and have additional features.


<img src="https://3p8owy1gdkoh452nrc36wbnp-wpengine.netdna-ssl.com/wp-content/uploads/2018/03/Deployment.png" style='background-color:white'></img>

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


<details><summary>

### DaemonSet 
[Concept](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) : | : [API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/daemon-set-v1/)

</summary>

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



<details><summary>

### ConfigMaps

[Concept](https://kubernetes.io/docs/concepts/configuration/configmap/) : - : 
[API](https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/config-map-v1/)

</summary>
A ConfigMap is an API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume.

A ConfigMap allows you to decouple environment-specific configuration from your container images, so that your applications are easily portable.

### Ex

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: game-demo
data:
  # property-like keys; each key maps to a simple value
  player_initial_lives: "3"
  ui_properties_file_name: "user-interface.properties"

  # file-like keys
  game.properties: |
    enemy.types=aliens,monsters
    player.maximum-lives=5    
  user-interface.properties: |
    color.good=purple
    color.bad=yellow
    allow.textmode=true    
---
apiVersion: v1
kind: Pod
metadata:
  name: configmap-demo-pod
spec:
  containers:
    - name: demo
      image: alpine
      command: ["sleep", "3600"]
      env:
        # Define the environment variable
        - name: PLAYER_INITIAL_LIVES # Notice that the case is different here
                                     # from the key name in the ConfigMap.
          valueFrom:
            configMapKeyRef:
              name: game-demo           # The ConfigMap this value comes from.
              key: player_initial_lives # The key to fetch.
        - name: UI_PROPERTIES_FILE_NAME
          valueFrom:
            configMapKeyRef:
              name: game-demo
              key: ui_properties_file_name
      volumeMounts:
      - name: config
        mountPath: "/config"
        readOnly: true
  volumes:
    # You set volumes at the Pod level, then mount them into containers inside that Pod
    - name: config
      configMap:
        # Provide the name of the ConfigMap you want to mount.
        name: game-demo
        # An array of keys from the ConfigMap to create as files
        items:
        - key: "game.properties"
          path: "game.properties"
        - key: "user-interface.properties"
          path: "user-interface.properties"
```
</details>



<details>
<summary>

### Secrets

[Concept](https://kubernetes.io/docs/concepts/configuration/secret/) : - : 
[API](https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/secret-v1/)

</summary>
A Secret is an object that contains a small amount of sensitive data such as a password, a token, or a key. 
The Secret resource contains two maps: data and stringData. The data field is used to store arbitrary data, encoded using base64. The stringData field is provided for convenience, and it allows you to provide Secret data as unencoded strings. The keys of data and stringData must consist of alphanumeric characters, -, _ or ..

### Ex

```yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  username1: YWRtaW4=   #echo -n 'admin' | base64 -> YWRtaW4
  password1: dDBwLVNlY3JldA==
stringData:
  username: admin
  password: t0p-Secret
...
---
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - name: mypod
    image: ghcr.io/rakshith-r/http_server:1.0
    volumeMounts:
    - name: foo
      mountPath: "/etc/foo"
      readOnly: true
  volumes:
  - name: foo
    secret:
      secretName: mysecret
...
```
</details>


<details>
<summary>

### Volumes 

[Concept](https://kubernetes.io/docs/concepts/storage/volumes/) : - : 
[API](https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/volume/)

</summary>
Volume represents a named volume in a pod that may be accessed by any container in the pod.

### Ex

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-pd
spec:
  containers:
  - image: k8s.gcr.io/test-webserver
    name: test-container
    volumeMounts:
    - mountPath: /cache
      name: cache-volume
  volumes:
  - name: cache-volume
    emptyDir: {}
```
</details>

<details>
<summary>

### Stateful Set 

[Concept](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/) : - : 
[API](https://kubernetes.io/docs/reference/kubernetes-api/workloads-resources/stateful-set-v1/)

</summary>
Manages the deployment and scaling of a set of Pods, and provides guarantees about the ordering and uniqueness of these Pods and sticky identity.

### Ex

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  serviceName: "nginx"
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```
</details>




<details>
<summary>

### Pod assignments 

[Taints & Toleration](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) : - : 
[Affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity)

</summary>
Node affinity, is a property of Pods that attracts them to a set of nodes (either as a preference or a hard requirement). Taints are the opposite -- they allow a node to repel a set of pods.

Tolerations are applied to pods, and allow (but do not require) the pods to schedule onto nodes with matching taints.

Taints and tolerations work together to ensure that pods are not scheduled onto inappropriate nodes. One or more taints are applied to a node; this marks that the node should not accept any pods that do not tolerate the taints.


### Ex
taint a node

`kubectl taint nodes node1 key1=value1:NoSchedule `

to remove Taint

`kubectl taint nodes node1 key1=value1:NoSchedule- `

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: http
spec:
  containers:
  - name: http
    image: ghcr.io/rakshith-r/http_server:1.0
    ports:
    - containerPort: 3000
  tolerations: 
  - key: "key2"
    operator: "Equal" #or "Exists"
    effect: "NoSchedule" #or PreferNoSchedule or NoExecute
    value : "value1" #not needed if opeartor is Exists
    tolerationSeconds: 6000 #for NoExecute
```

</details>