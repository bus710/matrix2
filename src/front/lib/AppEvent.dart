abstract class AppEvent {
  String type;
}

class SwitchEvent extends AppEvent {
  String type;

  SwitchEvent(type){
    this.type = type;
  }
}

