class Vehicle:
    def __init__(self, max_speed_kph, kilometers_per_liter):
        self.max_speed_kph = max_speed_kph
        self.kilometers_per_liter = kilometers_per_liter

    def get_price_for_trip(self, distance_kilometers, price_per_liter):
        denominator = self.kilometers_per_liter * price_per_liter
        if denominator == 0:
            raise Exception("cannot divide by zero")
        return distance_kilometers / denominator

    def get_cargo_cap_meters_cubed(self):
        pass


class Truck(Vehicle):
    def __init__(
        self,
        max_speed_kph,
        kilometers_per_liter,
        load_weight_kilos,
        bed_area_meters_squared,
    ):
        super().__init__(max_speed_kph, kilometers_per_liter)
        self.load_weight_kilos = load_weight_kilos
        self.bed_area_meters_squared = bed_area_meters_squared

    def get_price_for_trip(self, distance_kilometers, price_per_liter):
        base_cost = super().get_price_for_trip(distance_kilometers, price_per_liter)
        return base_cost + (self.load_weight_kilos * 0.01)

    def get_cargo_cap_meters_cubed(self):
        return self.bed_area_meters_squared * 2


class Car(Vehicle):
    def __init__(self, max_speed_kph, kilometers_per_liter, cargo_cap_meters_cubed):
        super().__init__(max_speed_kph, kilometers_per_liter)
        self.cargo_cap_meters_cubed = cargo_cap_meters_cubed

    def get_cargo_cap_meters_cubed(self):
        return self.cargo_cap_meters_cubed


# don't touch below this line


def test_vehicle(vehicle):
    print(f"Vehicle:")
    print(f" - Max Speed KPH: {vehicle.max_speed_kph:.2f}")
    print(f" - Kilometers per liter: {vehicle.kilometers_per_liter:.2f}")
    print(f"Calculating gas cost for 100 kilometers at $3.54 per liter...")
    print(f" - Gas cost: ${vehicle.get_price_for_trip(100, 3.54):.2f}")
    print("=====================================")


def test_truck(truck):
    print(f"Truck:")
    print(f" - Max Speed KPH: {truck.max_speed_kph:.2f}")
    print(f" - Kilometers per liter: {truck.kilometers_per_liter:.2f}")
    print(f" - Load Weight in Kilos: {truck.load_weight_kilos:.2f}")
    print(f" - Bed Area in meters squared: {truck.bed_area_meters_squared:.2f}")
    print(f"Calculating gas cost for 100 kilometers at $3.54 per liter...")
    print(f" - Gas cost: ${truck.get_price_for_trip(100, 3.54):.2f}")
    print(f" - Calculated cargo capacity: {truck.get_cargo_cap_meters_cubed():.2f}")
    print("=====================================")


def test_car(car):
    print(f"Car:")
    print(f" - Max Speed KPH: {car.max_speed_kph:.2f}")
    print(f" - Kilometers per liter: {car.kilometers_per_liter:.2f}")
    print(f" - Cargo Capacity in meters cubed: {car.cargo_cap_meters_cubed:.2f}")
    print(f"Calculating gas cost for 100 kilometers at $3.54 per liter...")
    print(f" - Gas cost: ${car.get_price_for_trip(100, 3.54):.2f}")
    print(f" - Calculated cargo capacity: {car.get_cargo_cap_meters_cubed():.2f}")
    print("=====================================")


def main():
    test_vehicle(Vehicle(100, 10))
    test_truck(Truck(100, 10, 2000, 5.0))
    test_car(Car(100, 10, 2.0))


main()
