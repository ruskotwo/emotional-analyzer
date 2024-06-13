import logging

from keras import Sequential
from pika.adapters.blocking_connection import BlockingChannel

from src.config import QUEUE_TO_ANALYSIS


class Worker:

    def __init__(self, model: Sequential, channel: BlockingChannel):
        self.model = model
        self.channel = channel

    def start(self):
        self.channel.basic_consume(queue=QUEUE_TO_ANALYSIS, on_message_callback=self.handle)
        self.channel.start_consuming()

    def handle(self, channel, method, properties, body):

        print('y vzal')

        data = body.decode()

        print(data)

        channel.basic_ack(delivery_tag=method.delivery_tag)
