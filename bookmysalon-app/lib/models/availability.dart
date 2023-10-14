// Availability Model
class Availability {
  final int availabilityID;
  final int salonID;
  final int serviceID;
  final String startDateTime;
  final String endDateTime;
  final String status;

  Availability({
    this.availabilityID,
    this.salonID,
    this.serviceID,
    this.startDateTime,
    this.endDateTime,
    this.status,
  });

  factory Availability.fromJson(Map<String, dynamic> json) {
    return Availability(
      availabilityID: json['availability_id'],
      salonID: json['salon_id'],
      serviceID: json['service_id'],
      startDateTime: json['start_date_time'],
      endDateTime: json['end_date_time'],
      status: json['status'],
    );
  }
}
