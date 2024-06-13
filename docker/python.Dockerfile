FROM tensorflow/tensorflow:latest-gpu-jupyter

COPY ai/requirements.txt /var/app/ai/requirements.txt
WORKDIR /var/app
RUN pip install --no-cache-dir --upgrade pip \
    && pip install --no-cache-dir -r ai/requirements.txt \
    && rm -rf /root/.cache/pip

#RUN python -m nltk.downloader stopwords