import 'package:flutter/material.dart';
import 'package:whatslikeapp/pages/authenticate.dart';

void main() {
  runApp(
    MaterialApp(
      initialRoute: '/',
      routes: {
        '/': (context) => const Authenticate(),
      },
    ),
  );
}
