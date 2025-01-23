package operator

const (
	nodeDaemonServiceAccountName  = "ksan-storage-nodedaemon"
	nodeDaemonContainerName       = "nodedaemon"
	nodeDaemonMemRequest          = "45Mi"
	nodeDaemonCPURequest          = "5m"
	priorityClassNameUserCritical = "openshift-user-critical"
)

// node daemon labels
const (
	appKubernetesPartOfLabel    = "app.kubernetes.io/part-of"
	appKubernetesNameLabel      = "app.kubernetes.io/name"
	appKubernetesManagedByLabel = "app.kubernetes.io/managed-by"
	appKubernetesComponentLabel = "app.kubernetes.io/component"

	nodeDaemonNameLabelVal      = "nodedaemon"
	nodeDaemonManagedByLabelVal = "ksan-storage-operator"
	nodeDaemonPartOfLabelVal    = "ksan-storage"
)
