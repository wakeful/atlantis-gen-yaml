# atlantis-gen-yaml
this tool generates `projects` section for your [atlantis.yaml](https://www.runatlantis.io/docs/repo-level-atlantis-yaml.html#terragrunt) file by parsing your [terragrunt](https://terragrunt.gruntwork.io) file(s) and `dependencies`.


### example case
let's assume we have a `example-acm` dir with this files next to our `atlantis.yaml`

```shell
$ find . -type f
./example-acm/record/terragrunt.hcl
./example-acm/certificate/terragrunt.hcl
./example-acm/validation/terragrunt.hcl
./atlantis.yaml
```

and content

`example-acm/certificate/terragrunt.hcl`
```hcl
terraform {
  source = "../src"
}

inputs = {
  fqdn = "my-new-domain.example.com"
}
```
`example-acm/record/terragrunt.hcl`
```hcl
terraform {
  source = "../src"
}

dependency "cert" {
  config_path = "../certificate"

  mock_outputs = {
    valid_name  = "fooBar"
    valid_type  = "CNAME"
    valid_value = "fooBar"
  }
}

inputs = {
  name  = dependency.cert.outputs.valid_name
  type  = dependency.cert.outputs.valid_type
  value = dependency.cert.outputs.valid_value
}
```

`example-acm/validation/terragrunt.hcl`
```hcl
terraform {
  source = "../src"
}

dependencies {
  paths = ["../record"]
}

dependency "cert" {
  config_path = "../certificate"

  mock_outputs = {
    arn = "fooBar"
  }
}

inputs = {
  arn = dependency.cert.outputs.arn
}
```
running the binary
```shell
$ gen-atlantis-yaml -path example-acm
```
will produce `YAML` file with all `dependencies` from your `terragrunt.hcl` file(s).
```yaml
projects:
- autoplan:
    enabled: true
    when_modified:
      - '*.hcl'
      - '*.tf*'
  dir: example-acm/certificate
  workflow: terragrunt
- autoplan:
    enabled: true
    when_modified:
      - '*.hcl'
      - '*.tf*'
      - ../certificate
  dir: example-acm/record
  workflow: terragrunt
- autoplan:
    enabled: true
    when_modified:
      - '*.hcl'
      - '*.tf*'
      - ../certificate
      - ../record
  dir: example-acm/validation
  workflow: terragrunt
```

## custom atlantis config?
you can include custom **atlantis** variables by adding them to `.atlantis-conf.yaml` file e.q.
```yaml
automerge: true
parallel_apply: false
parallel_plan: false
version: 3
```
would produce an `atlantis.yaml` file with the extra config files.
```yaml
automerge: true
parallel_apply: false
parallel_plan: false
version: 3
projects:
- autoplan:
    enabled: true
    when_modified:
      - '*.hcl'
      - '*.tf*'
  dir: example-acm/certificate
  workflow: terragrunt
- autoplan:
    enabled: true
    when_modified:
      - '*.hcl'
      - '*.tf*'
      - ../certificate
  dir: example-acm/record
  workflow: terragrunt
- autoplan:
    enabled: true
    when_modified:
      - '*.hcl'
      - '*.tf*'
      - ../certificate
      - ../record
  dir: example-acm/validation
  workflow: terragrunt
```

## ToDo
- [ ] add `patch` command to edit existing `atlantis.yaml` file
- [ ] publish pkg via `brew`
