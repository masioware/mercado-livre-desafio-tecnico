from random import randint, sample


def get_distribution_centers(max_list_size=5, max_available_centers=20):
    """Returns a random list of distribution centers with a limited size.
    - With no duplicates and a random length up to max_list_size!
        - Minimum 1 and max max_list_size (default 10).
        - Centers are chosen from a pool of 1 to max_available_centers.
    """
    num_centers = randint(1, max_list_size)
    all_possible_centers = range(1, max_available_centers + 1)
    chosen_centers = sample(all_possible_centers, num_centers)
    return [f"CD{num}" for num in chosen_centers]
