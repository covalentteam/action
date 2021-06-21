# Maintainer
This action prints a comment on your PR

## Inputs
- **token**: The personal access token is required

## Example usage

```yaml
- name: Running maintain action    
  uses: covalentteam/mantainer@v1
  with:
    event: ${{ toJson(github.event.pull_request_review_comment) }}
    token: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
```
