docker-compose -p histogram_demo down
docker network rm app_backend

echo "Starting environment"
cd infra
docker-compose -p histogram_demo up -d --build
echo "Environment started"

echo "Now, 'generator' generates and sends requests in 'app' application"
echo "Metrics available on: http://localhost:8080/metrics"
echo "Prometheus available on: http://localhost:9090"
echo "Grafana available on: http://localhost:3000/d/KYu1HT0Hl/http-servers-monitoring"
echo "Intervals dashboard on: http://localhost:3000/d/KYu1HT0Hl/http-servers-monitoring?viewPanel=16"
echo "Wait at least for 1 minute for having some data"