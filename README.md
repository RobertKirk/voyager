# Voyager

[![Go Report Card](https://goreportcard.com/badge/voyagermesh.dev/voyager)](https://goreportcard.com/report/voyagermesh.dev/voyager)
[![Build Status](https://github.com/voyagermesh/voyager/workflows/CI/badge.svg)](https://github.com/voyagermesh/voyager/actions?workflow=CI)
[![codecov](https://codecov.io/gh/voyagermesh/voyager/branch/master/graph/badge.svg)](https://codecov.io/gh/voyagermesh/voyager)
[![Docker Pulls](https://img.shields.io/docker/pulls/appscode/voyager.svg)](https://hub.docker.com/r/appscode/voyager/)
[![Slack](https://slack.appscode.com/badge.svg)](https://slack.appscode.com)
[![Twitter](https://img.shields.io/twitter/follow/voyagermesh.svg?style=social&logo=twitter&label=Follow)](https://twitter.com/intent/follow?screen_name=voyagermesh)

> Secure HAProxy Ingress Controller for Kubernetes

Voyager is a [HAProxy](http://www.haproxy.org/) backed [secure](#certificate) L7 and L4 [ingress](#ingress) controller for Kubernetes developed by [AppsCode](https://appscode.com). This can be used with any Kubernetes cloud providers including aws, gce, gke, azure, acs. This can also be used with bare metal Kubernetes clusters.

## Ingress
Voyager provides L7 and L4 loadbalancing using a custom Kubernetes [Ingress](https://voyagermesh.com/latest/guides/ingress/) resource. This is built on top of the [HAProxy](http://www.haproxy.org/) to support high availability, sticky sessions, name and path-based virtual hosting.
This also support configurable application ports with all the options available in a standard Kubernetes [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/).

## Certificate
Voyager can automatically provision and refresh SSL certificates (including wildcard certificates) issued from Let's Encrypt using a custom Kubernetes [Certificate](https://voyagermesh.com/latest/guides/certificate/) resource.

## Supported Versions
Please pick a version of Voyager that matches your Kubernetes installation.

| Voyager Version                                                                             | Docs                                                                 | Kubernetes Version | Prometheus operator Version |
|---------------------------------------------------------------------------------------------|----------------------------------------------------------------------|--------------------|-----------------------------|
| [v12.0.0-rc.2](https://github.com/voyagermesh/voyager/releases/tag/v12.0.0-rc.2) (uses CRD) | [User Guide](https://voyagermesh.com/docs/v12.0.0-rc.2/)             | 1.11.x+            | 0.34.0+                     |
| [v11.0.1](https://github.com/voyagermesh/voyager/releases/tag/v11.0.1) (uses CRD)           | [User Guide](https://voyagermesh.com/docs/v11.0.1/)                  | 1.9.x+             | 0.30.0+                     |
| [10.0.0](https://github.com/voyagermesh/voyager/releases/tag/10.0.0) (uses CRD)             | [User Guide](https://voyagermesh.com/docs/10.0.0/)                   | 1.9.x+             | 0.16.0+                     |
| [7.4.0](https://github.com/voyagermesh/voyager/releases/tag/7.4.0) (uses CRD)               | [User Guide](https://voyagermesh.com/docs/7.4.0/)                    | 1.8.x - 1.11.x     | 0.16.0+                     |
| [5.0.0](https://github.com/voyagermesh/voyager/releases/tag/5.0.0) (uses CRD)               | [User Guide](https://voyagermesh.com/docs/5.0.0/)                    | 1.7.x              | 0.12.0+                     |
| [3.2.2](https://github.com/voyagermesh/voyager/releases/tag/3.2.2) (uses TPR)               | [User Guide](https://github.com/voyagermesh/voyager/tree/3.2.2/docs) | 1.5.x - 1.7.x      | < 0.12.0                    |

## Installation
To install Voyager, please follow the guide [here](https://voyagermesh.com/latest/setup/install/).

## Using Voyager
Want to learn how to use Voyager? Please start [here](https://voyagermesh.com/latest/welcome/).

## Voyager API Clients
You can use Voyager api clients to programmatically access its CRD objects. Here are the supported clients:

- Go: [https://github.com/voyagermesh/voyager](/client/clientset/versioned)
- Java: https://github.com/voyagermesh/java

## Contribution guidelines
Want to help improve Voyager? Please start [here](https://voyagermesh.com/latest/welcome/contributing/).

---

**Voyager binaries collects anonymous usage statistics to help us learn how the software is being used and how we can improve it.
To disable stats collection, run the operator with the flag** `--enable-analytics=false`.

---

## Acknowledgement
 - docker-library/haproxy https://github.com/docker-library/haproxy
 - kubernetes/contrib https://github.com/kubernetes/contrib/tree/master/service-loadbalancer
 - kubernetes/ingress https://github.com/kubernetes/ingress
 - xenolf/lego https://github.com/appscode/lego
 - kelseyhightower/kube-cert-manager https://github.com/kelseyhightower/kube-cert-manager
 - PalmStoneGames/kube-cert-manager https://github.com/PalmStoneGames/kube-cert-manager
 - [Kubernetes cloudprovider implementation](https://github.com/kubernetes/kubernetes/tree/master/pkg/cloudprovider)
 - openshift/generic-admission-server https://github.com/openshift/generic-admission-server
 - TimWolla/haproxy-auth-request https://github.com/TimWolla/haproxy-auth-request

## Support

We use Slack for public discussions. To chit chat with us or the rest of the community, join us in the [AppsCode Slack team](https://appscode.slack.com/messages/C0XQFLGRM/details/) channel `#general`. To sign up, use our [Slack inviter](https://slack.appscode.com/).

If you have found a bug with Voyager or want to request for new features, please [file an issue](https://github.com/voyagermesh/voyager/issues/new).
