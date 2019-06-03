import 'package:flutter_web/material.dart';
import 'package:front/AppPage.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Matrix 2',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        fontFamily: 'OpenSans'
      ),
      home: AppPage(title: 'Matrix 2'),
    );
  }
}

