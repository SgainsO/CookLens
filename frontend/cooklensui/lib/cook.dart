import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';


class RecipIng extends StatelessWidget {
  const RecipIng({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return Row(
      children: [Expanded(child: MasterWidget())],
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
    try {
      print('Making API call to http://127.0.0.1:3500/scrape');
      final response = await http.get(Uri.parse('http://127.0.0.1:3500/scrape'));
      print('Response status: ${response.statusCode}');
      print('Response body: ${response.body}');
      
      if (response.statusCode == 200) {
        final data = json.decode(response.body);
        setState(() {
          allData = {
            'recipe': List<String>.from(data['recipe']),
            'ingredients': List<String>.from(data['ingredients'])
          };
          isLoading = false;
        });
        print('Data loaded successfully');
      } else {
        print('Error: HTTP ${response.statusCode}');
        setState(() {
          isLoading = false;
        });
      }
    } catch (e) {
      print('Exception occurred: $e');
      setState(() {
        isLoading = false;
      });
    }
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
