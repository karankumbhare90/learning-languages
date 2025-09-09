"""Functions to manage a users shopping cart items."""


def add_item(current_cart, items_to_add):
    """
        Add items to shopping cart.
        :param current_cart: dict - the current shopping cart.
        :param items_to_add: iterable - items to add to the cart.
        :return: dict - the updated user cart dictionary.
    """
    update_car = current_cart.copy()

    for item in items_to_add : 
        update_car[item] = update_car.get(item, 0) + 1

    return update_car

def read_notes(notes):
    """
        Create user cart from an iterable notes entry.
        :param notes: iterable of items to add to cart.
        :return: dict - a user shopping cart dictionary.
    """
    cart = {}
    for item in notes :
        cart[item] = 1

    return cart

def update_recipes(ideas, recipe_updates):
    """
        Update the recipe ideas dictionary.
        :param ideas: dict - The "recipe ideas" dict.
        :param recipe_updates: iterable -  with updates for the ideas section.
        :return: dict - updated "recipe ideas" dict.
    """
    updated_ideas = ideas.copy()

    for reciepe_name, new_ingredients in recipe_updates:
        updated_ideas[reciepe_name] = new_ingredients

    return updated_ideas

def sort_entries(cart):
    """Sort a users shopping cart in alphabetically order.

    :param cart: dict - a users shopping cart dictionary.
    :return: dict - users shopping cart sorted in alphabetical order.
    """

    return dict(sorted(cart.items(), key=lambda x : x[0]))


def send_to_store(cart, aisle_mapping):
    """Combine users order to aisle and refrigeration information.

    :param cart: dict - users shopping cart dictionary.
    :param aisle_mapping: dict - aisle and refrigeration information dictionary.
    :return: dict - fulfillment dictionary ready to send to store.
    """
    fulfillment_cart = {}
    for item, qty in cart.items():
        aisle, refrigeration = aisle_mapping[item]
        fulfillment_cart[item] = [qty, aisle, refrigeration]

    # Sort in reverse alphabetical order
    return dict(sorted(fulfillment_cart.items(), key=lambda x: x[0], reverse=True))


def update_store_inventory(fulfillment_cart, store_inventory):
    """Update store inventory levels with user order.

    :param fulfillment_cart: dict - fulfillment cart to send to store.
    :param store_inventory: dict - store available inventory
    :return: dict - store_inventory updated.
    """
    updated_inventory = store_inventory.copy()
    for item, (ordered_qty, aisle, refrigeration) in fulfillment_cart.items():
        current_qty, aisle_info, refrigeration_info = updated_inventory[item]
        remaining = current_qty - ordered_qty
        if remaining <= 0:
            updated_inventory[item] = ['Out of Stock', aisle_info, refrigeration_info]
        else:
            updated_inventory[item] = [remaining, aisle_info, refrigeration_info]
    return updated_inventory