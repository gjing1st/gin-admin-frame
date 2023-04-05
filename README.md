#
## 1 基本介绍
### 1.1 项目简介
基于gin的后台管理框架


## 2 使用说明
### 直接打包
cmd/publish/main.go
### docker化
使用./script/docker/Dockerfile打包
docker build -f ./script/docker/Dockerfile -t gaf:latest .
### 基于gitlab的自动打包
可参考.gitlab-ci.yml
### 基于KubeSphere的DevOps
可参考./script/jenkins和./script/k8s 可实现自动打包并部署至k8s


