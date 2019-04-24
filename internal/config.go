package bunker

// ContainerdSocketPath is the full path of the socket file that is used as the
// connection to containerd
const ContainerdSocketPath = "/run/bunker/bunkerd.sock"

// ContainerdNamespace is the namespace label used in the containerd context.
const ContainerdNamespace = "bunker"

// ContainerdDaemonPIDFile is the file where pid of bunkerd daemon is kept.
const ContainerdDaemonPIDFile = "/run/bunkerd.pid"
