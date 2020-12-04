package template

var (
	Cicd = `
variables:
    AlI_REGISTER_ACCOUNT: ${AliRegisterAccount}
    AlI_REGISTER_PASSWORD: ${AliRegisterPassword}
    IMAGE_VERSION: {{.Alias}}:v1.0
    K8S_ROLE_TOKEN: ${K8sRoleToken}
    K8S_ROLE_NAME: ${K8sRoleName}
stages:
  - build
  - push_image
  - deploy
build:
  stage: build
  image: docker:stable
  script:
    - docker build  -t registry.cn-hangzhou.aliyuncs.com/kafuluote/$IMAGE_VERSION .
  allow_failure: false
  only:
    - dev

  tags:
    - cms

push_image:
  image: docker:stable
  stage: push_image
  before_script:
    - docker login -u lyblxy@126.com -p zhuanke8  registry.cn-hangzhou.aliyuncs.com
  script:
      - echo "push image "
      - docker push registry.cn-hangzhou.aliyuncs.com/kafuluote/{{.Alias}}:v1.0
  allow_failure: false

  only:
    - dev

  tags:
    - cms

deploy:
  image: registry.cn-hangzhou.aliyuncs.com/kafuluote/kubectl:v1.0
  stage: deploy
  script:
    - kubectl config set-credentials $K8S_ROLE_NAME  --token $K8S_ROLE_TOKEN
    - kubectl config set-cluster ts --server https://192.168.70.128:6443
    - kubectl config set-context $K8S_ROLE_NAME@ts/kong --user $K8S_ROLE_NAME --cluster ts --namespace kong
    - kubectl config use-context $K8S_ROLE_NAME@ts/kong
    - kubectl set image -n kong deployments/{{.Alias}} {{.Alias}}=registry.cn-hangzhou.aliyuncs.com/kafuluote/$IMAGE_VERSION  --insecure-skip-tls-verify
  only:
    - master`

	Deploy = `---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: {{.Alias}}
  name: {{.Alias}}
  namespace: kong
spec:
  ports:
    - name: http
      port: 8000
      protocol: TCP
      targetPort: 8080
  selector:
    app: {{.Alias}}
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: {{.Alias}}
  name: {{.Alias}}
  namespace: kong
spec:
  selector:
    matchLabels:
      app: {{.Alias}}
  replicas: 1

  template:
    metadata:
      labels:
        app: {{.Alias}}
    spec:
      imagePullSecrets:
        - name: registry-secret
      containers:
        - image: registry.cn-hangzhou.aliyuncs.com/kafuluote/{{.Alias}}:v1.0
          name: {{.Alias}}
          ports:
            - containerPort: 8080`
)
