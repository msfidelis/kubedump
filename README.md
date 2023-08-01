# Kubedump - Simple tool to dump and restore all kubernetes resources across namespaces to simplify backups and migrations

<p>
  <a href="README.md" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="/" target="_blank">
    <img alt="Build CI" src="https://github.com/msfidelis/kubedump/workflows/kubedump%20ci/badge.svg" />
  </a>  
  <a href="/" target="_blank">
    <img alt="Release" src="https://github.com/msfidelis/kubedump/workflows/release%20packages/badge.svg" />
  </a>
  <a href="https://twitter.com/fidelissauro" target="_blank">
    <img alt="Twitter: fidelissauro" src="https://img.shields.io/twitter/follow/fidelissauro.svg?style=social" />
  </a>  
</p>

# Introduction

> Kubedump is a simple tool to make easy backups and workloads migrations between clusters.


# Installation 

### Docker 

```bash
docker pull fidelissauro/kubedump:latest
```

```bash
docker run --network -v ~/.kubeconfig:/home/root/.kubeconfig host -it fidelissauro/kubedump:latest dump chip
```


### MacOS amd64

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1/kubedump_0.1_darwin_arm64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```


### MacOS arm64

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1/kubedump_0.1_darwin_amd64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

### Linux amd64 

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1/kubedump_0.1_linux_amd64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

### Linux arm64 

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1/kubedump_0.1_linux_arm64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

### Freebsd amd64 

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1/kubedump_0.1_freebsd_amd64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

### Freebsd arm64 

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1/kubedump_0.1_freebsd_arm64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

# v0 Usage 

```bash
Usage:
  kubedump [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dump        dump all resources of a Kubernetes namespace
  help        Help about any command
  restore     restore all resources dumped using kubedump of a Kubernetes namespace

Flags:
  -h, --help   help for kubedump

Use "kubedump [command] --help" for more information about a command.
``````

## v0 Usage - Dump

```bash
kubedump dump --help
dump all resources of a Kubernetes namespace

Usage:
  kubedump dump [namespace] [flags]

Flags:
      --config-file string        kubedump config file location
  -d, --dry-run                   Perform a dry-run backup (no actual backup will be performed)
      --format string             Dump output format (default "yaml")
  -h, --help                      help for dump
      --kubectl-location string   Custom kubectl binary or alias (default "/usr/local/bin/kubectl")
      --project string            Project name (default "kubedump")
      --resources string          Kubernetes resources separated by comma (default "deployment,service,hpa,ingress,serviceaccount,daemonset,statefulset,job,cronjob")
```

## Dump namespace resources 

```bash
❯ kubedump dump chip --project cluster_x
Dumping 'deployment' of namespace 'chip'
Dumping 'service' of namespace 'chip'
Dumping 'hpa' of namespace 'chip'
Dumping 'ingress' of namespace 'chip'
Dumping 'serviceaccount' of namespace 'chip'
Dumping 'daemonset' of namespace 'chip'
Dumping 'statefulset' of namespace 'chip'
Dumping 'job' of namespace 'chip'
Dumping 'cronjob' of namespace 'chip'
```

An folder will be created with `project` name, with output structure like 

```
❯ tree cluster_x
cluster_x
└── chip
    ├── 00-namespace.yaml
    ├── cronjob.yaml
    ├── daemonset.yaml
    ├── deployment.yaml
    ├── hpa.yaml
    ├── ingress.yaml
    ├── job.yaml
    ├── service.yaml
    ├── serviceaccount.yaml
    └── statefulset.yaml
```

## Dump custom resources 

```bash
kubedump dump chip --project cluster_x --resources deployment,service,hpa,ingress,serviceaccount,daemonset,statefulset,job,cronjob,virtualservice,gateway,destinationrules
Dumping 'deployment' of namespace 'chip'
Dumping 'service' of namespace 'chip'
Dumping 'hpa' of namespace 'chip'
Dumping 'ingress' of namespace 'chip'
Dumping 'serviceaccount' of namespace 'chip'
Dumping 'daemonset' of namespace 'chip'
Dumping 'statefulset' of namespace 'chip'
Dumping 'job' of namespace 'chip'
Dumping 'cronjob' of namespace 'chip'
Dumping 'virtualservice' of namespace 'chip'
Dumping 'gateway' of namespace 'chip'
Dumping 'destinationrules' of namespace 'chip'
```

## Dump format option 

```
kubedump dump chip --project cluster_x --format yaml
kubedump dump chip --project cluster_y --format json
```

```
❯ tree cluster_x
cluster_x
└── chip
    ├── 00-namespace.yaml
    ├── cronjob.yaml
    ├── daemonset.yaml
    ├── deployment.yaml
    ├── hpa.yaml
    ├── ingress.yaml
    ├── job.yaml
    ├── service.yaml
    ├── serviceaccount.yaml
    └── statefulset.yaml

1 directory, 10 files
❯ tree cluster_y
cluster_y
└── chip
    ├── 00-namespace.json
    ├── cronjob.json
    ├── daemonset.json
    ├── deployment.json
    ├── hpa.json
    ├── ingress.json
    ├── job.json
    ├── service.json
    ├── serviceaccount.json
    └── statefulset.json

1 directory, 10 files
```


## v0 Usage - Restore 

```bash
❯ kubedump restore --help
restore all resources dumped using kubedump of a Kubernetes namespace

Usage:
  kubedump restore [namespace] [flags]

Flags:
      --config-file string        kubedump config file location
  -h, --help                      help for restore
      --kubectl-location string   Custom kubectl binary or alias (default "/usr/local/bin/kubectl")
      --project string            Project name (default "kubedump")
```


```bash
kubedump restore chip 

Restoring namespace 'chip'
Restoring kubedump/chip/00-namespace.yaml on namespace 'chip'
Restoring kubedump/chip/cronjob.yaml on namespace 'chip'
Restoring kubedump/chip/daemonset.yaml on namespace 'chip'
Restoring kubedump/chip/deployment.yaml on namespace 'chip'
Restoring kubedump/chip/hpa.yaml on namespace 'chip'
Restoring kubedump/chip/ingress.yaml on namespace 'chip'
Restoring kubedump/chip/job.yaml on namespace 'chip'
Restoring kubedump/chip/service.yaml on namespace 'chip'
Restoring kubedump/chip/serviceaccount.yaml on namespace 'chip'
Restoring kubedump/chip/statefulset.yaml on namespace 'chip'
```


