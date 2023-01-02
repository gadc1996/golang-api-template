make install:
	cp .env.example .env
	ln -s -f ${PWD}/hooks/pre-commit ${PWD}/.git/hooks/pre-commit 
test:
	docker build -q -t tests -f Dockerfile.test .  && docker run tests
dev:
	docker-compose up -d && ${BROWSER} http://localhost:8080
clean:
	docker-compose down