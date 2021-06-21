# Maintainer
This action prints a comment on your PR

## Inputs
- **token**: The personal access token is required

## Example usage

```yaml
  - name: Comment on Pull Request
    uses: covalentteam/maintainer@main
    with:
      owner: ${{ github.event.repository.owner.login }}
      pull_request_id: ${{ github.event.repository.id }}
      repository: ${{ github.event.repository.name }}
      token: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
```
