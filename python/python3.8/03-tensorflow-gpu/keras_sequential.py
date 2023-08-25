import tensorflow as tf
import datetime

#tf.debugging.set_log_device_placement(True)

x = tf.linspace(-2, 2, 201)
x = tf.cast(x, tf.float32)

def f(x):
  y = x**2 + 2*x - 5
  return y

y = f(x) + tf.random.normal(shape=[201])

new_model = tf.keras.Sequential([
    tf.keras.layers.Lambda(lambda x: tf.stack([x, x**2], axis=1)),
    tf.keras.layers.Dense(units=1,
kernel_initializer=tf.random.normal)])

new_model.compile(
    loss=tf.keras.losses.MSE,
    optimizer=tf.keras.optimizers.SGD(learning_rate=0.01))

tf.print(tf.constant('Fitting model...'))
history = new_model.fit(x, y,
                        epochs=100,
                        batch_size=32,
                        verbose=0)

new_model.save('./my_new_model')
tf.print(tf.constant('Model saved'))
while True:
        tf.print(datetime.datetime.now().strftime("%H:%M:%S.%f")[:-3])
        new_model.fit(x, y, epochs=100, batch_size=32, verbose=0)
