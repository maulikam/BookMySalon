// Review Model
class Review {
  final int reviewID;
  final int userID;
  final int salonID;
  final int rating;
  final String comment;
  final String datePosted;

  Review({
    this.reviewID,
    this.userID,
    this.salonID,
    this.rating,
    this.comment,
    this.datePosted,
  });

  // Factory constructor to create an instance of Review from a JSON map.
  factory Review.fromJson(Map<String, dynamic> json) {
    return Review(
      reviewID: json['review_id'],
      userID: json['user_id'],
      salonID: json['salon_id'],
      rating: json['rating'],
      comment: json['comment'],
      datePosted: json['date_posted'],
    );
  }
}
