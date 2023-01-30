EXECUTABLE=new-years-eve


.PHONY=clean start

## Application
start:
	docker-compose down
	docker-compose up -d
	MONGODB_URI="mongodb://mongo:mongo@localhost:27017" $(MAKE) -C backend/ start

## Infrastructure

deploy: clean
	cd backend && $(MAKE) prepare-for-deployment
	mv backend/api .
	cd infrastructure && $(MAKE) deploy

clean:
	rm -rf api || true
