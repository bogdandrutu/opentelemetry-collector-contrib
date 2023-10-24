[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# k8s_cluster

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### k8s.container.cpu_limit

Maximum resource limit set for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {cpu} | Gauge | Double |

### k8s.container.cpu_request

Resource requested for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {cpu} | Gauge | Double |

### k8s.container.ephemeralstorage_limit

Maximum resource limit set for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### k8s.container.ephemeralstorage_request

Resource requested for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### k8s.container.memory_limit

Maximum resource limit set for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### k8s.container.memory_request

Resource requested for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### k8s.container.ready

Whether a container has passed its readiness probe (0 for no, 1 for yes)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
|  | Gauge | Int |

### k8s.container.restarts

How many times the container has restarted in the recent past. This value is pulled directly from the K8s API and the value can go indefinitely high and be reset to 0 at any time depending on how your kubelet is configured to prune dead containers. It is best to not depend too much on the exact value but rather look at it as either == 0, in which case you can conclude there were no restarts in the recent past, or > 0, in which case you can conclude there were restarts in the recent past, and not try and analyze the value beyond that.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {restart} | Gauge | Int |

### k8s.container.storage_limit

Maximum resource limit set for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### k8s.container.storage_request

Resource requested for the container. See https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcerequirements-v1-core for details

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

### k8s.cronjob.active_jobs

The number of actively running jobs for a cronjob

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {job} | Gauge | Int |

### k8s.daemonset.current_scheduled_nodes

Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {node} | Gauge | Int |

### k8s.daemonset.desired_scheduled_nodes

Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {node} | Gauge | Int |

### k8s.daemonset.misscheduled_nodes

Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {node} | Gauge | Int |

### k8s.daemonset.ready_nodes

Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {node} | Gauge | Int |

### k8s.deployment.available

Total number of available pods (ready for at least minReadySeconds) targeted by this deployment

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.deployment.desired

Number of desired pods in this deployment

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.hpa.current_replicas

Current number of pod replicas managed by this autoscaler.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.hpa.desired_replicas

Desired number of pod replicas managed by this autoscaler.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.hpa.max_replicas

Maximum number of replicas to which the autoscaler can scale up.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.hpa.min_replicas

Minimum number of replicas to which the autoscaler can scale up.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.job.active_pods

The number of actively running pods for a job

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.job.desired_successful_pods

The desired number of successfully finished pods the job should be run with

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.job.failed_pods

The number of pods which reached phase Failed for a job

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.job.max_parallel_pods

The max desired number of pods the job should run at any given time

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.job.successful_pods

The number of pods which reached phase Succeeded for a job

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.namespace.phase

The current phase of namespaces (1 for active and 0 for terminating)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
|  | Gauge | Int |

### k8s.pod.phase

Current phase of the pod (1 - Pending, 2 - Running, 3 - Succeeded, 4 - Failed, 5 - Unknown)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
|  | Gauge | Int |

### k8s.replicaset.available

Total number of available pods (ready for at least minReadySeconds) targeted by this replicaset

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.replicaset.desired

Number of desired pods in this replicaset

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.replication_controller.available

Total number of available pods (ready for at least minReadySeconds) targeted by this replication_controller

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.replication_controller.desired

Number of desired pods in this replication_controller

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.resource_quota.hard_limit

The upper limit for a particular resource in a specific namespace. Will only be sent if a quota is specified. CPU requests/limits will be sent as millicores

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {resource} | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| resource | the name of the resource on which the quota is applied | Any Str |

### k8s.resource_quota.used

The usage for a particular resource in a specific namespace. Will only be sent if a quota is specified. CPU requests/limits will be sent as millicores

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {resource} | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| resource | the name of the resource on which the quota is applied | Any Str |

### k8s.statefulset.current_pods

The number of pods created by the StatefulSet controller from the StatefulSet version

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.statefulset.desired_pods

Number of desired pods in the stateful set (the `spec.replicas` field)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.statefulset.ready_pods

Number of pods created by the stateful set that have the `Ready` condition

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### k8s.statefulset.updated_pods

Number of pods created by the StatefulSet controller from the StatefulSet version

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {pod} | Gauge | Int |

### openshift.appliedclusterquota.limit

The upper limit for a particular resource in a specific namespace.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {resource} | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| k8s.namespace.name | The k8s namespace name. | Any Str |
| resource | the name of the resource on which the quota is applied | Any Str |

### openshift.appliedclusterquota.used

The usage for a particular resource in a specific namespace.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {resource} | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| k8s.namespace.name | The k8s namespace name. | Any Str |
| resource | the name of the resource on which the quota is applied | Any Str |

### openshift.clusterquota.limit

The configured upper limit for a particular resource.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {resource} | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| resource | the name of the resource on which the quota is applied | Any Str |

### openshift.clusterquota.used

The usage for a particular resource with a configured limit.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {resource} | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| resource | the name of the resource on which the quota is applied | Any Str |

## Optional Metrics

The following metrics are not emitted by default. Each of them can be enabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: true
```

### k8s.pod.status_reason

Current status reason of the pod (1 - Evicted, 2 - NodeAffinity, 3 - NodeLost, 4 - Shutdown, 5 - UnexpectedAdmissionError, 6 - Unknown)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
|  | Gauge | Int |

## Resource Attributes

| Name | Description | Values | Enabled |
| ---- | ----------- | ------ | ------- |
| container.id | The container id. | Any Str | true |
| container.image.name | The container image name | Any Str | true |
| container.image.tag | The container image tag | Any Str | true |
| k8s.container.name | The k8s container name | Any Str | true |
| k8s.cronjob.name | The k8s CronJob name | Any Str | true |
| k8s.cronjob.uid | The k8s CronJob uid. | Any Str | true |
| k8s.daemonset.name | The k8s daemonset name. | Any Str | true |
| k8s.daemonset.uid | The k8s daemonset uid. | Any Str | true |
| k8s.deployment.name | The name of the Deployment. | Any Str | true |
| k8s.deployment.uid | The UID of the Deployment. | Any Str | true |
| k8s.hpa.name | The k8s hpa name. | Any Str | true |
| k8s.hpa.uid | The k8s hpa uid. | Any Str | true |
| k8s.job.name | The k8s pod name. | Any Str | true |
| k8s.job.uid | The k8s job uid. | Any Str | true |
| k8s.kubelet.version | The version of Kubelet running on the node. | Any Str | false |
| k8s.kubeproxy.version | The version of Kube Proxy running on the node. | Any Str | false |
| k8s.namespace.name | The k8s namespace name. | Any Str | true |
| k8s.namespace.uid | The k8s namespace uid. | Any Str | true |
| k8s.node.name | The k8s node name. | Any Str | true |
| k8s.node.uid | The k8s node uid. | Any Str | true |
| k8s.pod.name | The k8s pod name. | Any Str | true |
| k8s.pod.qos_class | The k8s pod qos class name. One of Guaranteed, Burstable, BestEffort. | Any Str | false |
| k8s.pod.uid | The k8s pod uid. | Any Str | true |
| k8s.replicaset.name | The k8s replicaset name | Any Str | true |
| k8s.replicaset.uid | The k8s replicaset uid | Any Str | true |
| k8s.replicationcontroller.name | The k8s replicationcontroller name. | Any Str | true |
| k8s.replicationcontroller.uid | The k8s replicationcontroller uid. | Any Str | true |
| k8s.resourcequota.name | The k8s resourcequota name. | Any Str | true |
| k8s.resourcequota.uid | The k8s resourcequota uid. | Any Str | true |
| k8s.statefulset.name | The k8s statefulset name. | Any Str | true |
| k8s.statefulset.uid | The k8s statefulset uid. | Any Str | true |
| openshift.clusterquota.name | The k8s ClusterResourceQuota name. | Any Str | true |
| openshift.clusterquota.uid | The k8s ClusterResourceQuota uid. | Any Str | true |