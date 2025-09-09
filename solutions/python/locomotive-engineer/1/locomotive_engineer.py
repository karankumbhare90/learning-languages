"""Functions which helps the locomotive engineer to keep track of the train."""


def get_list_of_wagons(*wagon_ids):
    """Return a list of wagons.

    :param: arbitrary number of wagons.
    :return: list - list of wagons.
    """
    return list(wagon_ids)


def fix_list_of_wagons(each_wagons_id, missing_wagons):
    """Fix the list of wagons.

    :param each_wagons_id: list - the list of wagons.
    :param missing_wagons: list - the list of missing wagons.
    :return: list - list of wagons.
    """
    # Step 1: move first two wagons to the end
    reordered = each_wagons_id[2:] + each_wagons_id[:2]

    # Step 2: find the locomotive
    loco_index = reordered.index(1)

    # Step 3: insert missing wagons after locomotive
    fixed = reordered[:loco_index+1] + missing_wagons + reordered[loco_index+1:]
    return fixed


def add_missing_stops(route, **stops):
    """Add missing stops to route dict.

    :param route: dict - the dict of routing information.
    :param stops: keyword arguments - stop information.
    :return: dict - updated route dictionary.
    """
    ordered_stops = [stops[key] for key in sorted(stops.keys())]
    route["stops"] = ordered_stops
    return route


def extend_route_information(route, more_route_information):
    """Extend route information with more_route_information.

    :param route: dict - the route information.
    :param more_route_information: dict -  extra route information.
    :return: dict - extended route information.
    """
    extended = route.copy()
    extended.update(more_route_information)
    return extended


def fix_wagon_depot(wagons_rows):
    """Fix the list of rows of wagons.

    :param wagons_rows: list[list[tuple]] - the list of rows of wagons.
    :return: list[list[tuple]] - list of rows of wagons.
    """
    fixed_rows = list(map(list, zip(*wagons_rows)))
    return fixed_rows
