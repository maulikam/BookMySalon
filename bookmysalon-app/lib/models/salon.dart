// Salon Model
class Salon {
  final int salonID;
  final String name;
  final String address;
  final String contactDetails;
  final String photos;
  final double averageRating;

  Salon({
    this.salonID,
    this.name,
    this.address,
    this.contactDetails,
    this.photos,
    this.averageRating,
  });

  // Factory constructor to create an instance of Salon from a JSON map.
  factory Salon.fromJson(Map<String, dynamic> json) {
    return Salon(
      salonID: json['salon_id'],
      name: json['name'],
      address: json['address'],
      contactDetails: json['contact_details'],
      photos: json['photos'],
      averageRating: json['average_rating'],
    );
  }
}

// Service Model
class Service {
  final int serviceID;
  final int salonID;
  final String name;
  final String description;
  final String duration;
  final double price;

  Service({
    this.serviceID,
    this.salonID,
    this.name,
    this.description,
    this.duration,
    this.price,
  });

  // Factory constructor to create an instance of Service from a JSON map.
  factory Service.fromJson(Map<String, dynamic> json) {
    return Service(
      serviceID: json['service_id'],
      salonID: json['salon_id'],
      name: json['name'],
      description: json['description'],
      duration: json['duration'],
      price: json['price'],
    );
  }
}
