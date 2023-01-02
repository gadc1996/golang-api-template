install:
	cp .env.example .env
	ln -s -f ${PWD}/hooks/pre-commit ${PWD}/.git/hooks/pre-commit 
test:
	docker build -q -t tests -f docker/Dockerfile.test .  && docker run tests
dev:
	docker-compose up --file docker-compose.dev.yml -d && ${BROWSER} http://localhost:8080
clean:
	docker-compose down
architecture:
	eb init --platform docker --region us-east-1 ${APP_NAME} 
	eb create --region ${AWS_REGION} --platform docker --single Production
deploy:
	eb deploy
terminate:
	eb terminate
	aws elasticbeanstalk delete-application --application-name ${APP_NAME} --region ${AWS_REGION}
	rm -rf .elasticbeanstalk
