REPO			:= ghcr.io/datachainlab/iroha
TAG_VCPKG		:= 1.3.0-vcpkg
TAG_CUSTOM_IROHAD	:= 1.3.0-custom-irohad

.PHONY: docker-image-vcpkg
docker-image-vcpkg:
	docker build -t $(REPO):$(TAG_VCPKG) -f Dockerfiles/Dockerfile.vcpkg .

.PHONY: docker-image-custom-irohad
docker-image-custom-irohad:
	docker build -t $(REPO):$(TAG_CUSTOM_IROHAD) -f Dockerfiles/Dockerfile.custom-irohad .
