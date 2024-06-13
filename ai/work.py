import sys

import pika
from keras import models

from src import config
from src.worker import Worker

PATH_TO_MODEL = 'ai/models/model.keras'


def main():
    model = models.load_model(PATH_TO_MODEL)

    parameters = pika.URLParameters(config.RABBIT_MQ_DSN)
    connection = pika.BlockingConnection(parameters)
    channel = connection.channel()

    channel.queue_declare(queue=config.QUEUE_TO_ANALYSIS, durable=True)
    channel.queue_declare(queue=config.QUEUE_ANALYSIS_RESULT, durable=True)

    w = Worker(model, channel)
    w.start()


if __name__ == '__main__':
    sys.stdout.flush()
    main()
