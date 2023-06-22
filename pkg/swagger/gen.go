package swagger

//go:generate rm -rf server
//go:generate mkdir -p server
//go:generate swagger generate server --target server --name hello-api --spec swagger.yml --exclude-main
