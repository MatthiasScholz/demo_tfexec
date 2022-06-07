git_url := github.com/MatthiasScholz/demo_tfexec
init:
	go mod init $(git_url)

app := demo_tfexec
run:
	go run main.go

deps:
	go mod tidy

version:
	make -version
	go version

clean:
	go clean -modcache

user := matthias.scholz@thoughtworks.com
profile := tw-beach-push
login:
	saml2aws login --idp-account=$(profile) --region eu-central-1 --username=$(user)
