import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {

    SystemChrome.setPreferredOrientations([
      DeviceOrientation.landscapeLeft,
      DeviceOrientation.landscapeRight,
    ]);
    return MaterialApp(
      theme: ThemeData(
        // This is the theme of your application.
        //
        // TRY THIS: Try running your application with "flutter run". You'll see
        // the application has a purple toolbar. Then, without quitting the app,
        // try changing the seedColor in the colorScheme below to Colors.green
        // and then invoke "hot reload" (save your changes or press the "hot
        // reload" button in a Flutter-supported IDE, or press "r" if you used
        // the command line to start the app).
        //
        // Notice that the counter didn't reset back to zero; the application
        // state is not lost during the reload. To reset the state, use hot
        // restart instead.
        //
        // This works for code too, not just values: Most code changes can be
        // tested with just a hot reload.
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
      ),
      home: Scaffold(
        body: MasterWidget(),
      ),
    );
  }
}

class MasterWidget extends StatefulWidget {
  @override
  _MasterWidgetState createState() => _MasterWidgetState(); 
}

class _MasterWidgetState extends State<MasterWidget>{
  Map<String, List<String>> allData = {};
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    fetchAllData();
  }


  fetchAllData() async {
    final response = await http.get(Uri.parse('http://127.0.0.1:8080/scrape'));
    final data = json.decode(response.body);
    setState(() {
      allData = {
        'recipe': List<String>.from(data['recipe']),
        'ingredients': List<String>.from(data['ingredients'])
      };
      isLoading = false;
    });

  }
  @override
  Widget build(BuildContext context)
  {
    if (isLoading) return Center(child: CircularProgressIndicator());
    
    return Scaffold(
        body: Row(children: [ 
          Expanded(flex: 3, child: RecipTitleContainer(recipes: allData['recipe'] ?? [])),
          Expanded(flex: 2, child: IngreTitleContainer(ingredients: allData['ingredients'] ?? []))
        ],)
      );
  }

}

class IngreTitleContainer extends StatelessWidget {
  final List<String> ingredients;
  
  IngreTitleContainer({required this.ingredients});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(16.0),
      color: const Color.fromARGB(255, 179, 164, 205),
      child: Column(
        children: [
          Text(
            'Ingre List',
            style: TextStyle(
              color: Colors.white,
              fontSize: 24,
              fontWeight: FontWeight.bold,
            ),
          ), // <- Added missing closing parenthesis and comma
          Expanded(child: IngreContainer(ingredients: ingredients)), // <- Make sure this widget exists
        ],
      ),
    );
  }
}

class RecipTitleContainer extends StatelessWidget {
  final List<String> recipes;
  
  RecipTitleContainer({required this.recipes});

  @override
  Widget build(BuildContext context) {
    return Container(padding: EdgeInsets.all(16),
    color: const Color.fromARGB(255, 179, 164, 205),
      child: Column(
        children: [
          Text(
            'Recipe List',
            style: TextStyle(
              color: Colors.white,
              fontSize: 24,
              fontWeight: FontWeight.bold,
            ),
          ), // <- Added missing closing parenthesis and comma
          Expanded(child: RecipeContainer(recipes: recipes)), // <- Make sure this widget exists
        ],
      ),
    );
  }
}




  class IngreContainer extends StatelessWidget {
    final List<String> ingredients;
    
    IngreContainer({required this.ingredients});

    @override
    Widget build(BuildContext context) {
     return ListView.builder(
      itemCount: ingredients.length,
      itemBuilder: (context, index) {
        return ListTile(
          title: Text(ingredients[index]),
        );
      },
    );
  }
  }

  class RecipeContainer extends StatelessWidget {
    final List<String> recipes;
    
    RecipeContainer({required this.recipes});

    @override 
    Widget build(BuildContext context) {
      return ListView.builder(
        itemCount: recipes.length,
        itemBuilder: (context, index) {
          return ListTile(
            title: Text(recipes[index]),
          );
        },
      );
    }
  }