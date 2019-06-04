abstract class AppEvent {
  String type;
  String data;
}

class SwitchEvent extends AppEvent {
  String type;
  String data;

  SwitchEvent(type, data){
    this.type = type;
    this.data = data;
  }
}

