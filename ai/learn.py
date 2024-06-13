import os
import sys

import pandas as pd
from keras import models

from src.trainer import Trainer

PATH_TO_MODEL = 'ai/models/model.keras'
PATH_TO_DATASET = 'ai/datasets/tweet_emotions.csv'

def main():
    if len(sys.argv) > 1 and sys.argv[1] == 'force':
        print('Force create model')
    elif os.path.isfile(PATH_TO_MODEL):
        try:
            models.load_model(PATH_TO_MODEL)
            print('Model already exists')
            return
        except Exception:
            print('Invalid model, try re-create')
    else:
        print('Model does not exist, to create')

    df = pd.read_csv(
        PATH_TO_DATASET,
        usecols=['content', 'sentiment'],
        dtype={'content': 'string', 'sentiment': 'category'}
    )

    lr = Trainer(df)
    model = lr.learn()

    model.save(PATH_TO_MODEL)
    print('Model saved')

if __name__ == '__main__':
    main()
