// Appointment Model
class Appointment {
  final int appointmentID;
  final int userID;
  final int salonID;
  final int serviceID;
  final String dateTime;
  final String status;
  final String notificationSettings;

  Appointment({
    this.appointmentID,
    this.userID,
    this.salonID,
    this.serviceID,
    this.dateTime,
    this.status,
    this.notificationSettings,
  });

  factory Appointment.fromJson(Map<String, dynamic> json) {
    return Appointment(
      appointmentID: json['appointment_id'],
      userID: json['user_id'],
      salonID: json['salon_id'],
      serviceID: json['service_id'],
      dateTime: json['date_time'],
      status: json['status'],
      notificationSettings: json['notification_settings'],
    );
  }
}
