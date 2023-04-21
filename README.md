# DevOps MAI
Example repo to demonstrate basic DevOps workflow features.

# Workflow features
- Auto-incremented semver tag
- Auto-published releases
- Checked master builds
- Checked master tests

## Tag increment
Incremented on each PR merged into master if provided with specific label:
- `release:major`
- `release:minor`
- `release:patch`

Otherwise, `norelease` strategy is applied.

## Releases
Based on tags, catches specific PRs and displays them with unapplied changes
under `norelease` tag.

## Building & testing
- Go
- Python

- Linters
- Formatters

