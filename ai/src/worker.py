import json
import logging

import pandas as pd
from keras import Sequential
from pika.adapters.blocking_connection import BlockingChannel

from src import config


class Worker:

    def __init__(self, model: Sequential, channel: BlockingChannel):
        self.model = model
        self.channel = channel

    def start(self):
        self.channel.basic_consume(queue=config.QUEUE_TO_ANALYSIS, on_message_callback=self.handle)
        self.channel.start_consuming()

    def handle(self, channel, method, properties, body):
        data = {}

        try:
            data = json.loads(body.decode())
        except Exception as e:
            print('Invalid task body')
            print(e)
            channel.basic_ack(delivery_tag=method.delivery_tag)

        if data.get('messages') is None:
            print('Data dont has messages')
            channel.basic_ack(delivery_tag=method.delivery_tag)

        messages = data.get('messages')
        messages2analyze =  pd.Series(messages.values())

        preds = self.model.predict(messages2analyze).argmax(1).squeeze()

        index = 0
        for key in messages.keys():
            pred = preds[index]
            messages[key] = config.SENTIMENTS[pred]
            index += 1

        data['messages'] = messages

        channel.basic_publish(
            exchange='',
            routing_key=config.QUEUE_ANALYSIS_RESULT,
            body=json.dumps(data)
        )

        channel.basic_ack(delivery_tag=method.delivery_tag)
