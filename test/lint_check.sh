p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
make -C ../  fmt
golint ../fmg
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortimanager2/auth
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortimanager2/config
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortimanager2/request
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortimanager2/sdkcore
make -C ../  build

