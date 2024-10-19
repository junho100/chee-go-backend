url=209479273800.dkr.ecr.ap-northeast-2.amazonaws.com/prod-cheego556-backend-ecr

aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin "$(aws sts get-caller-identity --query Account --output text).dkr.ecr.ap-northeast-2.amazonaws.com"
docker pull $url:latest
docker stop cheego || true
docker rm cheego || true
docker run -d --name cheego -p 8080:8080 --log-driver=awslogs --log-opt awslogs-region="ap-northeast-2" --log-opt awslogs-group=/cheego/backend --log-opt awslogs-stream=backend --log-opt awslogs-create-group=true $url:latest
