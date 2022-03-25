run:
	docker run -p 8080:8080 --name contact-book --rm sprow/contact-book:v1.0
stop:
	docker stop contact-book