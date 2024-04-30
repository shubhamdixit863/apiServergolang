
.PHONY: install_gomock
install_gomock:
	   go install github.com/golang/mock/mockgen@v1.6.0


 .PHONY: mock-service
mock-service:
			 @echo "generating mock service"
			 mockgen -source=internal/services/auth.go > internal/services/mock/mockauth.go

.PHONY: docker-build
docker-build:
		 docker build  --no-cache -t apiservergolangtest .
