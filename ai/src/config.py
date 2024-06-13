import os

SENTIMENTS = ['anger', 'boredom', 'enthusiasm', 'fun', 'happiness', 'hate', 'love', 'neutral', 'relief', 'sadness',
              'surprise', 'worry']

RABBIT_MQ_DSN = os.environ['RABBIT_MQ_DSN']

QUEUE_TO_ANALYSIS = os.environ['QUEUE_TO_ANALYSIS']
QUEUE_ANALYSIS_RESULT = os.environ['QUEUE_ANALYSIS_RESULT']
