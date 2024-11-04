git add .
git commit -m "Cambio a bootstrap"
git push
GOARCH=amd64 GOOS=linux go build -o bootstrap main.go
del main.zip
zip main.zip bootstrap