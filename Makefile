build:
	docker build --no-cache -t cache .
run:
	docker run -p 3000:3000 cache
