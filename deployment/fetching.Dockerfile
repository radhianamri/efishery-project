from python:3.7.8-alpine3.12


WORKDIR /app
COPY fetching-py/ .


RUN echo "http://dl-8.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    apk update && \
    apk --no-cache --update-cache add python3-dev build-base openssl-dev libffi-dev gcc libgcc && \
    pip install -r prereq.txt && \
    pip install -r requirements.txt && \
    apk del  python3-dev build-base openssl-dev libffi-dev gcc libgcc && \
    rm -rf /var/cache/apk/* /var/tmp/* /tmp/* /root/.cache/*



EXPOSE 6000
CMD python3 manage.py --deployment_type PRODUCTION

