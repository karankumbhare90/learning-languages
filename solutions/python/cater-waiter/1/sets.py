"""Functions for compiling dishes and ingredients for a catering company."""


from sets_categories_data import (VEGAN,
                                  VEGETARIAN,
                                  KETO,
                                  PALEO,
                                  OMNIVORE,
                                  ALCOHOLS,
                                  SPECIAL_INGREDIENTS)


def clean_ingredients(dish_name, dish_ingredients):
    # Remove duplicates
    return (dish_name, set(dish_ingredients))


def check_drinks(drink_name, drink_ingredients):
    # """Append "Cocktail" (alcohol)  or "Mocktail" (no alcohol) to `drink_name`, based on `drink_ingredients`.

    if ALCOHOLS.intersection(drink_ingredients):
        return f"{drink_name} Cocktail"
    else : 
        return f"{drink_name} Mocktail"


def categorize_dish(dish_name, dish_ingredients):
    # """Categorize `dish_name` based on `dish_ingredients`.
    if dish_ingredients.issubset(VEGAN):
        category = "VEGAN"
    elif dish_ingredients.issubset(VEGETARIAN):
        category = "VEGETARIAN"
    elif dish_ingredients.issubset(PALEO):
        category = "PALEO"
    elif dish_ingredients.issubset(KETO):
        category = "KETO"
    else : 
        category = "OMNIVORE"

    return f"{dish_name}: {category}"
    


def tag_special_ingredients(dish):
    """Compare `dish` ingredients to `SPECIAL_INGREDIENTS`
    """

    dish_name, ingredients = dish
    return (dish_name, set(ingredients).intersection(SPECIAL_INGREDIENTS))


def compile_ingredients(dishes):
   return set().union(*dishes)


def separate_appetizers(dishes, appetizers):
    return list(set(dishes) - set(appetizers))


def singleton_ingredients(dishes, intersection):
    all_ingredients = set().union(*dishes)

    return all_ingredients - intersection
