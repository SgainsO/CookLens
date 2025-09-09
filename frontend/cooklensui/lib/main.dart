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
        body: Row(children: [ Expanded(flex: 3, child: RecipTitleContainer()),
         Expanded(flex: 2, child: IngreTitleContainer())],)
      ),
    );
  }
}

class IngreTitleContainer extends StatelessWidget {
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
          Expanded(child: IngreContainer()), // <- Make sure this widget exists
        ],
      ),
    );
  }
}

class RecipTitleContainer extends StatelessWidget {
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
          Expanded(child: RecipeContainer()), // <- Make sure this widget exists
        ],
      ),
    );
  }
}



  class IngreContainer extends StatefulWidget {
    @override
    State<IngreContainer> createState() => _IngreContainerState();
  }

  class RecipeContainer extends StatefulWidget {
    @override
    State<RecipeContainer> createState() => _RecipeContainerState();
  }

  class _RecipeContainerState extends State<RecipeContainer> {
      List<String> recips = [
    "Preheat the oven to 375°F (190°C).",
    "Chop onions, tomatoes, and garlic finely.",
    "Heat 2 tablespoons of olive oil in a pan.",
    "Add onions and sauté until golden brown.",
    "Stir in garlic and cook for 1 more minute.",
    "Add chopped tomatoes and simmer for 10 minutes.",
    "Season with salt, pepper, and oregano.",
    "Boil pasta in salted water until al dente.",
    "Drain pasta and mix with the tomato sauce.",
    "Grate parmesan cheese on top before serving.",
    "Whisk eggs and milk together in a bowl.",
    "Dip bread slices into egg mixture and fry until golden.",
    "Melt butter in a pan and cook pancakes until bubbly.",
    "Mix flour, sugar, and baking powder in a bowl.",
    "Add milk and eggs to form a smooth batter.",
    "Pour batter into a greased baking dish.",
    "Bake for 30 minutes until golden and set.",
    "Marinate chicken with lemon juice and spices.",
    "Grill chicken on medium heat for 6–8 minutes each side.",
    "Serve hot with rice or salad."
  ];

    @override 
    Widget build(BuildContext context) {
      return ListView.builder(
        itemCount: recips.length,
        itemBuilder: (context, index) {
          return ListTile(
            title: Text(recips[index]),
          );
        },
      );
    }
  } 

  class _IngreContainerState extends State<IngreContainer> {
    List<String> Ingres = [
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
      itemCount: Ingres.length,
      itemBuilder: (context, index) {
        return ListTile(
          title: Text(Ingres[index]),
        );
      },
    );
  }
  }