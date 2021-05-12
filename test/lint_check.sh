p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
make -C ../  fmt
golint ../fmg
golint ../vendor/github.com/fortinetdev/forti-sdk-go2/fmg/auth
golint ../vendor/github.com/fortinetdev/forti-sdk-go2/fmg/config
golint ../vendor/github.com/fortinetdev/forti-sdk-go2/fmg/request
golint ../vendor/github.com/fortinetdev/forti-sdk-go2/fmg/sdkcore
make -C ../  build

