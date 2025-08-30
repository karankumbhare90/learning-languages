"""Functions to help play and score a game of blackjack."""

def value_of_card(card):
    """Determine the scoring value of a card.
    
    Face cards (J, Q, K) = 10
    Ace (A) = 1 (for now)
    Number cards = their value
    """
    if card in ['J', 'Q', 'K']:
        return 10
    elif card == 'A':
        return 1
    else:
        return int(card)


def higher_card(card_one, card_two):
    """Return the higher value card, or both if equal."""
    v1, v2 = value_of_card(card_one), value_of_card(card_two)
    if v1 > v2:
        return card_one
    elif v2 > v1:
        return card_two
    else:
        return card_one, card_two


def value_of_ace(card_one, card_two):
    """Return the most advantageous value for an upcoming Ace (1 or 11)."""
    v1, v2 = value_of_card(card_one), value_of_card(card_two)

    # Special handling if one of the cards is Ace
    if card_one == 'A' or card_two == 'A':
        return 1

    # If 11 doesn’t bust (<= 21), return 11; otherwise return 1
    if v1 + v2 + 11 <= 21:
        return 11
    return 1


def is_blackjack(card_one, card_two):
    """Return True if the hand is a blackjack (Ace + 10-value card)."""
    ten_cards = ['10', 'J', 'Q', 'K']
    return (card_one == 'A' and card_two in ten_cards) or \
           (card_two == 'A' and card_one in ten_cards)


def can_split_pairs(card_one, card_two):
    """Return True if the cards can be split into pairs (same value)."""
    return value_of_card(card_one) == value_of_card(card_two)


def can_double_down(card_one, card_two):
    """Return True if the two-card hand totals 9, 10, or 11."""
    total = value_of_card(card_one) + value_of_card(card_two)
    return total in [9, 10, 11]
