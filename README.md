# Ggg

Go [Grafana](https://grafana.com) Generator. Describe Grafana dashboards using Go DSL.

## What is it?

This is a Go-based DSL that'll help to generate Grafana dashboards. It is designed so that dashboards
should be written code-first and not by trying to translate Grafana-generated JSON. You don't need to
be familiar with Grafana dashboard JSON schema but some familiarity with Grafana itself is a must.

Ggg is primarily intended to be used in Terraform with Grafana provider but isn't limited to it.

## Why on earth we need yet another Grafana dashboard generator?

There are several alternatives in the wild:
* [Jsonnet](https://jsonnet.org/) + [Grafonnet](https://github.com/grafana/grafonnet-lib)
* Python-based [grafanalib](https://github.com/weaveworks/grafanalib)

Our team uses Grafana very heavily. We have thousands of dashboards and some of them are quite complex.
To tackle this complexity we generate dashboards from code.

The main programming language in our team is Go, and we have a deep expertise in it. We have to use
familiar tools to code dashboard generation logic for it to be fully comprehendable by all team members.
Both Jsonnet and Python aren't such a tools for us. Furthermore, we already have a lot of convenient and
ready to use tools for the Go stack.

It would be of real help if Grafana developers maintained a complete and precise formal schema of their
dashboards. But this isn't the case, unfortunately. While there is some effort in this direction, the
problem still hasn't been solved.

Jsonnet being a somewhat exotic language is missing a solid IDE support. This prevents people not familiar
with it to effectively use it in rare occasions when they need a new dashboard. And it also isn't easy to
convince people that they need to learn yet another programming language.

Grafanalib uses Python which is more familiar to us than Jsonnet. But the main drawback is that it still
somewhat low level and this forces us to have a higher level layer. Maintaining our own Python packages
with higher level layer requires setting up proper infra which we don't have at the moment. And we also
don't have enough resources and expertise to set it up and maintain.

## Getting started

Initialize a Go modules project:

```shell
$ go mod init <some-name>
$ go get github.com/g7r/ggg
```

Add a `main.go` file:

```golang
package main

import (
	"github.com/g7r/ggg"
	"github.com/g7r/ggg/tftool"
)

func main() {
	dashboard := ggg.NewDashboard(func(builder *ggg.DashboardBuilder) {
        // ....
	})
	
	tftool.Main(dashboard)
}
```

You can see the dashboard JSON by invoking your `main.go` directly:

```shell
$ go run ./main.go 
```

Describe Grafana dashboard with Terraform:

```terraform
terraform {
  required_providers {
    grafana = {
      source = "grafana/grafana"
      version = "1.18.0"
    }
  }
}

provider "grafana" {
  url  = "<your-grafana-url>"
  auth = "<your-grafana-auth-token>"
}

data "external" "my_dashboard" {
  program = ["go", "run", "./main.go"]
}

resource "grafana_dashboard" "my_dashboard" {
  config_json = data.external.my_dashboard.result.dashboard
}
```