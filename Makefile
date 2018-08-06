TAG ?= echo
REV = $$(git rev-parse --short HEAD)


build-docker:
	docker build -t ${TAG}-server:${REV} . -f Dockerfile.server
	docker build -t ${TAG}-client:${REV} . -f Dockerfile.client

publish:
	docker tag  ${TAG}-server:${REV} msmallinsky/${TAG}-server:${REV}
	docker push msmallinsky/${TAG}-server:${REV}
	docker tag  ${TAG}-client:${REV} msmallinsky/${TAG}-client:${REV}
	docker push msmallinsky/${TAG}-client:${REV}
	
