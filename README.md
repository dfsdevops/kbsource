# KBsources

Kubernetes Base Sources. Bring consistency to your config repos!


### The Problem

kustomize-based configuration repositories for Kubernetes benefit from maintaining a declarative set of manifests as a "base" resource within a centralized git repository. When bringing new software into such a git repository, one must decipher the upstream repositories installation methods and painstakingly translate it into a local format. Often times there isn't in a kustomize-friendly format, or it's a helm chart that needs to be templated first and then teased apart into separate files.

Then, months down the line, all of these upstream projects get updates, and everything must be updated and compared against the base in your config repo. Worse, sometimes you miss something becuase there isn't.

Furthermore, sometimes there are specific modifications you make to namespace, RBAC, labels to the base, or the order ot fields, or the indentation of lists in yaml, or the. And it's hard to maintain these differences across updates.


### The solution

For the most part, all of these problems could be addressed through some type of automation. Objects in kubernetes follow a predictable structure so therefore if you can "bake" a set of resources from an upstream, you can design a pipeline that reads those resourcs, massages them into the files you want in the format you want. There currently isn't a lot of existing tooling in this area and people still do a lot by hand.

`kbsource` takes a set of kubernetes resources from stdin, and writes each resource into individual files following a naming scheme of your choice.

This allows the possibility of a set of CI pipelines that can read from upstream helm chart updates and manifest release packages, template them, separate them into files, and overwrite the existing files in your repo in a consistent, git-diffable way. One can then utilize other tools such as the wonderful https://github.com/snarlysodboxer/predictable-yaml to autoformat further if desired. Throw in a kustomize stage that adds patches or transformations! Run it by hand, have it open an automatic PR or just have it merge directly! the possibilities are endless.
### Further improvements past initial version

mutate labels/namespace
autoformatting/auto-ordering


### Roadmap

- Add initial README
- Setup intial cobra-CLI structure/verbs
- Take stdin stream of k8s resources (yaml-only for now), read them into an internal state.
- Marshall output into separate yaml files named after each resource name and `Kind`
- Add config format/flag to represent (perhaps something like `--format 'KIND.name.apiVersion.yaml'` at it's most explicit.).
- create example commands utiltizing upstream helm/kustomize projects.
- allow for grouping of one type of resource (all configmaps, all deployments, etc.) in a single file.
- (maybe) Allow for grouping by label (e.g. component) into a single file.
