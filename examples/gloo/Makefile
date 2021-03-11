all:
	$(MAKE) install_gloo install_flagger install_app install_loadtester

install_gloo:
	helm repo add gloo https://storage.googleapis.com/solo-public-helm

	kubectl create ns gloo-system | true
	helm upgrade -i gloo gloo/gloo \
	--namespace gloo-system

install_flagger:
	git clone git@github.com:mayankshah1607/flagger.git -b v1.5.0-grofers

	helm upgrade -i flagger ./flagger/charts/flagger \
	--namespace gloo-system \
	--set prometheus.install=true \
	--set meshProvider=gloo \
	--set image.repository=docker.io/mayankshah1607/flagger-grofers \
	--set image.tag=latest \

	rm -rf flagger/

uninstall_flagger:
	helm uninstall flagger -n gloo

install_loadtester:
	kubectl -n test apply -k https://github.com/fluxcd/flagger//kustomize/tester?ref=main

install_app:
	kubectl create ns test | true
	kubectl -n test apply -k https://github.com/fluxcd/flagger//kustomize/podinfo?ref=main

uninstall_app:
	kubectl -n test delete -k https://github.com/fluxcd/flagger//kustomize/podinfo?ref=main

install_canary:
	kubectl apply -f canary.yaml
	kubectl apply -f virtual_service.yaml

uninstall_canary:
	kubectl delete -f canary.yaml
	kubectl delete -f virtual_service.yaml

