"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This module provides helper functions to calculate baking and preparation times
for a lasagna recipe.
"""

# Constants
EXPECTED_BAKE_TIME = 40        # minutes
PREPARATION_TIME = 2           # minutes per layer


def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed (in minutes).
    :return: int - remaining bake time (in minutes).

    Takes the actual minutes the lasagna has been baking and returns how many
    minutes it still needs to bake based on the EXPECTED_BAKE_TIME.
    """
    return EXPECTED_BAKE_TIME - elapsed_bake_time


def preparation_time_in_minutes(number_of_layers):
    """Calculate the preparation time.

    :param number_of_layers: int - the number of lasagna layers.
    :return: int - total preparation time (in minutes).

    Assumes each layer takes PREPARATION_TIME minutes to prepare.
    """
    return PREPARATION_TIME * number_of_layers


def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Calculate the total elapsed cooking time.

    :param number_of_layers: int - number of layers in the lasagna.
    :param elapsed_bake_time: int - elapsed baking time (in minutes).
    :return: int - total time spent cooking (in minutes).

    Adds preparation time to the elapsed baking time to get the total.
    """
    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
