import 'package:flutter/material.dart';
import 'cook.dart';

class LandingPage extends StatelessWidget {
  const LandingPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Enter a website url to get started!'),
            TextField(),
            ElevatedButton(
              onPressed: () {
                Navigator.pushNamed(context, '/recipe');
              },
              child: Text('Go!'),
            ),
            ElevatedButton(onPressed: null, child: Text('Random')),
          ],
        ),
      ),
    );
  }
}
