# Template for training model 
# when input datas need lots of memory (like images)
from queue import Queue
import threading
import time
import numpy as np

BATCH_QUEUE = Queue(3)
datas = list(range(100))
labels = list(range(100))

def load_data(folder, batch_size=16):
    # Load image files from disk and do some
    # augmentation or preprocessing
    while True:
        X_batch = []
        Y_batch = []
        for start in range(0, len(datas), batch_size):
            X_batch = np.array( datas[start:start+batch_size])
            Y_batch = np.array(labels[start:start+batch_size])
            
            BATCH_QUEUE.put((X_batch, Y_batch))
            print("[BATCH_QUEUE] PUT:", X_batch)


def generator(batch_size):
    threading.Thread(target=load_data, args=("folder", batch_size), daemon=True).start()
    while True:
        yield BATCH_QUEUE.get()


def fit(batch_size, epochs):
    gen = generator(batch_size)
    steps_per_epoch = len(datas) // batch_size + 1
    for e in range(1, epochs+1):
        for step in range(steps_per_epoch):
            data, label = next(gen)
            print(f"[EPOCH-{e}][STEP-{step}] GET: {data}")
            time.sleep(2)


if __name__ == "__main__":
    fit(batch_size=16, epochs=5)
