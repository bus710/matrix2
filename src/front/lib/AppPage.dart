import 'dart:async';

import 'package:flutter_web/material.dart';
import 'package:flutter_web/widgets.dart';
// import 'package:front/bloc.dart';
// import 'package:front/event.dart';

import 'package:front/AppUtil.dart';

class AppPage extends StatefulWidget {
  AppPage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _AppPageState createState() => _AppPageState();
}

class _AppPageState extends State<AppPage> {
  // final _bloc = AppBLoC();
  StreamController<String> controlStream;
  bool switchValue = false;
  List<bool> activeList; // to check enabled/disabled
  List<String> rgbList; // to store each point's color
  List<List> colorList; // to store rgbList's instances
  int rSliderValue;
  int gSliderValue;
  int bSliderValue;

  _AppPageState() {
    controlStream = new StreamController();

    activeList = new List<bool>(64);
    for (int i = 0; i < activeList.length; i++) {
      activeList[i] = false;
    }

    rgbList = new List<String>();
    rgbList = '0,0,0'.split(',');

    colorList = new List<List>();
    for (int i = 0; i < 64; i++) {
      colorList.add(rgbList);
    }

    rSliderValue = 0;
    gSliderValue = 0;
    bSliderValue = 0;
  }

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(centerTitle: true, title: Text(widget.title)),
      body: Center(
        child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              // _getText(),
              // _getButton(),
              // _getSwitch(),
              _getDisplay(),
              _getController(),
            ]),
      ),
    );
  }

  @override
  dispose() {
    super.dispose();
    // _bloc.dispose();
  }

  Container _getDisplay() {
    return Container(
      child: StreamBuilder(
          stream: controlStream.stream,
          initialData: 'Empty',
          builder: (context, snapshot) {
            _setColor(snapshot.data.toString());

            return Container(
              height: 350,
              width: 350,
              child: GridView.count(
                crossAxisCount: 8,
                padding: const EdgeInsets.symmetric(vertical: 3, horizontal: 3),
                children: List.generate(64, (index) {
                  return _getPoint(index);
                }),
              ),
            );
          }),
    );
  }

  void _setColor(String data) {
    switch (data) {
      case 'All':
        {
          for (int i = 0; i < activeList.length; i++) {
            activeList[i] = true;
          }
        }
        break;
      case 'None':
        {
          for (int i = 0; i < activeList.length; i++) {
            activeList[i] = false;
          }
        }
        break;
      case 'Apply':
        {
          // apply();
          // _bloc.app_event_sink.add(SwitchEvent(false));
        }
        break;
      case 'Empty':
        {/* Do nothing */}
        break;
      default:
        {
          List<String> tmp = data.split(":");
          rgbList = tmp[1].split(",");

          for (int i = 0; i < 64; i++) {
            if (activeList[i]) {
              colorList[i] = rgbList;
            }
          }
        }
    }
  }

  CustomPaint _getPoint(int number) {
    return CustomPaint(
      foregroundPainter: MyPainter(activeList[number]),
      child: GestureDetector(
        onTap: () {
          setState(() {
            if (activeList[number] == true) {
              activeList[number] = false;
            } else {
              activeList[number] = true;
            }
            controlStream.sink.add('Empty');
            print('$number - ' + activeList[number].toString());
          });
        },
        child: Container(
          decoration: new BoxDecoration(
            color: Color.fromARGB(
                0xff,
                int.parse(colorList[number][0]) * 3,
                int.parse(colorList[number][1]) * 3,
                int.parse(colorList[number][2]) * 3),
            borderRadius: new BorderRadius.all(
              Radius.circular(5),
            ),
          ),
          margin: const EdgeInsets.symmetric(vertical: 3, horizontal: 3),
          padding: const EdgeInsets.symmetric(vertical: 3, horizontal: 3),
          width: 1,
          height: 1,
          // color: Colors.black,
        ),
      ),
    );
  }

  Container _getController() {
    return Container(
      child: Padding(
        padding: const EdgeInsets.only(left: 5, top: 20, right: 5, bottom: 10),
        child: Center(
          child: Column(children: <Widget>[
            _sliderAndValue('R', 150, rSliderValue),
            _sliderAndValue('G', 150, gSliderValue),
            _sliderAndValue('B', 150, bSliderValue),
            _buttons(),
          ]),
        ),
      ),
    );
  }

  Container _sliderAndValue(String name, double width, int _sliderValue) {
    return Container(
      padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 60),
      child: Row(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: <Widget>[
            Text(name, style: TextStyle(fontSize: 18)),
            Container(
              width: 240,
              height: 25,
              child: Slider(
                min: 0,
                max: 63,
                value: _sliderValue.toDouble(),
                onChanged: (value) => _sliderChanged(name, value),
              ),
            ),
            Text(_sliderValue.toString(), style: TextStyle(fontSize: 18)),
          ]),
    );
  }

  void _sliderChanged(String name, double value) {
    setState(() {
      switch (name) {
        case 'R':
          {
            rSliderValue = value.round();
          }
          break;
        case 'G':
          {
            gSliderValue = value.round();
          }
          break;
        case 'B':
          {
            bSliderValue = value.round();
          }
          break;
        default:
          break;
      }
      print('sliderValue:$rSliderValue,$gSliderValue,$bSliderValue');
      controlStream.sink
          .add('sliderValue:$rSliderValue,$gSliderValue,$bSliderValue');
    });
  }

  Container _buttons() {
    return Container(
      padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 30),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: <Widget>[
          _mutualButton('All', 100, Theme.of(context).accentColor),
          _mutualButton('None', 100, Theme.of(context).accentColor),
          _mutualButton('Apply', 50, Colors.red),
        ],
      ),
    );
  }

  Padding _mutualButton(
    String name,
    double width,
    Color color,
  ) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 20, horizontal: 1),
      child: SizedBox(
        width: width,
        height: 40,
        child: RaisedButton(
          onPressed: () => setState(() {
                // print('$name');
                print('outgoing data: ' + name);
                controlStream.sink.add(name);
              }),
          child: _conditionalTextOrIcon(name),
          textColor: Colors.white,
          splashColor: Colors.white,
          color: color,
          elevation: 4,
        ),
      ),
    );
  }

  Widget _conditionalTextOrIcon(String name) {
    // https://github.com/flutter/flutter_web/tree/master/examples/gallery/web/assets
    if (name != "Apply") {
      return Text(name, style: TextStyle(fontSize: 20));
    } else {
      return new Icon(Icons.send);
    }
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
  //           },
  //         );
  //       });
  // }
}
