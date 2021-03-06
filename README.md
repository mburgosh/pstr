# PatientSky Template Renderer ( PSTR )
This project is to render templates used for Kubernetes deployment

## --all <bool>
Renders Deployment, Service, Autoscaler & Ingress from Json

## --limit <string>
Limits to specific service name, can be used in conjunction with --deploy, --service, --autoscaler & --ingress to narrow output

## --deploy <bool>
Renders only deployment

## --service <bool>
Renders only service

## --genericservice <bool>
Renders only generic service

## --autoscaler
Renders only HPA ( AutoScaler )

## --ingress
Renders only Ingress rules

## --namespace <string>
Set namespace to use in template rendering

## --build <string>
Set build. This propagates the value of "Deploy_build"

## --hostname <string>
Comma delimited list of hostnames to use for template rendering.

## --output <string>
Path to where to write output YAML files.

If not specified files will be written in the current folder


**WARNING, If multiple services has Service.#.Ports.External.HTTP set this will generate multiple ingress rules with the same hostname if --limit is not used together with**


## notepad
bamboo_deploy_release="34" bamboo_buildNumber="43d24" CONSUL_APPLICATION="consul_app" cluster_ip="127.0.0.1" CONSUL_ENVIRONMENT="consul_env" CONSUL_PASSWORD="consul_pass"  CONSUL_URL="http://consul" CONSUL_USERNAME="consul_user" git_repo="http://git.repo" ssh_key="rsa1234" NEW_RELIC_LICENSE_KEY="9876er54321" go run *.go  --build_id dfb1337 --namespace=hptest --all --hostname=test.domain.com,test2.another.com --file ./serviceDefinition.json --output ./out --limit app


bamboo_deploy_release="34" bamboo_buildNumber="43d24" CONSUL_APPLICATION="consul_app" cluster_ip="127.0.0.1" CONSUL_ENVIRONMENT="consul_env" CONSUL_PASSWORD="consul_pass"  CONSUL_URL="http://consul" CONSUL_USERNAME="consul_user" git_repo="http://git.repo" ssh_key="rsa1234" NEW_RELIC_LICENSE_KEY="9876er54321" ./pstr --build_id dfb1337 --namespace=hptest --all --hostname=test.domain.com,test2.another.com --file ./serviceDefinition.json --output ./out --limit app
