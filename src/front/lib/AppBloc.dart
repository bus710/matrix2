import 'dart:async';
import 'dart:convert';
import 'package:front/AppEvent.dart';

/* Please use html.WebSocket for websocket instead of these:
- import 'package:web_socket_channel/io.dart';
- import 'package:web_socket_channel/web_socket_channel.dart';
The detail can be found:
- https://api.dartlang.org/stable/2.3.1/dart-html/WebSocket-class.html */
import 'dart:html' as html;

class AppBLoC {
  /* App Stream from the socket to the front.
  - app_sink is used to put some data to the stream.
  - stream_app is connected to the front side components.  */
  final _appStreamController = StreamController<bool>();
  StreamSink<bool> get app_sink => _appStreamController.sink;
  Stream<bool> get stream_app => _appStreamController.stream;

  /* App Event: the front > the event > the socket. */
  final _appEventController = StreamController<AppEvent>();
  Sink<AppEvent> get app_event_sink => _appEventController.sink;

  // init web socket channel
  var _webSocket;

  // constructor
  AppBLoC() {
    // Connect the event and the handler.
    _appEventController.stream.listen(_handler);

    // Init the socket and its handlers.
    socketInit();
  }

  /* This handler connects these: 
  - The event of button/slider control
  - The websocket.
  - So this is the place to make some action (i.e. marshaling). */
  _handler(AppEvent event) {
    print("event - " + event.type.toString());
    if (_webSocket != null && _webSocket.readyState == html.WebSocket.OPEN) {
      _webSocket
          .send(json.encode({"type": event.type.toString(), "data": event.data}));
    } else {
      print('WebSocket not connected, message data not sent');
    }
  }

  dispose() {
    _appStreamController.close();
    _appEventController.close();
    _webSocket.close();
  }

  socketInit() {
    _webSocket = new html.WebSocket(
        'ws://' + html.window.location.hostname + ':3000/message');

    _webSocket.onOpen.listen((e) {
      print("opened");
    });

    _webSocket.onClose.listen((e) {
      print("closed");
    });

    _webSocket.onMessage.listen((e) {
      String type = json.decode(e.data)["type"];
      String data = json.decode(e.data)["data"];
      print("received - ${type} / ${data}");

      // if (data == "true") {
      //   app_sink.add(true);
      // } else if (data == "false") {
      //   app_sink.add(false);
      // } else {
      //   app_sink.add(false);
      // }
    });
  }

  /* "to json" can be found:
  - https://medium.com/flutter-community/working-with-apis-in-flutter-8745968103e9
  - https://github.com/PoojaB26/ParsingJSON-Flutter/blob/master/lib/model/post_model.dart
  - https://flutter.dev/docs/development/data-and-backend/json
   */
}
