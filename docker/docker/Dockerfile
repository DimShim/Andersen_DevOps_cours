FROM alpine:3.7

# specifying ports for listening
EXPOSE 80

# copy in container src files
COPY src/nginx.conf \
	app/app.py \
	app/templates/index.html \
        app/templates/description.html \
	app/templates/preview.html /

# install requiered components
RUN apk add --update --no-cache \
	nginx \
	curl \
	python3 \
	py3-flask \
&&	wget -c https://files.pythonhosted.org/packages/d6/56/\
484163d2f941f8c96de24fecf4885e3101913a2f4b31a1d50b4eb8eb244a/\
emoji-1.2.0.tar.gz \
&&	tar -xvf emoji-1.2.0.tar.gz \
&&	cd emoji-1.2.0 \
&&	find . -mtime +10950 -print -exec touch {} \; \
&&	python3 setup.py install \
&&	rm -rf /emoji-1.2.0

# generate directories 
RUN	mkdir -p /etc/supervisor.d \	
			/run/nginx

# move src files to correct locations
RUN	rm /etc/nginx/conf.d/default.conf \
&&	mv /nginx.conf /etc/nginx/conf.d/ \
&&	mkdir -p /emojis_loopback/templates \
&&	mv /index.html /preview.html /description.html /emojis_loopback/templates/ \
&&	mv app.py /emojis_loopback/

# launch server
CMD	python3 /emojis_loopback/app.py & \
	nginx -g 'daemon off;'
