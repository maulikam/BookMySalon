// lib/helpers/salon_api_helper.dart

import 'dart:convert';
import 'package:http/http.dart' as http;

class SalonApiHelper {
  final String baseUrl = "http://localhost:8080";

  // Create a new salon
  Future<Map<String, dynamic>> createSalon(Map<String, dynamic> salonData) async {
    final response = await http.post(
      Uri.parse('$baseUrl/salon'),
      headers: {"Content-Type": "application/json"},
      body: jsonEncode(salonData),
    );

    if (response.statusCode == 201) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to create salon. Error: ${response.body}');
    }
  }

  // Get details of a salon by ID
  Future<Map<String, dynamic>> getSalonDetails(int salonID) async {
    final response = await http.get(Uri.parse('$baseUrl/salon/$salonID'));

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to fetch salon details. Error: ${response.body}');
    }
  }

  // Update salon details
  Future<void> updateSalonDetails(Map<String, dynamic> salonData) async {
    final response = await http.put(
      Uri.parse('$baseUrl/salon/update'),
      headers: {"Content-Type": "application/json"},
      body: jsonEncode(salonData),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to update salon. Error: ${response.body}');
    }
  }

  // Continue from previous `SalonApiHelper` class...

  // Delete a salon by ID
  Future<void> deleteSalon(int salonID) async {
    final response = await http.delete(Uri.parse('$baseUrl/salon/$salonID'));

    if (response.statusCode != 200) {
      throw Exception('Failed to delete salon. Error: ${response.body}');
    }
  }

  // List all salons
  Future<List<Map<String, dynamic>>> listAllSalons() async {
    final response = await http.get(Uri.parse('$baseUrl/salons'));

    if (response.statusCode == 200) {
      List<dynamic> responseBody = jsonDecode(response.body);
      return responseBody.cast<Map<String, dynamic>>();
    } else {
      throw Exception('Failed to list all salons. Error: ${response.body}');
    }
  }

  // Search salons by name
  Future<List<Map<String, dynamic>>> searchSalonsByName(String name) async {
    final response = await http.get(Uri.parse('$baseUrl/salons/search?name=$name'));

    if (response.statusCode == 200) {
      List<dynamic> responseBody = jsonDecode(response.body);
      return responseBody.cast<Map<String, dynamic>>();
    } else {
      throw Exception('Failed to search salons by name. Error: ${response.body}');
    }
  }


 }
