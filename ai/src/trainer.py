import matplotlib.pyplot as plt
import numpy as np
import tensorflow as tf
from keras import layers, models
from pandas import DataFrame
from sklearn.metrics import ConfusionMatrixDisplay
from sklearn.metrics import classification_report
from sklearn.model_selection import train_test_split
from sklearn.utils.class_weight import compute_class_weight
from tensorflow import keras

from src.config import SENTIMENTS

MAX_FEATURES = 5_000
MAX_SEQ_LEN = 256
EMBEDDING_DIM = 128
RANDOM_STATE = 50


def plot_history(history):
    acc, val_acc = history['accuracy'], history['val_accuracy']
    loss, val_loss = history['loss'], history['val_loss']
    x = range(1, len(acc) + 1)

    plt.figure(figsize=(12, 5))
    plt.subplot(1, 2, 1)

    plt.plot(x, acc, 'b', label='Training acc')
    plt.plot(x, val_acc, 'r', label='Validation acc')
    plt.title('Training and validation accuracy')
    plt.legend()
    plt.subplot(1, 2, 2)

    plt.plot(x, loss, 'b', label='Training loss')
    plt.plot(x, val_loss, 'r', label='Validation loss')
    plt.title('Training and validation loss')
    plt.legend()
    plt.show()


class Trainer:

    def __init__(self, df: DataFrame):
        df.sentiment = df.sentiment.cat.remove_unused_categories()
        df.sentiment = df.sentiment.cat.reorder_categories(SENTIMENTS)

        self.train_df, self.test_df = train_test_split(df, test_size=0.2, random_state=RANDOM_STATE)

        print(f'{len(self.train_df)=}, {len(self.test_df)=}')

        self.class_weight = dict(enumerate(
            compute_class_weight(
                class_weight="balanced",
                classes=np.unique(df.sentiment),
                y=df.sentiment
            )
        ))

    def learn(self) -> models.Sequential:
        layer = layers.TextVectorization(
            max_tokens=MAX_FEATURES,
            output_sequence_length=MAX_SEQ_LEN,
            output_mode='int'
        )
        layer.adapt(self.train_df.content)

        model = models.Sequential([
            keras.Input(shape=(1,), dtype=tf.string),
            layer,
            layers.Embedding(MAX_FEATURES, EMBEDDING_DIM),

            layers.SpatialDropout1D(0.2),
            layers.GlobalMaxPooling1D(),
            layers.Dropout(0.4),
            layers.Dense(256, activation='gelu'),
            layers.Dropout(0.4),
            layers.Dense(len(SENTIMENTS), activation='softmax'),
        ])

        model.compile(
            optimizer='adam',
            loss='sparse_categorical_crossentropy',
            metrics=['accuracy']
        )

        x_train, y_train = self.train_df.content, self.train_df.sentiment.cat.codes
        x_test, y_test = self.test_df.content, self.test_df.sentiment.cat.codes

        history = model.fit(
            x=x_train,
            y=y_train,
            validation_data=(x_test, y_test),
            batch_size=256,
            epochs=3,
            verbose=1,
            class_weight=self.class_weight,
            callbacks=[keras.callbacks.EarlyStopping(patience=3)],
        )
        test_loss, test_acc = model.evaluate(x_test, y_test, verbose=2)

        plot_history(history.history)

        y_pred = model.predict(x_test).argmax(1)

        print(classification_report(
            y_test, y_pred, target_names=SENTIMENTS
        ))
        ConfusionMatrixDisplay.from_predictions(
            y_test, y_pred, display_labels=SENTIMENTS
        )

        return model
