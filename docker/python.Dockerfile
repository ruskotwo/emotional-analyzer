FROM tensorflow/tensorflow:latest-gpu-jupyter

RUN apt update && apt install supervisor

ADD . /var/app/
COPY ai/supervisor/ /etc/supervisor/conf.d/

WORKDIR /var/app
RUN pip install --no-cache-dir --upgrade pip \
    && pip install --no-cache-dir -r ai/requirements.txt \
    && rm -rf /root/.cache/pip

