class User {
  final int id;
  final String username;
  final String password; // Ideally, the password should never be exposed to the client.
  final String email;
  final String profileImage;
  final String dateJoined;
  final String lastLogin;

  User({
    this.id,
    this.username,
    this.password,
    this.email,
    this.profileImage,
    this.dateJoined,
    this.lastLogin,
  });

  // Factory constructor to create an instance of User from a JSON map.
  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      username: json['username'],
      password: json['password'],
      email: json['email'],
      profileImage: json['profile_image'],
      dateJoined: json['date_joined'],
      lastLogin: json['last_login'],
    );
  }
}
