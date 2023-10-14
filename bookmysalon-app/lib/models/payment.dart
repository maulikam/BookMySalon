]// Transaction Model
class Transaction {
  final int transactionID;
  final int userID;
  final double amount;
  final String date;
  final String status;
  final String paymentMethod;

  Transaction({
    this.transactionID,
    this.userID,
    this.amount,
    this.date,
    this.status,
    this.paymentMethod,
  });

  factory Transaction.fromJson(Map<String, dynamic> json) {
    return Transaction(
      transactionID: json['transaction_id'],
      userID: json['user_id'],
      amount: json['amount'].toDouble(),
      date: json['date'],
      status: json['status'],
      paymentMethod: json['payment_method'],
    );
  }
}

// Invoice Model
class Invoice {
  final int invoiceID;
  final int transactionID;
  final String details;
  final String dateIssued;

  Invoice({
    this.invoiceID,
    this.transactionID,
    this.details,
    this.dateIssued,
  });

  factory Invoice.fromJson(Map<String, dynamic> json) {
    return Invoice(
      invoiceID: json['invoice_id'],
      transactionID: json['transaction_id'],
      details: json['details'],
      dateIssued: json['date_issued'],
    );
  }
}

// Promotion Model
class Promotion {
  final int promotionID;
  final String description;
  final double discountAmount;
  final String validFrom;
  final String validTo;

  Promotion({
    this.promotionID,
    this.description,
    this.discountAmount,
    this.validFrom,
    this.validTo,
  });

  factory Promotion.fromJson(Map<String, dynamic> json) {
    return Promotion(
      promotionID: json['promotion_id'],
      description: json['description'],
      discountAmount: json['discount_amount'].toDouble(),
      validFrom: json['valid_from'],
      validTo: json['valid_to'],
    );
  }
}
