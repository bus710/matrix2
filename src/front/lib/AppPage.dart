import 'package:flutter_web/material.dart';
import 'package:flutter_web/widgets.dart';
// import 'package:front/bloc.dart';
// import 'package:front/event.dart';

class AppPage extends StatefulWidget {
  final String title;

  AppPage({Key key, this.title}) : super(key: key);

  @override
  _AppPageState createState() => _AppPageState();
}

class _AppPageState extends State<AppPage> {
  // final _bloc = AppBLoC();
  bool switchValue = false;

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(widget.title)),
      body: Center(
        child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              // _getText(),
              // _getButton(),
              // _getSwitch(),
            ]),
      ),
    );
  }

  @override
  dispose() {
    super.dispose();
    // _bloc.dispose();
  }

  // _getText() {
  //   return StreamBuilder(
  //     stream: _bloc.stream_app,
  //     initialData: false,
  //     builder: (context, snapshot){
  //       return Center(
  //         child: Text(snapshot.data.toString()),
  //       );
  //     },);
  // }

  // _getButton() {
  //   return RaisedButton(
  //     child: Text("Flip"),
  //     onPressed: () {
  //       print("flip");
  //       _bloc.app_event_sink.add(SwitchEvent(false));
  //     }
  //   );
  // }

  // _getSwitch() {
  //   return StreamBuilder(
  //       stream: _bloc.stream_app,
  //       initialData: switchValue,
  //       builder: (context, snapshot) {
  //         print("snapshot - " + snapshot.data.toString());
  //         switchValue = snapshot.data;
  //         return Switch(
  //           value: switchValue,
  //           onChanged: (bool v) {
  //             print("changed - " + switchValue.toString() + "/" + v.toString());
  //             _bloc.app_event_sink.add(SwitchEvent(switchValue));
  //           },
  //         );
  //       });
  // }
}