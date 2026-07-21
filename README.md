# se-tooling-ci-common
Shared CI and Testing config for SE Tooling

## PR Title Check

`common/github_workflows/pr-title-check.yaml` is a reusable workflow that
verifies every pull request title begins with a JIRA ticket reference, e.g.
`[JIRA-1234] title of PR`.

To use it, add a small caller workflow to the target repository that invokes
this reusable workflow on pull request events. It fails if the title does not
start with a bracketed uppercase project key and issue number
(regex `^\[[A-Z][A-Z0-9]+-[0-9]+\].+`).

Create `.github/workflows/pr-title-check.yaml` in the target repository with
the following content:

```yaml
name: PR Title Check
on:
  pull_request:
    types: [opened, edited, reopened, synchronize]

jobs:
  check-title:
    uses: canonical/se-tooling-ci-common/common/github_workflows/pr-title-check.yaml@main
```
