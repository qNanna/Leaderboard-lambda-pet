# Install npm before run commands
# Also read Makefile
## Tested and written on mac m3

### Install
  #### global
    - make npm-install-global 
  #### local
    - make npm-install-local
  #### go dependencies
    - go mod tidy

### Run local
  - make run

### Deploy
  - make deploy

### Remove lambdas
  - make remove