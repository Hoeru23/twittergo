git add .
git commit -m "Leer Tweet - 68"
git push
SET GOOS=linux
SET GOARCH=amd64
go build -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap