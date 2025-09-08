import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

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
        body: RecipeTitleContainer(),
      ),
    );
  }
}

class RecipeTitleContainer extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(16.0),
      color: Colors.deepPurple,
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
          Expanded(child: RecipeContainer()), // <- Make sure this widget exists
        ],
      ),
    );
  }
}
  class RecipeContainer extends StatefulWidget {
    @override
    State<RecipeContainer> createState() => _RecipeContainerState();
  }

  class _RecipeContainerState extends State<RecipeContainer> {
    List<String> recipes = [
      'Spaghetti Carbonara',
      'Chicken Alfredo',
      'Beef Stroganoff',
      'Vegetable Stir Fry',
      'Tacos',
      'Caesar Salad',
      'Grilled Cheese Sandwich',
      'Pancakes',
      'Chocolate Chip Cookies',
      'Apple Pie'
    ];

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