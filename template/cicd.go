package template

var (
	Cicd = `stages:
  - build
  - push_image
  - deploy
build:
  stage: build
  image: docker:stable
  script:
    - docker build  -t registry.cn-hangzhou.aliyuncs.com/kafuluote/{{.Alias}}:v1.0 .
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
    - kubectl config set-credentials dashboard-admin-token-gzgqd --token eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkYXNoYm9hcmQtYWRtaW4tdG9rZW4tZ3pncWQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGFzaGJvYXJkLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiMGFkYmVlNTEtOTJiMS00YjU3LThjODctZDU4MmI4MDQ5ZmU5Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOmRhc2hib2FyZC1hZG1pbiJ9.GazzLI1Pgc6mfz6EreEvaVKvGcSw6a_Y2OE4N3hYd0_B7D-gvlwIVgniVtDRpRd5QQE0uv-QANmk5Mmqkr4jgjGL6gDZ04QKhJnJAU_XBouugTBzI0-cUOynQqF_aZXsL305dT-DG_d3lrMJSsuc50fUQlsjeitQ3UDqJ_mAgsNCGlOR4emMPYhwyuKw3qWWSgE-I1VyLGpMzE6CrISSD_tyq3Vjr-M3lQPKzaKo_5R0rdGnEZXYgkZa9DaQDDj1_FXoWom_u-yA8k0Pash4PRu2sQ3uoG_eHuwBt3ggTCnm9Mkm1U3AMpyTGMrrbeg3V3TIBg5evkJmd55zqzc9LQ
    - kubectl config set-cluster ts --server https://192.168.70.128:6443
    - kubectl config set-context dashboard-admin-token-gzgqd@ts/kong --user dashboard-admin-token-gzgqd --cluster ts --namespace kong
    - kubectl config use-context dashboard-admin-token-gzgqd@ts/kong
    - kubectl set image -n kong deployments/{{.Alias}} {{.Alias}}=registry.cn-hangzhou.aliyuncs.com/kafuluote/{{.Alias}}:v1.0  --insecure-skip-tls-verify
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
