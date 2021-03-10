all:
	$(MAKE) install_gloo install_app

install_gloo:
	helm repo add gloo https://storage.googleapis.com/solo-public-helm

	kubectl create ns gloo-system
	helm upgrade -i gloo gloo/gloo \
	--namespace gloo-system

install_flagger:
	helm repo add flagger https://flagger.app

	helm upgrade -i flagger flagger/flagger \
	--namespace gloo-system \
	--set prometheus.install=true \
	--set meshProvider=gloo \
	--set image.repository=docker.io/mayankshah1607/flagger-grofers \
	--set image.tag=latest

uninstall_flagger:
	helm uninstall flagger -n gloo

install_loadtester:
	kubectl -n test apply -k https://github.com/fluxcd/flagger//kustomize/tester?ref=main

install_app:
	kubectl create ns test | true
	kubectl -n test apply -k https://github.com/fluxcd/flagger//kustomize/podinfo?ref=main

uninstall_app:
	kubectl -n test delete -k https://github.com/fluxcd/flagger//kustomize/podinfo?ref=main

