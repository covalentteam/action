# Maintainer
This action prints a comment on your PR

## Inputs
- **token**: The personal access token is required

## Example usage

```yamlcle
- name: Running maintain action    
  uses: covalentteam/mantainer@v1
  with:
    token: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
```
