Создание версии v1.0.0
go mod init
touch gofunc.go
vi gofunc.go
git add .
git commit -a -m "v1.0.0"
git push
git tag v1.0.0
git push -q origin v1.0.0

Создание версии v1.1.0
git commit -a -m "v1.1.0"
git push
git tag v1.1.0
git push -q origin v1.1.0
