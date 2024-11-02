git add .
git commit -m "Primera subida"
git push
go build -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip main.exe