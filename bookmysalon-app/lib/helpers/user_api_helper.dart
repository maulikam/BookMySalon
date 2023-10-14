import 'dart:convert';
import 'package:http/http.dart' as http;
import 'models/user.dart';  // Import your User model here

class ApiHelper {
  final String _baseUrl = "http://localhost:8080";  // Change to your server's IP if not on localhost

  // Register
  Future<String?> registerUser(User user) async {
    final response = await http.post(
      Uri.parse("$_baseUrl/register"),
      headers: {
        'Content-Type': 'application/json',
      },
      body: jsonEncode(user.toJson()),
    );

    if (response.statusCode == 200) {
      return response.body;
    } else {
      throw Exception("Failed to register user: ${response.body}");
    }
  }

  // Login
  Future<String?> loginUser(User user) async {
    final response = await http.post(
      Uri.parse("$_baseUrl/login"),
      headers: {
        'Content-Type': 'application/json',
      },
      body: jsonEncode(user.toJson()),
    );

    if (response.statusCode == 200) {
      return response.body;
    } else {
      throw Exception("Failed to login: ${response.body}");
    }
  }

  // Fetch Profile
  Future<User?> getUserProfile(String token) async {
    final response = await http.get(
      Uri.parse("$_baseUrl/profile"),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );

    if (response.statusCode == 200) {
      return User.fromJson(jsonDecode(response.body));
    } else {
      throw Exception("Failed to fetch profile: ${response.body}");
    }
  }

// Update Profile
  Future<void> updateUserProfile(User user, String token) async {
    final response = await http.put(
      Uri.parse("$_baseUrl/profile"),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode(user.toJson()),
    );

    if (response.statusCode != 200) {
      throw Exception("Failed to update profile: ${response.body}");
    }
  }

  // Change Password
  Future<void> changeUserPassword(
      String oldPassword, String newPassword, String token) async {
    final response = await http.put(
      Uri.parse("$_baseUrl/change-password"),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode({
        'oldPassword': oldPassword,
        'newPassword': newPassword,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception("Failed to change password: ${response.body}");
    }
  }

  // Delete Account
  Future<void> deleteAccount(String token) async {
    final response = await http.delete(
      Uri.parse("$_baseUrl/profile"),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );

    if (response.statusCode != 200) {
      throw Exception("Failed to delete account: ${response.body}");
    }
  }}

