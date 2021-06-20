# Maintainer
This action prints a comment on your PR

## Inputs
- **token**: The personal access token is required

## Examples

```yaml
- name: Running maintain action    
  uses: covalentteam/mantainer@v1
  with:
    number: ${{ github.event.pull_request_review_comment.pull_request.id }}
    owner: covalentteam
    repo: maintainer
    token: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
```
