import 'package:flutter/material.dart';
import 'cook.dart';

class LandingPage extends StatelessWidget {
  const LandingPage({super.key});


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromARGB(255, 179, 164, 205),
      body: Stack(
        children: [
          // Top-aligned title
          Align(
            alignment: Alignment.topCenter,
            child: Padding(
              padding: const EdgeInsets.only(top: 10), // move it down a bit
              child: Text(
                'CookLens',
                style: TextStyle(
                  fontSize: 48,
                  fontWeight: FontWeight.bold,
                  color: Colors.white,
                ),
              ),
            ),
          ),

          // Center-aligned content
          const CenterContent(),
        ],
      ),
    );
  }
}

class CenterContent extends StatefulWidget {
  const CenterContent({super.key}); 

  @override
  State<CenterContent> createState() => _CenterContentState();
}

class _CenterContentState extends State<CenterContent> {
  
  final inputControl = TextEditingController(); 

  @override
  void dispose() {
    inputControl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return  Align(
            alignment: Alignment.center,
            child: Column(
              mainAxisSize: MainAxisSize.min, // shrink-wrap content
              children: [
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 20),
                  child: TextField(
                    decoration: InputDecoration(
                      hintText: 'Website URL',
                    ),
                    textAlign: TextAlign.center,
                    style: TextStyle(
                      color: Colors.black,
                      fontSize: 22,
                      fontStyle: FontStyle.italic,
                    ),
                    controller: inputControl,
                  ),
                ),
                SizedBox(height: 20),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        minimumSize: Size(100, 50),
                      ),
                      onPressed: () {
                        Navigator.pushNamed(context, '/recipe', arguments: {'link': inputControl.text});
                      },
                      child: Text('Go!'),
                    ),
                    SizedBox(width: 20), // Add spacing between buttons
                    ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        minimumSize: Size(100, 50),
                      ),
                      onPressed: null,
                      child: Text('Random'),
                    ),
                  ],
                ),
        ],
      ),
    );
  }
}