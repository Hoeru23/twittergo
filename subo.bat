git add .
git commit -m "Cambio a bootstrap"
git push
go build -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip main.exe