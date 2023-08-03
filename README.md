# Kubedump - Simple tool to dump and restore kubernetes resources

![logo](.github/img/logo.png)

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

## Docker 

```bash
docker pull fidelissauro/kubedump:latest
```

```bash
docker run --network -v ~/.kubeconfig:/home/root/.kubeconfig host -it fidelissauro/kubedump:latest dump chip
```


## MacOS amd64

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1.1/kubedump_0.1.1_darwin_arm64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```


## MacOS arm64

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1.1/kubedump_0.1.1_darwin_amd64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

## Linux amd64 

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1.1/kubedump_0.1.1_linux_amd64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

## Linux arm64 

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1.1/kubedump_0.1.1_linux_arm64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

## Freebsd amd64 

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1.1/kubedump_0.1.1_freebsd_amd64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

## Freebsd arm64 

```bash
wget https://github.com/msfidelis/kubedump/releases/download/v0.1.1/kubedump_0.1.1_freebsd_arm64 -O kubedump 
mv kubedump /usr/local/bin 
chmod +x /usr/local/bin/kubedump
```

#  Usage

```bash
kubedump --help

Usage:
  kubedump [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  dump         dump all resources of a Kubernetes namespace
  dump-file    dump all resources of with file custom configs
  help         Help about any command
  restore      restore all resources dumped using kubedump of a Kubernetes namespace
  restore-file restore all resources of with custom configs from configuration files

Flags:
  -h, --help   help for kubedump

Use "kubedump [command] --help" for more information about a command.
``````

## Usage - Dump

```bash
kubedump dump --help

dump all resources of a Kubernetes namespace

Usage:
  kubedump dump [namespace] [flags]

Flags:
  -d, --dry-run                   Perform a dry-run backup (no actual backup will be performed)
      --format string             Dump output format (default "yaml")
  -h, --help                      help for dump
      --kubectl-location string   Custom kubectl binary or alias (default "kubectl")
      --project string            Project name (default "kubedump")
      --resources string          Kubernetes resources separated by comma (default "deployment,service,hpa,ingress,serviceaccount,daemonset,statefulset,job,cronjob")
```

### Dump namespace resources 

```bash
❯ kubedump dump chip --project cluster_x

2023/08/02 21:48:08 INFO Starting dump namespace=chip
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=namespace
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=deployment
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=service
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=hpa
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=ingress
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=serviceaccount
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=daemonset
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=statefulset
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=job
2023/08/02 21:48:08 INFO Dumping resources namespace=chip resource=cronjob
2023/08/02 21:48:08 INFO Success namespace=chip output_files=./cluster_x/chip
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

### Dump custom resources 

```bash
kubedump dump chip --project cluster_x --resources deployment,service,hpa,ingress,serviceaccount,daemonset,statefulset,job,cronjob,virtualservice,gateway,destinationrules

2023/08/02 21:49:19 INFO Starting dump namespace=chip
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=namespace
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=deployment
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=service
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=hpa
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=ingress
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=serviceaccount
2023/08/02 21:49:19 WARN No resource found in namespace namespace=chip resource=serviceaccount
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=daemonset
2023/08/02 21:49:19 WARN No resource found in namespace namespace=chip resource=daemonset
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=statefulset
2023/08/02 21:49:19 WARN No resource found in namespace namespace=chip resource=statefulset
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=job
2023/08/02 21:49:19 WARN No resource found in namespace namespace=chip resource=job
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=cronjob
2023/08/02 21:49:19 WARN No resource found in namespace namespace=chip resource=cronjob
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=virtualservice
2023/08/02 21:49:19 ERRO Error to Dump resource namespace=chip resource=virtualservice file=error exit status 1="missing value"
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=gateway
2023/08/02 21:49:19 ERRO Error to Dump resource namespace=chip resource=gateway file=error exit status 1="missing value"
2023/08/02 21:49:19 INFO Dumping resources namespace=chip resource=destinationrules
2023/08/02 21:49:20 ERRO Error to Dump resource namespace=chip resource=destinationrules file=error exit status 1="missing value"
2023/08/02 21:49:20 INFO Success namespace=chip output_files=./cluster_x/chip
```

### Dump format option 

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


## Usage - Restore 

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

2023/08/02 21:50:13 INFO Restoring resources namespace=chip resource=namespace
2023/08/02 21:50:13 INFO Resources restored: namespace=chip resources=kubedump/chip/00-namespace.yaml
2023/08/02 21:50:13 INFO Resources restored: namespace=chip resources=kubedump/chip/deployment.yaml
2023/08/02 21:50:13 INFO Resources restored: namespace=chip resources=kubedump/chip/hpa.yaml
2023/08/02 21:50:13 INFO Resources restored: namespace=chip resources=kubedump/chip/ingress.yaml
2023/08/02 21:50:13 INFO Resources restored: namespace=chip resources=kubedump/chip/service.yaml
```


## Using Config Files 

You can use `dump-file` and `restore-file` command with yaml file configuration to automate dumps between a lot of namespaces at same time. 

You can create a `kubedump-file.yaml` example like this: 

```yaml
project: cluster_x
format: yaml
namespaces: 
  - chip
  - whois
resources:
  - deployment
  - service
  - hpa
  - ingress
  - serviceaccount
  - daemonset
  - statefulset
  - jobs
  - cronjob
  - secret
  - configmap

```

```bash
kubedump dump-file --config-file kubedump-file.yaml

2023/08/02 21:52:13 INFO Dumping resources namespace=chip resource=namespace
2023/08/02 21:52:13 INFO Dumping resources namespace=chip resource=deployment
2023/08/02 21:52:14 INFO Dumping resources namespace=chip resource=service
2023/08/02 21:52:14 INFO Dumping resources namespace=chip resource=hpa
2023/08/02 21:52:14 INFO Dumping resources namespace=chip resource=ingress
2023/08/02 21:52:14 INFO Dumping resources namespace=chip resource=serviceaccount
2023/08/02 21:52:14 WARN No resource found in namespace namespace=chip resource=serviceaccount
2023/08/02 21:52:14 INFO Success namespace=chip output_files=./cluster_x/chip
2023/08/02 21:52:14 INFO Dumping resources namespace=whois resource=namespace
2023/08/02 21:52:14 INFO Dumping resources namespace=whois resource=deployment
2023/08/02 21:52:14 INFO Dumping resources namespace=whois resource=service
2023/08/02 21:52:14 INFO Dumping resources namespace=whois resource=hpa
2023/08/02 21:52:14 INFO Dumping resources namespace=whois resource=ingress
2023/08/02 21:52:14 INFO Dumping resources namespace=whois resource=serviceaccount
2023/08/02 21:52:14 WARN No resource found in namespace namespace=whois resource=serviceaccount
2023/08/02 21:52:14 INFO Success namespace=whois output_files=./cluster_x/whois
```


```bash
kubedump restore-file --config-file kubedump-file.yaml

2023/08/02 21:52:45 INFO Restoring resources namespace=chip resource=namespace
2023/08/02 21:52:45 INFO Resources restored: namespace=chip resources=cluster_x/chip/00-namespace.yaml
2023/08/02 21:52:45 INFO Resources restored: namespace=chip resources=cluster_x/chip/deployment.yaml
2023/08/02 21:52:45 INFO Resources restored: namespace=chip resources=cluster_x/chip/hpa.yaml
2023/08/02 21:52:46 INFO Resources restored: namespace=chip resources=cluster_x/chip/ingress.yaml
2023/08/02 21:52:46 INFO Resources restored: namespace=chip resources=cluster_x/chip/service.yaml
2023/08/02 21:52:46 INFO Restoring resources namespace=whois resource=namespace
2023/08/02 21:52:46 INFO Resources restored: namespace=whois resources=cluster_x/whois/00-namespace.yaml
2023/08/02 21:52:46 INFO Resources restored: namespace=whois resources=cluster_x/whois/deployment.yaml
2023/08/02 21:52:46 INFO Resources restored: namespace=whois resources=cluster_x/whois/hpa.yaml
2023/08/02 21:52:46 INFO Resources restored: namespace=whois resources=cluster_x/whois/ingress.yaml
2023/08/02 21:52:46 INFO Resources restored: namespace=whois resources=cluster_x/whois/service.yaml
```