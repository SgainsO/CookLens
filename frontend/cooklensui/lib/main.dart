import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'land.dart';
import 'cook.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    SystemChrome.setPreferredOrientations([
      DeviceOrientation.landscapeLeft,
      DeviceOrientation.landscapeRight,
    ]);
    
    return MaterialApp(
      title: 'CookLens',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
      ),
      initialRoute: '/',
      routes: {
        '/': (context) => const LandingPageWrapper(),
        '/recipe': (context) => const RecipIng(),
      },
    );
  }
}

class LandingPageWrapper extends StatelessWidget {
  const LandingPageWrapper({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [Expanded(child: LandingPage())],
      ),
    );
  }
}
