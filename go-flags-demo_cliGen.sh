 pwd=`pwd`
 cd $GOPATH/src/github.com/go-easygen/easygen/test
 easygen commandlineGoFlags.header,commandlineGoFlags.ityped.tmpl,commandlineGoFlags "$pwd/go-flags-demo"_cli | gofmt > "$pwd/go-flags-demo"_cliDef.go
