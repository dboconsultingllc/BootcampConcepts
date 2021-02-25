# Bootcamp Concepts



Collection of notes and sample code during 10 week cloud skills bootcamp.

Topics Covered

* Python Code.

* Powershell Code.

* Linters and Testers.

Blah blah blah
[another link here](https://google.com).

more info here.

## Getting Started
Links to Repo below.

Sample code snippet:

```python
import model_search
from model_search import constants
from model_search import single_trainer
from model_search.data import csv_data

trainer = single_trainer.SingleTrainer(
    data=csv_data.Provider(
        label_index=0,
        logits_dimension=2,
        record_defaults=[0, 0, 0, 0],
        filename="model_search/data/testdata/csv_random_data.csv"),
    spec=constants.DEFAULT_DNN)

trainer.try_models(
    number_models=200,
    train_steps=1000,
    eval_steps=100,
    root_dir="/tmp/run_example",
    batch_size=32,
    experiment_name="example",
    experiment_owner="model_search_user")
```



## Something else here
Some other code reference here `some concept`. 

```python
class Provider(object, metaclass=abc.ABCMeta):
  """A data provider interface.

  The Provider abstract class that defines three function for Estimator related
  training that return the following:
    * An input function for training and test input functions that return
      features and label batch tensors. It is responsible for parsing the
      dataset and buffering data.
    * The feature_columns for this dataset.
    * problem statement.
  """

  def get_input_fn(self, hparams, mode, batch_size: int):
    """Returns an `input_fn` for train and evaluation.

    Args:
      hparams: tf.HParams for the experiment.
      mode: Defines whether this is training or evaluation. See
        `estimator.ModeKeys`.
      batch_size: the batch size for training and eval.

    Returns:
      Returns an `input_fn` for train or evaluation.
    """

  def get_serving_input_fn(self, hparams):
    """Returns an `input_fn` for serving in an exported SavedModel.

    Args:
      hparams: tf.HParams for the experiment.

    Returns:
      Returns an `input_fn` that takes no arguments and returns a
        `ServingInputReceiver`.
    """

  @abc.abstractmethod
  def number_of_classes(self) -> int:
    """Returns the number of classes. Logits dim for regression."""

  def get_feature_columns(
      self
  ) -> List[Union[feature_column._FeatureColumn,
                  feature_column_v2.FeatureColumn]]:
    """Returns a `List` of feature columns."""
```




## Conclusion
Search for anything here:
https://google.com/