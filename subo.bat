git add .
git commit -m "Cambio a bootstrap"
git push
$$Env:GOARCH=arm64; $$Env:GOOS=linux; go build -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap