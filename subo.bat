git add .
git commit -m "Consulta Relacion - 84"
git push
SET GOOS=linux
SET GOARCH=amd64
go build -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap